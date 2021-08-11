package mattermost

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"
	"github.com/pkg/errors"
)

var ErrUnexpectedSignMethod = errors.New("unexpected signing method")
var ErrMissingHeader = errors.Errorf("missing %s: Bearer header", apps.OutgoingAuthHeader)
var ErrActingUserMismatch = errors.New("JWT claim doesn't match actingUserID in context")

type callHandler func(http.ResponseWriter, *http.Request, *apps.CallRequest)

func Init(router *mux.Router, m *apps.Manifest, staticAssets fs.FS, localMode bool) {
	router.HandleFunc(constants.ManifestPath, fManifest(m))
	router.HandleFunc(constants.InstallPath, extractCall(fInstall, localMode))
	router.HandleFunc(constants.BindingsPath, extractCall(fBindings, localMode))

	// OK responses
	router.HandleFunc(constants.BindingPathOK+"/{type}", extractCall(fOK, localMode))
	router.HandleFunc(constants.BindingPathOKEmpty+"/{type}", extractCall(fEmptyOK, localMode))

	// Navigate responses
	router.HandleFunc(constants.BindingPathNavigateExternal+"/{type}", extractCall(fNavigateExternal, localMode))
	router.HandleFunc(constants.BindingPathNavigateInternal+"/{type}", extractCall(fNavigateInternal, localMode))
	router.HandleFunc(constants.BindingPathNavigateInvalid+"/{type}", extractCall(fNavigateInvalid, localMode))

	// Error responses
	router.HandleFunc(constants.BindingPathError+"/{type}", extractCall(fError, localMode))
	router.HandleFunc(constants.BindingPathErrorEmpty+"/{type}", extractCall(fEmptyError, localMode))
	router.HandleFunc(constants.BindingPathMarkdownFormError+"/{type}", extractCall(fMarkdownFormError, localMode))
	router.HandleFunc(constants.BindingPathMarkdownFormErrorMissingField+"/{type}", extractCall(fMarkdownFormErrorMissingField, localMode))

	// Form responses
	router.HandleFunc(constants.BindingPathFormOK+"/{type}", extractCall(fFormOK, localMode))
	router.HandleFunc(constants.BindingPathFullFormOK+"/{type}", extractCall(fFullFormOK, localMode))
	router.HandleFunc(constants.BindingPathRedefineFormOK+"/{type}", extractCall(fFormRedefine, localMode))
	router.HandleFunc(constants.BindingPathEmbeddedFormOK+"/{type}", extractCall(fFormEmbedded, localMode))
	router.HandleFunc(constants.BindingPathFullDisabledOK+"/{type}", extractCall(fFullFormDisabledOK, localMode))
	router.HandleFunc(constants.BindingPathDynamicFormOK+"/{type}", extractCall(fDynamicFormOK, localMode))
	router.HandleFunc(constants.BindingPathFormInvalid+"/{type}", extractCall(fFormInvalid, localMode))
	router.HandleFunc(constants.BindingPathMultiselectForm+"/{type}", extractCall(fFormMultiselect, localMode))
	router.HandleFunc(constants.BindingPathWithButtonsOK+"/{type}", extractCall(fFormWithButtonsOK, localMode))
	router.HandleFunc(constants.BindingPathMarkdownForm+"/{type}", extractCall(fFormWithMarkdownError, localMode))
	router.HandleFunc(constants.BindingPathMarkdownFormWithMissingError+"/{type}", extractCall(fFormWithMarkdownErrorMissingField, localMode))

	// Lookup responses
	router.HandleFunc(constants.BindingPathLookupOK+"/{type}", extractCall(fLookupOK, localMode))
	router.HandleFunc(constants.BindingPathLookupEmpty+"/{type}", extractCall(fLookupEmpty, localMode))
	router.HandleFunc(constants.BindingPathLookupMultiword+"/{type}", extractCall(fLookupMultiword, localMode))
	router.HandleFunc(constants.BindingPathLookupInvalid+"/{type}", extractCall(fLookupInvalid, localMode))

	// Other
	router.HandleFunc(constants.BindingPathHTML+"/{type}", extractCall(fHTML, localMode))
	router.HandleFunc(constants.BindingPathUnknown+"/{type}", extractCall(fUnknown, localMode))

	// Static files
	router.PathPrefix(constants.StaticAssetPath).Handler(http.StripPrefix("/", http.FileServer(http.FS(staticAssets))))

	// Subscription Commands
	router.HandleFunc(constants.SubscribeCommand+"/submit", extractCall(fSubscriptionsCommand(m), localMode))

	// Notifications
	router.HandleFunc(constants.NotifyBotMention, extractCall(fSubscriptionsBotMention(m), localMode))
	router.HandleFunc(constants.NotifyBotJoinedChannel, extractCall(fSubscriptionsBotJoinedChannel(m), localMode))
	router.HandleFunc(constants.NotifyBotLeftChannel, extractCall(fSubscriptionsBotLeftChannel(m), localMode))
	router.HandleFunc(constants.NotifyBotJoinedTeam, extractCall(fSubscriptionsBotJoinedTeam(m), localMode))
	router.HandleFunc(constants.NotifyBotLeftTeam, extractCall(fSubscriptionsBotLeftTeam(m), localMode))

	// OpenDialog
	router.HandleFunc(constants.OtherPathOpenDialog+"/submit", extractCall(postOpenDialogTest(m), localMode))
	router.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogNoResponse, postOpenDialogTestNoResponse)
	router.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogEmptyResponse, postOpenDialogTestEmptyResponse)
	router.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogEphemeralResponse, postOpenDialogTestEphemeralResponse)
	router.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogUpdateResponse, postOpenDialogTestUpdateResponse)
	router.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogBadResponse, postOpenDialogTestBadResponse)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := errors.Errorf("path not found: %s", r.URL.Path)
		httputils.WriteJSON(w, apps.NewErrorCallResponse(err))
	})
}

func extractCall(f callHandler, localMode bool) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := apps.CallRequestFromJSONReader(r.Body)
		if err != nil {
			utils.WriteBadRequestError(rw, err)
			return
		}

		if localMode {
			claims, err := checkJWT(r)
			if err != nil {
				utils.WriteBadRequestError(rw, err)
				return
			}

			if data.Context.ActingUserID != "" && data.Context.ActingUserID != claims.ActingUserID {
				utils.WriteBadRequestError(rw, ErrActingUserMismatch)
				return
			}

			utils.DumpObject(data)
		}

		f(rw, r, data)
	}
}

func checkJWT(req *http.Request) (*apps.JWTClaims, error) {
	authValue := req.Header.Get(apps.OutgoingAuthHeader)
	if !strings.HasPrefix(authValue, "Bearer ") {
		return nil, ErrMissingHeader
	}

	jwtoken := strings.TrimPrefix(authValue, "Bearer ")
	claims := apps.JWTClaims{}
	_, err := jwt.ParseWithClaims(jwtoken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %v", ErrUnexpectedSignMethod, token.Header["alg"])
		}
		return []byte(constants.AppSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}

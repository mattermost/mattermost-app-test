package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"

	"github.com/mattermost/mattermost-app-test/path"
)

var ErrUnexpectedSignMethod = errors.New("unexpected signing method")
var ErrMissingHeader = errors.Errorf("missing %s: Bearer header", apps.OutgoingAuthHeader)
var ErrActingUserMismatch = errors.New("JWT claim doesn't match actingUserID in context")

type callHandler func(*apps.CallRequest) apps.CallResponse

func initHTTP(r *mux.Router) {
	r.HandleFunc(path.Manifest, httputils.DoHandleJSON(manifest))

	handleCall(r, path.Install, handleInstall)
	handleCall(r, path.Bindings, handleBindings)

	// Functional command responses
	handleCall(r, path.ProfileView, handleProfileView)
	handleCall(r, path.ProfileCommand, handleProfileCommand)

	// OK responses
	handleCall(r, path.OK, handleOK)
	handleCall(r, path.OKEmpty, handleOKEmpty)

	// // Navigate responses
	// handleCall(r, constants.NavigateExternal, fNavigateExternal)
	// handleCall(r, constants.NavigateInternal, fNavigateInternal)
	// handleCall(r, constants.NavigateInvalid, fNavigateInvalid)

	// // Error responses
	// handleCall(r, constants.Error, fError)
	// handleCall(r, constants.ErrorEmpty, fEmptyError)
	// handleCall(r, constants.MarkdownFormError, fMarkdownFormError)
	// handleCall(r, constants.MarkdownFormErrorMissingField, fMarkdownFormErrorMissingField)

	// // Form responses
	// handleCall(r, constants.Form, fFormOK)
	// handleCall(r, constants.FormFull, fFullFormOK)
	// r.HandleFunc(constants.FormRedefine, handle(fFormRedefine, LocalMode))
	// r.HandleFunc(constants.FormEmbedded, handle(fFormEmbedded, LocalMode))
	// r.HandleFunc(constants.FormFullDisabled, handle(fFullFormDisabledOK, LocalMode))
	// r.HandleFunc(constants.FormDynamic, handle(fDynamicFormOK, LocalMode))
	// r.HandleFunc(constants.FormInvalid, handle(fFormInvalid, LocalMode))
	// r.HandleFunc(constants.FormMultiselect, handle(fFormMultiselect, LocalMode))
	// r.HandleFunc(constants.FormWithButtons, handle(fFormWithButtonsOK, LocalMode))
	// r.HandleFunc(constants.FormMarkdown, handle(fFormWithMarkdownError, LocalMode))
	// r.HandleFunc(constants.FormMarkdownWithMissingError, handle(fFormWithMarkdownErrorMissingField, LocalMode))

	// // Lookup responses
	// r.HandleFunc(constants.Lookup, handle(fLookupOK, LocalMode))
	// r.HandleFunc(constants.LookupEmpty, handle(fLookupEmpty, LocalMode))
	// r.HandleFunc(constants.LookupMultiword, handle(fLookupMultiword, LocalMode))
	// r.HandleFunc(constants.LookupInvalid, handle(fLookupInvalid, LocalMode))

	// // Other
	// r.HandleFunc(constants.HTMLPath, handle(fHTML, LocalMode))
	// r.HandleFunc(constants.UnknownPath, handle(fUnknown, LocalMode))

	// Static files
	r.PathPrefix(path.Static).Handler(http.StripPrefix("/", http.FileServer(http.FS(staticAssets))))

	// // Subscription Commands
	// r.HandleFunc(constants.SubscribeCommand+"/submit", handle(fSubscriptionsCommand(m), LocalMode))

	// // Global Notifications
	// r.HandleFunc(constants.NotifyUserCreated, handle(fSubscriptionsUserCreated(m), LocalMode))
	// r.HandleFunc(constants.NotifyBotMention, handle(fSubscriptionsBotMention(m), LocalMode))
	// r.HandleFunc(constants.NotifyBotJoinedChannel, handle(fSubscriptionsBotJoinedChannel(m), LocalMode))
	// r.HandleFunc(constants.NotifyBotLeftChannel, handle(fSubscriptionsBotLeftChannel(m), LocalMode))
	// r.HandleFunc(constants.NotifyBotJoinedTeam, handle(fSubscriptionsBotJoinedTeam(m), LocalMode))
	// r.HandleFunc(constants.NotifyBotLeftTeam, handle(fSubscriptionsBotLeftTeam(m), LocalMode))

	// // Channel Notifications
	// r.HandleFunc(constants.NotifyUserJoinedChannel, handle(fSubscriptionsUserJoinedChannel(m), LocalMode))
	// r.HandleFunc(constants.NotifyUserLeftChannel, handle(fSubscriptionsUserLeftChannel(m), LocalMode))
	// r.HandleFunc(constants.NotifyPostCreated, handle(fSubscriptionsPostCreated(m), LocalMode))

	// // Team Notifications
	// r.HandleFunc(constants.NotifyUserJoinedTeam, handle(fSubscriptionsUserJoinedTeam(m), LocalMode))
	// r.HandleFunc(constants.NotifyUserLeftTeam, handle(fSubscriptionsUserLeftTeam(m), LocalMode))
	// r.HandleFunc(constants.NotifyChannelCreated, handle(fSubscriptionsChannelCreated(m), LocalMode))

	// // OpenDialog
	// r.HandleFunc(constants.OtherPathOpenDialog+"/submit", handle(postOpenDialogTest(m), LocalMode))
	// r.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogNoResponse, postOpenDialogTestNoResponse)
	// r.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogEmptyResponse, postOpenDialogTestEmptyResponse)
	// r.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogEphemeralResponse, postOpenDialogTestEphemeralResponse)
	// r.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogUpdateResponse, postOpenDialogTestUpdateResponse)
	// r.HandleFunc(constants.OtherPathOpenDialog+constants.OtherOpenDialogBadResponse, postOpenDialogTestBadResponse)

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := errors.Errorf("path not found: %s", r.URL.Path)
		_ = httputils.WriteJSON(w, apps.NewErrorResponse(err))
	})
}

func handleCall(router *mux.Router, path string, f callHandler) {
	router.HandleFunc(path, handle(f))
}

func handle(f callHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		creq, err := apps.CallRequestFromJSONReader(r.Body)
		if err != nil {
			httputils.WriteError(w, utils.NewInvalidError(err))
			return
		}

		if localMode {
			claims, err := checkJWT(r)
			if err != nil {
				httputils.WriteError(w, utils.NewInvalidError(err))
				return
			}

			if creq.Context.ActingUserID != "" && creq.Context.ActingUserID != claims.ActingUserID {
				httputils.WriteError(w, utils.NewInvalidError(ErrActingUserMismatch))
				return
			}

			// log.Println(utils.Pretty(creq))
		}

		cresp := f(creq)
		httputils.WriteJSON(w, cresp)
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
		return []byte(AppSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}

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

	// Form responses
	router.HandleFunc(constants.BindingPathFormOK+"/{type}", extractCall(fFormOK, localMode))
	router.HandleFunc(constants.BindingPathFullFormOK+"/{type}", extractCall(fFullFormOK, localMode))
	router.HandleFunc(constants.BindingPathDynamicFormOK+"/{type}", extractCall(fDynamicFormOK, localMode))
	router.HandleFunc(constants.BindingPathFormInvalid+"/{type}", extractCall(fFormInvalid, localMode))

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

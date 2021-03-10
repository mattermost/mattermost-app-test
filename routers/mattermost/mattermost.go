package mattermost

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-test/constants"
	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
	"github.com/pkg/errors"
)

var ErrUnexpectedSignMethod = errors.New("unexpected signing method")
var ErrMissingHeader = errors.Errorf("missing %s: Bearer header", api.OutgoingAuthHeader)

type callHandler func(http.ResponseWriter, *http.Request, *api.JWTClaims, *apps.Call)

func Init(router *mux.Router, m *apps.Manifest, staticAssets fs.FS) {
	router.HandleFunc(constants.ManifestPath, fManifest(m))
	router.HandleFunc(constants.InstallPath, extractCall(fInstall))
	router.HandleFunc(constants.BindingsPath, extractCall(fBindings))

	// OK responses
	router.HandleFunc(constants.BindingPathOK, extractCall(fOK))
	router.HandleFunc(constants.BindingPathOKEmpty, extractCall(fEmptyOK))

	// Navigate responses
	router.HandleFunc(constants.BindingPathNavigateExternal, extractCall(fNavigateExternal))
	router.HandleFunc(constants.BindingPathNavigateInternal, extractCall(fNavigateInternal))
	router.HandleFunc(constants.BindingPathNavigateInvalid, extractCall(fNavigateInvalid))

	// Error responses
	router.HandleFunc(constants.BindingPathError, extractCall(fError))
	router.HandleFunc(constants.BindingPathErrorEmpty, extractCall(fEmptyError))

	// Form responses
	router.HandleFunc(constants.BindingPathFormOK, extractCall(fFormOK))
	router.HandleFunc(constants.BindingPathFormInvalid, extractCall(fFormInvalid))

	// Lookup responses
	router.HandleFunc(constants.BindingPathLookupOK, extractCall(fLookupOK))
	router.HandleFunc(constants.BindingPathLookupEmpty, extractCall(fLookupEmpty))
	router.HandleFunc(constants.BindingPathLookupMultiword, extractCall(fLookupMultiword))
	router.HandleFunc(constants.BindingPathLookupInvalid, extractCall(fLookupInvalid))

	// Other
	router.HandleFunc(constants.BindingPathHTML, extractCall(fHTML))
	router.HandleFunc(constants.BindingPathUnknown, extractCall(fUnknown))

	// Static files
	router.PathPrefix(constants.StaticAssetPath).Handler(http.StripPrefix("/", http.FileServer(http.FS(staticAssets))))
}

func extractCall(f callHandler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		claims, err := checkJWT(r)
		if err != nil {
			utils.WriteBadRequestError(rw, err)
			return
		}

		data, err := apps.UnmarshalCallFromReader(r.Body)
		if err != nil {
			utils.WriteBadRequestError(rw, err)
			return
		}

		str, _ := json.MarshalIndent(data, "", " ")
		log.Printf("%s", str)
		f(rw, r, claims, data)
	}
}

func checkJWT(req *http.Request) (*api.JWTClaims, error) {
	authValue := req.Header.Get(api.OutgoingAuthHeader)
	if !strings.HasPrefix(authValue, "Bearer ") {
		return nil, ErrMissingHeader
	}

	jwtoken := strings.TrimPrefix(authValue, "Bearer ")
	claims := api.JWTClaims{}
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

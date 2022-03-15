package main

import (
	"fmt"
	"log"
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
var ErrActingUserMismatch = errors.New("JWT claim doesn't match actingUser.Id in context")

type callHandler func(*apps.CallRequest) apps.CallResponse

func initHTTP(r *mux.Router) {
	r.HandleFunc(path.Manifest, httputils.DoHandleJSON(manifest))

	// Static files
	r.PathPrefix(path.Static).Handler(http.StripPrefix("/", http.FileServer(http.FS(staticAssets))))

	handleCall(r, path.Install, handleInstall)
	handleCall(r, path.Bindings, handleBindings)

	initHTTPEmbedded(r)
	initHTTPError(r)
	initHTTPForms(r)
	initHTTPLookup(r)
	initHTTPNavigate(r)
	initHTTPOK(r)
	initHTTPSubscriptions(r)

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := errors.Errorf("path not found: %s", r.URL.Path)
		_ = httputils.WriteJSON(w, apps.NewErrorResponse(err))
	})
}

func handle(f callHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		creq, err := apps.CallRequestFromJSONReader(r.Body)
		if err != nil {
			httputils.WriteError(w, utils.NewInvalidError(err))
			return
		}

		err = checkJWT(r, creq)
		if err != nil {
			httputils.WriteError(w, utils.NewInvalidError(err))
			return
		}

		cresp := f(creq)

		if len(creq.Values) > 0 {
			log.Println(creq.Path, creq.Context.Subject, utils.ToJSON(creq.Values), cresp.Type)
		} else {
			log.Println(creq.Path, creq.Context.Subject, cresp.Type)
		}

		_ = httputils.WriteJSON(w, cresp)
	}
}

func handleCall(router *mux.Router, path string, f callHandler) {
	router.HandleFunc(path, handle(f))
}

func handleError(text string) callHandler {
	return func(_ *apps.CallRequest) apps.CallResponse {
		return apps.CallResponse{
			Type: apps.CallResponseTypeError,
			Text: text,
		}
	}
}

func handleForm(f apps.Form) callHandler {
	return func(_ *apps.CallRequest) apps.CallResponse {
		return apps.NewFormResponse(f)
	}
}

type lookupResponse struct {
	Items []apps.SelectOption `json:"items"`
}

func handleLookup(items []apps.SelectOption) callHandler {
	return func(_ *apps.CallRequest) apps.CallResponse {
		return apps.NewDataResponse(lookupResponse{items})
	}
}

func checkJWT(req *http.Request, creq *apps.CallRequest) error {
	if !localMode {
		return nil
	}

	authValue := req.Header.Get(apps.OutgoingAuthHeader)
	if !strings.HasPrefix(authValue, "Bearer ") {
		return ErrMissingHeader
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
		return err
	}

	if creq.Context.ActingUser != nil && creq.Context.ActingUser.Id != claims.ActingUserID {
		return utils.NewInvalidError(ErrActingUserMismatch)
	}

	return nil
}

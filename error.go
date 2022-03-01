package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"

	"github.com/mattermost/mattermost-app-test/path"
)

func initHTTPError(r *mux.Router) {
	handleCall(r, path.ErrorDefault, handleError("Error"))
	handleCall(r, path.ErrorEmpty, handleError(""))
	handleCall(r, path.ErrorMarkdownForm, handleErrorMarkdownForm)
	handleCall(r, path.ErrorMarkdownFormMissingField, handleErrorMarkdownFormMissingField)
	r.HandleFunc(path.Error404, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	r.HandleFunc(path.Error500, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	r.HandleFunc(path.InvalidUnknownType, httputils.DoHandleJSON(apps.CallResponse{
		Type: "unknown",
	}))

	r.HandleFunc(path.InvalidHTML, httputils.DoHandleData("text/html", []byte(`
<!DOCTYPE html>
<html>
	<head>
	</head>
	<body>
		<p>HTML example</p>
	</body>
</html>
`)))
}

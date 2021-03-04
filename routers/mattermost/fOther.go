package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-test/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func fUnknown(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	resp := apps.CallResponse{
		Type: "unknown",
	}
	utils.WriteCallResponse(w, resp)
}

func fHTML(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	html := `
		<!DOCTYPE html>
		<html>
			<head>
			</head>
			<body>
				<p>HTML example</p>
			</body>
		</html>
		`

	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(html))
}

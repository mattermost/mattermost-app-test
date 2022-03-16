module github.com/mattermost/mattermost-app-test

go 1.16

require (
	github.com/aws/aws-lambda-go v1.23.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.10.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/mattermost/mattermost-plugin-apps v0.7.1-0.20220315231457-c14f3d4a76de
	github.com/mattermost/mattermost-server/v6 v6.0.0-20220315170027-7710a2fe3741
	github.com/pkg/errors v0.9.1
)

replace github.com/mattermost/mattermost-plugin-apps v0.7.1-0.20220114173732-f10616dc8752 => ../mattermost-plugin-apps

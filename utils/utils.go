package utils

import (
	"encoding/json"
	"log"
	"strings"
)

func GetStringFromMapInterface(in map[string]interface{}, key, def string) string {
	if len(in) == 0 {
		return def
	}

	v, ok := in[key]
	if !ok {
		return def
	}

	out, ok := v.(string)
	if !ok {
		return def
	}

	return out
}

func GetIconURL(siteURL, name, appID string) string {
	return strings.TrimRight(siteURL, "/") + "/plugins/com.mattermost.apps/apps/" + string(appID) + "/static/" + name
}

func DumpObject(c interface{}) {
	b, _ := json.MarshalIndent(c, "", "    ")
	log.Printf("%s\n", string(b))
}

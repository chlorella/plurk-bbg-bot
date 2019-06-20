package pkg

import (
	"flag"
	"fmt"
	"github.com/clsung/plurgo/plurkgo"
	"github.com/garyburd/go-oauth/oauth"
	"log"
)

var credPath = flag.String("config", "config.json", "Path to configuration file containing the application's credentials.")
var lang = "ja"


func getProfile (accessToken *oauth.Credentials){
	result, err := plurgo.CallAPI(accessToken, "/APP/Profile/getOwnProfile", map[string]string{})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	fmt.Println(string(result))
}

func addAllAsFriend (accessToken *oauth.Credentials){
	result, err := plurgo.CallAPI(accessToken, "/APP/Alerts/addAllAsFriends", map[string]string{})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	fmt.Println(string(result))
}

func sendPlurk (accessToken *oauth.Credentials, content string, qualifier string) {
	var data = map[string]string{}
	data["content"] = content
	data["qualifier"] = qualifier
	data["lang"] = lang
	result, err := plurgo.CallAPI(accessToken, "/APP/Timeline/plurkAdd", data)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	fmt.Println(string(result))
}
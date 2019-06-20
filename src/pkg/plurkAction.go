package pkg

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/clsung/plurgo/plurkgo"
	"github.com/garyburd/go-oauth/oauth"
	"io/ioutil"
	"log"
)

var credPath = flag.String("config", "config.json", "Path to configuration file containing the application's credentials.")
var lang = "ja"

func GetPlurkOauth() (*oauth.Credentials, error) {
	flag.Parse()
	plurkOAuth, err := plurgo.ReadCredentials(*credPath)
	if err != nil {
		log.Fatalf("Error reading credential, %v", err)
		return nil, err
	}
	accessToken, authorized, err := plurgo.GetAccessToken(plurkOAuth)

	if authorized {
		bytes, err := json.MarshalIndent(plurkOAuth, "", "  ")
		if err != nil {
			log.Fatalf("failed to store credential: %v", err)
			return nil, err
		}
		err = ioutil.WriteFile(*credPath, bytes, 0700)
		if err != nil {
			log.Fatalf("failed to write credential: %v", err)
			return nil, err
		}
	}
	return accessToken, nil
}

func GetProfile(accessToken *oauth.Credentials) map[string]string {
	result, err := plurgo.CallAPI(accessToken, "/APP/Profile/getOwnProfile", map[string]string{})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	var resultData = map[string]string{}

	if err := json.Unmarshal(result, &resultData); err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Print(resultData)
	return resultData
}

func AddAllAsFriend(accessToken *oauth.Credentials) map[string]string {
	var resultData = map[string]string{}
	result, err := plurgo.CallAPI(accessToken, "/APP/Alerts/addAllAsFriends", map[string]string{})
	if err != nil {
		log.Fatalf("failed: %v", err)
	}

	if err := json.Unmarshal(result, &resultData); err != nil {
		log.Fatalf("failed: %v", err)
	}

	log.Print(resultData)
	return resultData
}

func SendPlurk(accessToken *oauth.Credentials, content string, qualifier string) {
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

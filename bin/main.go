package main

import (
	plurkAction "github.com/chlorella/plurk-bbg-bot/src/pkg"
	"github.com/robfig/cron"
	"log"
)

func main() {
	c := cron.New()
	accessToken, err := plurkAction.GetPlurkOauth()

	profile := plurkAction.GetProfile(accessToken)
	log.Print(profile["alerts_count"])

	if err != nil{
		c.AddFunc("@hourly", func() {plurkAction.AddAllAsFriend(accessToken)})
	}
	c.Start()
	//plurkAction.AddAllAsFriend(accessToken)
	//plurkAction.SendPlurk(accessToken, "(dice10)(dice)", "plays")
	defer c.Stop()

	select{}
}


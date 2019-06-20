package main

import plurkAction "github.com/chlorella/plurk-bbg-bot/src/pkg"

func main() {
	accessToken, _ := plurkAction.GetPlurkOauth()

	plurkAction.GetProfile(accessToken)
	//plurkAction.AddAllAsFriend(accessToken)
	plurkAction.SendPlurk(accessToken, "(dice10)(dice)", "plays")
}


package main

import plurkAction "github.com/chlorella/plurk-bbg-bot/src/pkg"

func main() {
	accessToken, _ := plurkAction.GetPlurkOauth()

	plurkAction.GetPro()
}


package main

import (
	"github.com/nwoik/calibotapi/utils"
)

func main() {
	clans := utils.ReadFile("./clan.json")

	// member := utils.CreateMember("12345", "nikka", "deez")

	// clans.Add(*member)
	clan := utils.AddClan(clans, "My Clan", "129712342", "123131")
	clan.GetMembers()

	utils.WriteFile("./clan2.json", clans)
}

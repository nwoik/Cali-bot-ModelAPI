package main

import (
	c "github.com/nwoik/calibotapi/clans"
	"github.com/nwoik/calibotapi/utils"
)

func main() {
	clans := c.NewClans().Open("./clan.json")
	// members := utils.Open("members.json")

	clan := utils.CreateClan("My Clan", "129712342", "123131")
	clan.AddMember(utils.CreateMember("12345", "nikka", "deez", "123131414"))

	clans.AddClan(clan)
	// members = append(members, utils.CreateMember("12345", "nikka", "deez", "123131414"))

	clans.Close("./clan.json")
	// utils.Close("./members.json", members)
}

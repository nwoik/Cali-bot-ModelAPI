package main

import (
	c "github.com/nwoik/calibotapi/clans"
	"github.com/nwoik/calibotapi/utils"
)

func main() {
	clans := c.NewClans().Open("./clan.json")
	members := utils.Open("members.json")
	// member := utils.CreateMember("12345", "nikka", "deez", "123131414")

	// clans.Add(*member)
	clan := utils.AddClan(clans, "My Clan", "129712342", "123131")
	clan.SetMembers(utils.AddMember(clan.GetMembers(), "12345", "nikka", "deez", "123131414"))
	members = utils.AddMember(members, "12345", "nikka", "deez", "123131414")

	clans.Close("./clan2.json")
	utils.Close("./members.json", members)
}

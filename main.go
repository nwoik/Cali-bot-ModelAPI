package main

import (
	c "github.com/nwoik/calibotapi/clan"
	m "github.com/nwoik/calibotapi/member"
)

func main() {
	clans := c.Open("./clan.json")
	members := m.Open("members.json")

	clan := c.CreateClan("My Clan", "129712342", "123131")
	clan.AddMember(m.CreateMember("12345", "nikka", "deez", "123131414"))
	clan.BlacklistMember(m.CreateMember("12345", "nikka", "deez", "123131414"))

	clans = append(clans, clan)
	members = append(members, m.CreateMember("12345", "nikka", "deez", "123131414"))

	c.Close("./clan.json", clans)
	m.Close("./members.json", members)
}

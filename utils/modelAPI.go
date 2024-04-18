package utils

import (
	clans "github.com/nwoik/calibotapi/clans"
)

func AddClan(clans *clans.Clans, name string, clanid string, guildid string) *clans.Clan {
	clan := CreateClan(name, clanid, guildid)
	return clans.AddClan(clan)
}

func AddMember(members []*clans.Member, nick string, ign string, igid string, userid string) []*clans.Member {
	member := CreateMember(nick, ign, igid, userid)

	members = append(members, &member)

	return members
}

func CreateClan(name string, clanid string, guildid string) clans.Clan {
	newClan := clans.NewClan().
		SetName(name).
		SetClanID(clanid).
		SetGuildID(guildid)

	return *newClan
}

func CreateMember(nick string, ign string, igid string, userid string) clans.Member {
	newMember := clans.NewMember().
		SetNick(nick).
		SetIGN(ign).
		SetIGID(igid).
		SetUserID(userid)

	return *newMember
}

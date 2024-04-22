package clans

type Clan struct {
	Name       string    `json:"name"`
	GuildID    string    `json:"guildid"`
	MemberRole string    `json:"memberrole"`
	LeaderID   string    `json:"leaderid"`
	ClanID     string    `json:"clanid"`
	Members    []*Member `json:"members"`
	Blacklist  []*Member `json:"blacklist"`
}

func NewClan() *Clan {
	return &Clan{}
}

func (clan *Clan) SetName(name string) *Clan {
	clan.Name = name
	return clan
}

func (clan *Clan) SetGuildID(id string) *Clan {
	clan.GuildID = id
	return clan
}

func (clan *Clan) SetMemberRole(role string) *Clan {
	clan.MemberRole = role
	return clan
}

func (clan *Clan) SetLeaderID(id string) *Clan {
	clan.LeaderID = id
	return clan
}

func (clan *Clan) SetClanID(id string) *Clan {
	clan.ClanID = id
	return clan
}

func (clan *Clan) SetMembers(members []*Member) *Clan {
	clan.Members = members
	return clan
}

func (clan *Clan) AddMember(member *Member) *Member {
	clan.Members = append(clan.Members, member)

	return member
}

func (clan *Clan) BlacklistMember(member *Member) *Member {
	clan.Blacklist = append(clan.Blacklist, member)

	return member
}

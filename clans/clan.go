package clans

type Clan struct {
	Name    string    `json:"name"`
	GuildID string    `json:"guildid"`
	ClanID  string    `json:"clanid"`
	Members []*Member `json:"members"`
}

func NewClan() *Clan {
	return &Clan{}
}

func (clan *Clan) GetName() string {
	return clan.Name
}

func (clan *Clan) GetGuildID() string {
	return clan.GuildID
}

func (clan *Clan) GetClanID() string {
	return clan.ClanID
}

func (clan *Clan) GetMembers() []*Member {
	members := clan.Members

	if members == nil {
		clan.SetMembers(make([]*Member, 0))
		members = clan.Members
	}

	return members
}

func (clan *Clan) SetName(name string) *Clan {
	clan.Name = name
	return clan
}

func (clan *Clan) SetGuildID(id string) *Clan {
	clan.GuildID = id
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
	clan.Members = clan.GetMembers()
	clan.Members = append(clan.GetMembers(), member)

	return member
}

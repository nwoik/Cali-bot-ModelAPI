package clans

type Clans struct {
	ClansList []*Clan `json:"clansList"`
}

func (clans *Clans) AddClan(clan Clan) *Clan {
	clans.ClansList = clans.GetClans()
	clans.ClansList = append(clans.ClansList, &clan)

	return &clan
}

func (clans *Clans) GetClans() []*Clan {
	theClans := clans.ClansList

	if theClans == nil {
		clans.SetClans(make([]*Clan, 0))
		theClans = clans.ClansList
	}

	return theClans
}

func (clans *Clans) SetClans(clansList []*Clan) *Clans {
	clans.ClansList = clansList
	return clans
}

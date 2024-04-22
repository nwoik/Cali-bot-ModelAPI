package clans

type Member struct {
	UserID string `json:"userid"`
	Nick   string `json:"nick"`
	IGN    string `json:"ign"`
	IGID   string `json:"igid"`
}

func NewMember() *Member {
	return &Member{}
}

func (member *Member) SetUserID(id string) *Member {
	member.UserID = id
	return member
}

func (member *Member) SetNick(nick string) *Member {
	member.Nick = nick
	return member
}

func (member *Member) SetIGN(ign string) *Member {
	member.IGN = ign
	return member
}

func (member *Member) SetIGID(igid string) *Member {
	member.IGID = igid
	return member
}

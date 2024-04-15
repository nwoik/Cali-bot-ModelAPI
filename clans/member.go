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

func (member *Member) GetUserID(id string) string {
	return member.UserID
}

func (member *Member) GetNick(nick string) string {
	return member.Nick
}

func (member *Member) GetIGN(ign string) string {
	return member.IGN
}

func (member *Member) GetIGID(igid string) string {
	return member.IGID
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

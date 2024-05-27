package breaks

type Break struct {
	UserID string `json:"userid" bson:"userid"`
	Reason string `json:"reason" bson:"reason"`
}

func NewBreak(userid string, reason string) *Break {
	return &Break{
		UserID: userid,
		Reason: reason,
	}
}

func (b *Break) SetUserID(id string) *Break {
	b.UserID = id
	return b
}

func (b *Break) SetReason(reason string) *Break {
	b.Reason = reason
	return b
}

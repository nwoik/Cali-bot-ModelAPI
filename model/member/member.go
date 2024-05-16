package member

import (
	"encoding/json"
	"fmt"
	"os"
)

type Rank string

const (
	MEMBER  Rank = "member"
	OFFICER Rank = "officer"
	LEADER  Rank = "leader"
)

type Member struct {
	UserID     string `json:"userid" bson:"userid"`
	Nick       string `json:"nick" bson:"nick"`
	IGN        string `json:"ign" bson:"ign"`
	IGID       string `json:"igid" bson:"igid"`
	ClanID     string `json:"clanid" bson:"clanid"`
	Rank       string `json:"rank" bson:"rank"`
	DateJoined string `json:"datejoined" bson:"datejoined"`
}

func NewMember() *Member {
	return &Member{}
}

func CreateMember(nick string, ign string, igid string, userid string) *Member {
	newMember := NewMember().
		SetNick(nick).
		SetIGN(ign).
		SetIGID(igid).
		SetUserID(userid)

	return newMember
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

func (member *Member) SetClanID(clanid string) *Member {
	member.ClanID = clanid
	return member
}

func (member *Member) SetRank(rank string) *Member {
	member.Rank = rank
	return member
}

func (member *Member) SetDateJoined(dateJoined string) *Member {
	member.DateJoined = dateJoined
	return member
}

func Open(filePath string) []*Member {
	var members []*Member
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return nil
	}

	// Parse JSON data
	err = json.Unmarshal(jsonData, &members)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return nil
	}

	return members
}

func Close(filePath string, members []*Member) {

	// Open the JSON file for reading
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// // Rewind the file pointer to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error rewinding file pointer:", err)
		return
	}

	// // Truncate the file to remove any existing content
	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	encodedData, err := json.MarshalIndent(members, "", "    ")
	if err != nil {
		fmt.Println("Error encoding JSON data:", err)
		return
	}

	// Write the encoded JSON data to the file
	if _, err := file.Write(encodedData); err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}

	fmt.Println("Data has been written to", filePath)
}

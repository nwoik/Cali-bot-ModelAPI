package utils

import (
	"encoding/json"
	"fmt"
	"os"

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

func Open(filePath string) []*clans.Member {
	var members []*clans.Member
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

func Close(filePath string, members []*clans.Member) {

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

	// // Encode and write updated JSON data back to the file
	err = json.NewEncoder(file).Encode(&members)
	if err != nil {
		fmt.Println("Error encoding JSON data:", err)
		return
	}

	fmt.Println("Data has been written to", filePath)
}

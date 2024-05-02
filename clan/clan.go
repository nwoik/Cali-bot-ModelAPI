package clans

import (
	"encoding/json"
	"fmt"
	"os"
)

type Clan struct {
	Name        string   `json:"name"`
	GuildID     string   `json:"guildid"`
	LeaderRole  string   `json:"leaderrole"`
	OfficerRole string   `json:"officerrole"`
	MemberRole  string   `json:"memberrole"`
	ExtraRoles  []string `json:"extraroles"`
	LeaderID    string   `json:"leaderid"`
	ClanID      string   `json:"clanid"`
	Blacklist   []string `json:"blacklist"`
}

func NewClan() *Clan {
	return &Clan{}
}

func CreateClan(name string, clanid string, guildid string) *Clan {
	newClan := NewClan().
		SetName(name).
		SetClanID(clanid).
		SetGuildID(guildid)

	return newClan
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

func (clan *Clan) SetOfficerRole(role string) *Clan {
	clan.OfficerRole = role
	return clan
}

func (clan *Clan) SetLeaderRole(role string) *Clan {
	clan.LeaderRole = role
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

func (clan *Clan) AddRole(role string) string {
	clan.ExtraRoles = append(clan.ExtraRoles, role)

	return role
}

func (clan *Clan) BlacklistMember(userid string) string {
	clan.Blacklist = append(clan.Blacklist, userid)

	return userid
}

func Open(filePath string) []*Clan {
	var clans []*Clan
	// Read JSON file
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return nil
	}

	// Parse JSON data
	err = json.Unmarshal(jsonData, &clans)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return nil
	}

	return clans
}

func Close(filePath string, clans []*Clan) {

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

	encodedData, err := json.MarshalIndent(clans, "", "    ")

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

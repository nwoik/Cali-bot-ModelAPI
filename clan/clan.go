package clans

import (
	"encoding/json"
	"fmt"
	"os"

	m "github.com/nwoik/calibotapi/member"
)

type Clan struct {
	Name      string      `json:"name"`
	GuildID   string      `json:"guildid"`
	Roles     []string    `json:"Roles"`
	LeaderID  string      `json:"leaderid"`
	ClanID    string      `json:"clanid"`
	Members   []*m.Member `json:"members"`
	Blacklist []*m.Member `json:"blacklist"`
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

func (clan *Clan) SetRoles(roles []string) *Clan {
	clan.Roles = roles
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

func (clan *Clan) SetMembers(members []*m.Member) *Clan {
	clan.Members = members
	return clan
}

func (clan *Clan) AddRole(role string) string {
	clan.Roles = append(clan.Roles, role)

	return role
}

func (clan *Clan) AddMember(member *m.Member) *m.Member {
	clan.Members = append(clan.Members, member)

	return member
}

func (clan *Clan) BlacklistMember(member *m.Member) *m.Member {
	clan.Blacklist = append(clan.Blacklist, member)

	return member
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

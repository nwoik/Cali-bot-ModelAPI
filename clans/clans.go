package clans

import (
	"encoding/json"
	"fmt"
	"os"
)

type Clans struct {
	ClansList []*Clan `json:"clansList"`
}

func NewClans() *Clans {
	return &Clans{}
}

func (clans *Clans) AddClan(clan *Clan) *Clan {
	clans.ClansList = append(clans.ClansList, clan)

	return clan
}

func (clans *Clans) SetClans(clansList []*Clan) *Clans {
	clans.ClansList = clansList
	return clans
}

func (clans *Clans) Open(filePath string) *Clans {
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

func (clans *Clans) Close(filePath string) {

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

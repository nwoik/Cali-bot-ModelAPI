package clans

import (
	"encoding/json"
	"fmt"
	"os"
)

type Members struct {
	MembersList []*Member `json:"members"`
}

func NewMembers() *Members {
	return &Members{}
}

func (members *Members) GetMembers() []*Member {
	membersList := members.MembersList

	if membersList == nil {
		members.SetMembers(make([]*Member, 0))
		membersList = members.MembersList
	}

	return membersList
}

func (members *Members) SetMembers(membersList []*Member) *Members {
	members.MembersList = membersList
	return members
}

func (members *Members) Open(filePath string) *Members {
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

func (members *Members) Close(filePath string) {

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

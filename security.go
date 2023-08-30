package main

import (
	"container/list"
	"encoding/json"
	"os"
)

// Structs declaration
type admin_list struct {
	admins []string `json:'admins'`
}

// Exported Functions

func ReadAdminList(adminlist list.List) bool {
	path := "./cache/adminlist.json"
	adminStruct := admin_list{}
	jsonString, err := os.ReadFile(path)
	if err != nil {

		return false
	}
	jsonErr := json.Unmarshal(jsonString, &adminStruct.admins)
	if jsonErr != nil {
		return false
	}
	for i := 0; i < len(adminStruct.admins); i++ {
		adminlist.PushBack(adminStruct.admins[i])
	}
	return true
}

func WriteAdminList(adminlist *list.List, response *string) {
	err := os.WriteFile("./cache/adminlist.json", []byte(createAdminJson(adminlist, response)), 0666)
	if err != nil {
		*response += "\nError creating adminlist.json"
	}
}

// Non-exported Functions

func createAdminJson(adminlist *list.List, response *string) string {
	adminstruct := admin_list{admins: nil}
	var admins []string
	i := adminlist.Front()
	for i != adminlist.Back().Next() {
		admins = append(admins, i.Value.(string))
	}
	adminstruct.admins = admins

	adminJson, err := json.Marshal(adminstruct)
	if err != nil {
		*response += "\nError creating adminlist.json text"
	}
	return string(adminJson)
}

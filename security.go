package main

import (
	"container/list"
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

// Structs declaration
type admin_list struct {
	admins []uuid.UUID `json:'admins'`
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

// Non-exported Functions

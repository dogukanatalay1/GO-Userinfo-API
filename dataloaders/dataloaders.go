package dataloaders

import (
	"encoding/json"
	m "../models"
	util "../utils"
)

func LoadUsers() []m.User {
  bytes, _ := util.ReadFile("../json/users.json")
  var users []m.User
  json.Unmarshal([]byte(bytes), &users)
  return users
}

func LoadInterests() []m.Interest {

}

func LoadInterestMappings() []m.InterestMapping {

}

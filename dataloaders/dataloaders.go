package dataloaders

import (
	"encoding/json"
	. "go-userinfo/models"
	util "go-userinfo/utils"
)

func LoadUsers() []User {
  bytes, _ := util.ReadFile("../json/users.json")
  var users []User
  json.Unmarshal([]byte(bytes), &users)
  return users
}

func LoadInterests() []Interest {
  bytes, _ := util.ReadFile("../json/interests.json")
  var interests []Interest
  json.Unmarshal([]byte(bytes), &interests)
  return interests
}

func LoadInterestMappings() []InterestMapping {
  bytes, _ := util.ReadFile("../json/userInterestMapping.json")
  var data []InterestMapping
  json.Unmarshal([]byte(bytes), &data)
  return data
}


package dataloaders

import (
	"encoding/json"
	"https://github.com/dogukanatalay1/GO-Userinfo-API/tree/master/Models"
	util "../utils"
)

func LoadUsers() []User {
  bytes, _ := util.ReadFile("../json/users.json")
  var users []User
  json.Unmarshal([]byte(bytes), &users)
  return users
}

func LoadInterests() []Interest {

}

func LoadInterestMappings() []InterestMapping {

}

package handlers

import (
	"encoding/json"
	"fmt"
	. "go-userinfo/dataloaders"
	. "go-userinfo/models"
	"net/http"
)

func Run() {
  http.HandleFunc("/", Handler)
  http.ListenAndServe(":8080", nil)	
}

func Handler(w http.ResponseWriter, r *http.Request) {
  // page object
  page := Page{ID: 7, Name: "Users", Description: "User List", URI: "/users"}
  // data loaders
  users := LoadUsers()
  interests := LoadInterests()
  interestMappings := LoadInterestMappings()
  // -- --
  var newUsers []User 

  for _ , user := range users {
    for _ , interestMapping := range interestMappings {
      if user.ID == interestMapping.UserID {
        for _ , interest := range interests {
          if interestMapping.InterestID == interest.ID {
            user.Interests = append(user.Interests, interest)
            // append returns the added version (like map in JS)
          }
        }
      }
    }
    newUsers = append(newUsers, user)
    fmt.Println(user.Firstname)
  }

  viewModel := UserViewModel{Page: page, Users: newUsers} 
  data, _ := json.Marshal(viewModel)
  w.Write([]byte(data))
}



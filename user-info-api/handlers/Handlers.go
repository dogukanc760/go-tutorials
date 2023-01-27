package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	d "../data-loaders"
	m "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	page := m.Page{
		ID:          7,
		Name:        "Kullanıcılar",
		Description: "Kullanıcı Listesi Verir",
		URI:         "/users",
	}

	// data loaders
	users := d.LoadUsers()
	interests := d.LoadInterest()
	interstMapping := d.LoadInsterestMappings()

	// işlem
	var newUsers []m.User

	for _, user := range users {
		for _, interestMapping := range interstMapping {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interest = append(user.Interest, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
		fmt.Println(newUsers[0].FirstName)
	}

	viewModel := m.UserViewModel{
		Page:  page,
		Users: newUsers,
	}

	data, _ := json.Marshal(viewModel)

	w.Write([]byte(data))
}

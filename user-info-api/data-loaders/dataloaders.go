package dataloaders

import (
	"encoding/json"

	. "../models"
	u "../utils"
)

func LoadUsers() []User {
	bytes, _ := u.ReadFile("../json/users.json")
	var users []User
	json.Unmarshal([]byte(bytes), &users)
	return users
}

func LoadInterest() []Interest {
	bytes, _ := u.ReadFile("../json/interests.json")
	var interest []Interest
	json.Unmarshal([]byte(bytes), &interest)
	return interest
}

func LoadInsterestMappings() []InterestMapping {
	bytes, _ := u.ReadFile("../json/userInterestMappings.json")
	var interestMappings []InterestMapping
	json.Unmarshal([]byte(bytes), &interestMappings)
	return interestMappings
}

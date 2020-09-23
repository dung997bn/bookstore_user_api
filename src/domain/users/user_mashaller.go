package users

import "encoding/json"

//PrivateUser struct
type PrivateUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//PublicUser struct
type PublicUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"`
}

//Marshall return user type
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
			LastName:    user.LastName,
			FirstName:   user.FirstName,
			Email:       user.Email,
		}
	}

	//if use json, 2 struct must be same json key
	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

//Marshall return users type
func (users Users) Marshall(isPublic bool) []interface{} {
	var result = make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.Marshall(isPublic)
	}
	return result
}

package api42

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	userListBase = "https://api.intra.42.fr/v2/campus/31/users?per_page=150&page=%d"
)

type UserProfileResponse struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PoolYear     string `json:"pool_year"`
	ProfileImage struct {
		Link string `json:"link"`
	} `json:"image"`

	Staff  bool `json:"staff?"`
	Active bool `json:"active?"`
}

type UserProfile struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	Image     string `json:"image"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	PoolYear  string `json:"pool_year"`
}

func RequestPageUser(token string, page int) (*resty.Response, error) {
	client := resty.New()

	userList := fmt.Sprintf(userListBase, page)

	return client.R().
		SetAuthToken(token).
		Get(userList)
}

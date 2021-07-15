package cafection

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type authRequest struct {
	AlternatePassword string `json:"alternatePassword"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Username          string `json:"username"`
}

// AuthResponse contains the results of an authentication with the cafection
// management API.
type AuthResponse struct {
	Success           bool   `json:"success"`
	ErrorMessage      string `json:"errorMessage"`
	Token             string `json:"token"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	AlternatePassword string `json:"alternatePassword"`
	DebugMode         bool   `json:"debugMode"`
	DefaultGroupID    int    `json:"defaultGroupID"`
	DefaultGroupName  string `json:"defaultGroupName"`
	CanCreateVirtual  bool   `json:"canCreateVirtual"`
	Username          string `json:"username"`
}

// Auth takes the username and password of an account, and makes an HTTP
// request to authenticate with the cafection API, and returns the result in the
// form of the AuthResponse.
func Auth(username string, password string) (*AuthResponse, error) {
	const url = "https://mngtool.cafection.com/mng3/api/Login/login"
	params := authRequest{
		AlternatePassword: "",
		Email:             username,
		Password:          password,
		Username:          username,
	}
	reqJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var authRes AuthResponse
	err = json.Unmarshal(body, &authRes)
	if err != nil {
		return nil, err
	}
	return &authRes, nil
}

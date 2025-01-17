package models

type Secret struct {
	Host     string `json:"host"`
	UserName string `json:"usrnm"`
	Password string `json:"psswrd"`
	Jwt      string `json:"jwtSign"`
	Database string `json:"database"`
}

package model

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

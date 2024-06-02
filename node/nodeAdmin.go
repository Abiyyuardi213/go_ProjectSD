package node

import "time"

type DataAdmin struct {
	Id       int
	Name     string
	Email    string
	Phone    string
	Username string
	Password string
}

type AdminLL struct {
	DBAdmin DataAdmin
	Next    *AdminLL
}

type LoginHistory struct {
	AdminID   int
	Username  string
	LoginTime time.Time
	Next *LoginHistory
}
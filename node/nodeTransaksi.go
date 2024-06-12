package node

import "time"

type TransaksiKeluar struct {
	AdminID    int
	Username   string
	NamaBarang string
	Quantity   int
	Time       time.Time
	Next *TransaksiKeluar
}
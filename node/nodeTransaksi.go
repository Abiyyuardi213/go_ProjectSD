package node

import "time"

type BarangTransaksi struct {
	SerialNumber int
	Quantity     int
}

type Transaksi struct {
	ID            int
	AdminID int
	Username string
	Barang        []BarangTransaksi
	TransaksiTime time.Time
	Next *Transaksi
}
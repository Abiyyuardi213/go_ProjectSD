package node

type DataBarang struct {
	SerialNumber int
	Name string
	Price string
	Stock int
}

type BarangLL struct {
	DBBarang DataBarang
	Next *BarangLL
}
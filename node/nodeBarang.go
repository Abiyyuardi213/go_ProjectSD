package node

type DataBarang struct {
	SerialNumber int
	Name string
	Stock string
	Price string
}

type BarangLL struct {
	DBBarang DataBarang
	Next *BarangLL
}
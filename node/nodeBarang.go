package node

type DataBarang struct {
	SerialNumber int
	Name string
	Stock int
	Price float64
}

type BarangLL struct {
	DBBarang DataBarang
	Next *BarangLL
}
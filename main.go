package main

import "fmt"

type Cart struct {
	item map[string]int
}

func newCart() Cart {
	return Cart{
		item: map[string]int{},
	}
}

func (cart *Cart) tambahProduk(kodeProduk string, kuantitas int) {
	quantity, exist := cart.item[kodeProduk]

	if exist {
		kuantitas += quantity
	}

	cart.item[kodeProduk] = kuantitas
}

func (cart *Cart) hapusProduk(kodeProduk string) {
	delete(cart.item, kodeProduk)
}

func (cart *Cart) tampilkanCart() {
	for index, res := range cart.item {
		fmt.Println(index, res)
	}
}
func main() {
	cart := newCart()

	cart.tambahProduk("pisang hijau", 2)
	cart.tambahProduk("semangka kuning", 3)
	cart.tambahProduk("apel merah", 4)
	cart.tambahProduk("apel merah", 1)
	cart.tambahProduk("apel merah", 2)

	cart.hapusProduk("semangka kuning")
	cart.hapusProduk("semangka merah")

	cart.tampilkanCart()

}

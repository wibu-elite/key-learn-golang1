package main

import "fmt"

var (
	namaGroup  string
	alamat     string
	nomorTlp   uint
	lamaRent   uint
	totalBayar uint
)

func main() {

	fmt.Println(`                 Studio Al Izzah
	================================
	Isi Daftar Penyewa Alat Berikut `)

	fmt.Print("Nama Group :")
	fmt.Scan(&namaGroup)

	fmt.Print("Alamat :")
	fmt.Scan(&alamat)

	fmt.Print("Nomor Telepon :")
	fmt.Scan(&nomorTlp)

	fmt.Print("Lama Rental :")
	fmt.Scan(&lamaRent)

	fmt.Print("Total Bayar :")
	totalBayar = sumRent(lamaRent)
	fmt.Println(totalBayar)

}

func sumRent(lamaRent uint) uint {

	if lamaRent == 1 {
		totalBayar = 1 * 200000
		return totalBayar
	} else {
		totalBayar = 200000 + (lamaRent-1)*90000
		return totalBayar
	}
}

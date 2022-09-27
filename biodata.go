package main

import (
	"fmt"
)

type Siswa struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var absen int
	fmt.Printf("Masukan Absen Siswa : ")
	_, err := fmt.Scan(&absen)
	if err != nil {
		fmt.Println(err)
	}
	var data_siswa1 = Siswa{
		nama:      "M Naufa Dzulfiqar",
		alamat:    "Jalan Pojok tengah Cimahi",
		pekerjaan: "Mahasiswa",
		alasan:    "Karena Golang sangat ringan",
	}
	var data_siswa2 = Siswa{
		nama:      "ujang",
		alamat:    "Jakarta",
		pekerjaan: "Mahasiswa",
		alasan:    "Karena Golang sangat Keren",
	}
	var data_siswa3 = Siswa{
		nama:      "Dadang",
		alamat:    "Bandung",
		pekerjaan: "Wirausaha",
		alasan:    "Karena Golang sangat Hebat",
	}
	var data_siswa4 = Siswa{
		nama:      "Udin",
		alamat:    "Sumedang",
		pekerjaan: "Developer",
		alasan:    "Karena Golang Dari google",
	}
	var data_siswa5 = Siswa{
		nama:      "solihin",
		alamat:    "papua",
		pekerjaan: "Developer",
		alasan:    "Karena Golang Dari mantap",
	}
	if absen == 1 {
		fmt.Printf("Data Absen 1 adalah \n")
		fmt.Printf("Nama : %+v\n ", data_siswa1.nama)
		fmt.Printf("Alamat : %+v\n ", data_siswa1.alamat)
		fmt.Printf("pekerjaan : %+v\n ", data_siswa1.pekerjaan)
		fmt.Printf("alasan : %+v\n ", data_siswa1.alasan)
	} else if absen == 2 {
		fmt.Printf("Data Absen 2 adalah \n")
		fmt.Printf("Nama : %+v\n ", data_siswa2.nama)
		fmt.Printf("Alamat : %+v\n ", data_siswa2.alamat)
		fmt.Printf("pekerjaan : %+v\n ", data_siswa2.pekerjaan)
		fmt.Printf("alasan : %+v\n ", data_siswa2.alasan)
	} else if absen == 3 {
		fmt.Printf("Data Absen 3 adalah \n")
		fmt.Printf("Nama : %+v\n ", data_siswa3.nama)
		fmt.Printf("Alamat : %+v\n ", data_siswa3.alamat)
		fmt.Printf("pekerjaan : %+v\n ", data_siswa3.pekerjaan)
		fmt.Printf("alasan : %+v\n ", data_siswa3.alasan)
	} else if absen == 4 {
		fmt.Printf("Data Absen 4 adalah \n")
		fmt.Printf("Nama : %+v\n ", data_siswa4.nama)
		fmt.Printf("Alamat : %+v\n ", data_siswa4.alamat)
		fmt.Printf("pekerjaan : %+v\n ", data_siswa4.pekerjaan)
		fmt.Printf("alasan : %+v\n ", data_siswa4.alasan)
	} else {
		fmt.Printf("Data Absen 5 adalah \n")
		fmt.Printf("Nama : %+v\n ", data_siswa5.nama)
		fmt.Printf("Alamat : %+v\n ", data_siswa5.alamat)
		fmt.Printf("pekerjaan : %+v\n ", data_siswa5.pekerjaan)
		fmt.Printf("alasan : %+v\n ", data_siswa5.alasan)
	}
}

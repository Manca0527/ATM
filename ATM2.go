package main

import "fmt"

type Akun struct {
	Akun  string
	Pin   int
	saldo float64
}

func main() {
	dataAkun := []Akun{
		{"Ilyas", 1111, 200000},
		{"Ali", 2222, 250000},
	}

	fmt.Println("=== Selamat Datang di ATM ===")

	var Login *Akun
	var namaAkun string
	var pin int

	fmt.Println("\n=== Login ATM ===")
	fmt.Print("Masukan Nama Akun: ")
	fmt.Scanln(&namaAkun)

	fmt.Print("Masukan Pin: ")
	fmt.Scanln(&pin)

	loginBerhasil := false
	for i := range dataAkun {
		if dataAkun[i].Akun == namaAkun && dataAkun[i].Pin == pin {
			Login = &dataAkun[i]
			loginBerhasil = true
			break
		}
	}

	if loginBerhasil {
		fmt.Println("Login Berhasil")

		isRunning := true

		for isRunning {
			var Pilihan int

			fmt.Println("\n=== Menu ATM ===")
			fmt.Println("1. Setor Tunai")
			fmt.Println("2. Tarik Tunai")
			fmt.Println("3. Transfer")
			fmt.Println("4. Cek Saldo")
			fmt.Println("5. Exit")
			fmt.Print("Pilih Menu (1-5): ")
			fmt.Scanln(&Pilihan)

			if Pilihan == 1 {
				var jumlah float64
				fmt.Print("Masukan Jumlah Setor Tunai: Rp")
				fmt.Scanln(&jumlah)

				jumlahValid := jumlah > 0
				if !jumlahValid {
					fmt.Println("Jumlah Tidak Valid!!!")
				} else {
					Login.saldo += jumlah
					fmt.Printf("Setoran berhasil! Saldo Baru: Rp%.2f\n", Login.saldo)
				}

			} else if Pilihan == 2 {
				var jumlah float64
				fmt.Print("Masukan Jumlah Tarik Tunai: Rp")
				fmt.Scanln(&jumlah)

				jumlahValid := jumlah > 0
				saldoCukup := jumlah <= Login.saldo

				if !jumlahValid {
					fmt.Println("Jumlah Tidak Valid!!!")
				} else if !saldoCukup {
					fmt.Println("Saldo Tidak Cukup!!!")
				} else {
					Login.saldo -= jumlah
					fmt.Printf("Penarikan Berhasil! Saldo Baru: Rp%.2f\n", Login.saldo)
				}

			} else if Pilihan == 3 {
				var NamaAkun string
				var jumlah float64

				fmt.Print("Masukan Nama Akun Tujuan: ")
				fmt.Scanln(&NamaAkun)

				fmt.Print("Masukan Jumlah Transfer: Rp")
				fmt.Scanln(&jumlah)

				var akunTujuan *Akun
				akunDitemukan := false
				for i := range dataAkun {
					if dataAkun[i].Akun == NamaAkun {
						akunTujuan = &dataAkun[i]
						akunDitemukan = true
						break
					}
				}

				jumlahValid := jumlah > 0
				saldoCukup := jumlah <= Login.saldo

				if !akunDitemukan {
					fmt.Println("Akun Tujuan Tidak Ditemukan!!!")
				} else if !jumlahValid {
					fmt.Println("Jumlah Tidak Valid!!!")
				} else if !saldoCukup {
					fmt.Println("Saldo Tidak Cukup!!!")
				} else {
					Login.saldo -= jumlah
					akunTujuan.saldo += jumlah
					fmt.Printf("Transfer Berhasil! Saldo Baru: Rp%.2f\n", Login.saldo)
				}

			} else if Pilihan == 4 {
				fmt.Printf("Saldo Anda: Rp%.2f\n", Login.saldo)

			} else if Pilihan == 5 {
				isRunning = false
				fmt.Println("\nLogout berhasil!")
				fmt.Println("\n=================================")
				fmt.Println("   TERIMA KASIH TELAH GUNAKAN   ")
				fmt.Println("           ATM KAMI             ")
				fmt.Println("=================================")

			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		}

	} else {
		fmt.Println("Login Gagal!!!")
	}
}

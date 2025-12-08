package main

import "fmt"

type Akun struct {
	NoRekening int
	Pin        int
	Saldo      float64
}

func main() {
	dataAkun := []Akun{
		{123456, 1111, 1000000},
		{234567, 2222, 2500000},
		{345678, 3333, 500000},
		{456789, 4444, 750000},
		{567890, 5555, 3000000},
	}

	fmt.Println("=================================")
	fmt.Println("   SELAMAT DATANG DI ATM BANK   ")
	fmt.Println("=================================")

	var akunLogin *Akun
	maxPercobaan := 3
	percobaan := 0

	for percobaan < maxPercobaan {
		var noRek, pin int

		fmt.Println("\n=== LOGIN ATM ===")
		fmt.Print("Masukkan Nomor Rekening: ")
		fmt.Scan(&noRek)

		fmt.Print("Masukkan PIN: ")
		fmt.Scan(&pin)

		found := false
		for i := range dataAkun {
			if dataAkun[i].NoRekening == noRek && dataAkun[i].Pin == pin {
				akunLogin = &dataAkun[i]
				found = true
				fmt.Println("Login berhasil!")
				break
			}
		}

		if found {
			break
		}

		percobaan++
		fmt.Printf("PIN salah! Sisa percobaan: %d\n", maxPercobaan-percobaan)
	}

	if akunLogin == nil {
		fmt.Println("\n*** KARTU DIBLOKIR ***")
		fmt.Println("\n=================================")
		fmt.Println("   TERIMA KASIH TELAH GUNAKAN   ")
		fmt.Println("           ATM KAMI             ")
		fmt.Println("=================================")
		return
	}

	for {
		var pilihan int

		fmt.Println("\n=================================")
		fmt.Println("         MENU UTAMA ATM          ")
		fmt.Println("=================================")
		fmt.Println("1. Info Saldo")
		fmt.Println("2. Setor Tunai")
		fmt.Println("3. Tarik Tunai")
		fmt.Println("4. Transfer")
		fmt.Println("5. Logout")
		fmt.Println("=================================")
		fmt.Print("Pilih menu (1-5): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Printf("\nSaldo Anda: Rp %.2f\n", akunLogin.Saldo)

		} else if pilihan == 2 {
			var jumlah float64
			fmt.Print("\nMasukkan Jumlah Setoran: Rp ")
			fmt.Scan(&jumlah)

			if jumlah <= 0 {
				fmt.Println("Jumlah tidak valid!")
			} else {
				akunLogin.Saldo += jumlah
				fmt.Printf("Setoran berhasil! Saldo baru: Rp %.2f\n", akunLogin.Saldo)
			}

		} else if pilihan == 3 {
			var jumlah float64
			fmt.Print("\nMasukkan Jumlah Penarikan: Rp ")
			fmt.Scan(&jumlah)

			if jumlah <= 0 {
				fmt.Println("Jumlah tidak valid!")
			} else if jumlah > akunLogin.Saldo {
				fmt.Println("Saldo tidak mencukupi!")
			} else {
				akunLogin.Saldo -= jumlah
				fmt.Printf("Penarikan berhasil! Saldo baru: Rp %.2f\n", akunLogin.Saldo)
			}

		} else if pilihan == 4 {
			var noRekTujuan int
			var jumlah float64

			fmt.Print("\nMasukkan Nomor Rekening Tujuan: ")
			fmt.Scan(&noRekTujuan)

			fmt.Print("Masukkan Jumlah Transfer: Rp ")
			fmt.Scan(&jumlah)

			var akunTujuan *Akun
			for i := range dataAkun {
				if dataAkun[i].NoRekening == noRekTujuan {
					akunTujuan = &dataAkun[i]
					break
				}
			}

			if akunTujuan == nil {
				fmt.Println("Rekening tujuan tidak ditemukan!")
			} else if akunTujuan.NoRekening == akunLogin.NoRekening {
				fmt.Println("Tidak dapat transfer ke rekening sendiri!")
			} else if jumlah <= 0 {
				fmt.Println("Jumlah transfer harus lebih dari 0!")
			} else if jumlah > akunLogin.Saldo {
				fmt.Println("Saldo tidak mencukupi!")
			} else {
				akunLogin.Saldo -= jumlah
				akunTujuan.Saldo += jumlah
				fmt.Printf("\nTransfer berhasil!\n")
				fmt.Printf("Saldo Anda sekarang: Rp %.2f\n", akunLogin.Saldo)
			}

		} else if pilihan == 5 {
			fmt.Println("\nLogout berhasil!")
			fmt.Println("\n=================================")
			fmt.Println("   TERIMA KASIH TELAH GUNAKAN   ")
			fmt.Println("           ATM KAMI             ")
			fmt.Println("=================================")
			return

		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

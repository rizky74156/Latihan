package main

import (
	"fmt"
)

type login struct{
	username string
	password string
	balance int
}

var akun login

var data = []login{
	{"rizky", "123", 0},
	{"ackerman", "123", 100000},
}

func input(message string)string{
	var input string
	fmt.Print(message)
	fmt.Scan(&input)
	return input
}

func inputAngka(massage string)int{
    var input int
    fmt.Print(massage)
    fmt.Scan(&input)
    return input
}

func login_page(){
    fmt.Println("Anda Memilih 1")
    fmt.Println("---Silahkan Isi Pendaftaran---")
    if masuk(input("Masukkan Username : "),input("Masukkan Password : ")) {
        fmt.Println("Login Berhasil")
        main2()
    }else{
        fmt.Println("Username dan password tidak ditemukan")
    }
}

func masuk(user, pass string)bool{
	for _, dataDaftar := range data{
		if dataDaftar.username == user && dataDaftar.password == pass{
			akun = dataDaftar
			return true
		}
	}
	return false
}

func cekUser(user, pass string)bool{
	for _, dataDaftar := range data{
		if dataDaftar.username == user && dataDaftar.password == pass{
			return true
		}
	}
	return false
}

func batasUser1(user string)bool{
	return len(user) > 3
}

func batasUser2(user string)bool{
	return len(user) < 10
}

func cekRegister(user string, newUser login){
	if batasUser1(user) && batasUser2(user){
		data = append(data, newUser)
		fmt.Println(data)
	}else if !batasUser1(user){
		fmt.Println("Panjang username harus lebih dari 3")
		pageRegister()
	}else if !batasUser2(user){
		fmt.Println("Panjang username harus kurang dari 10")
		pageRegister()
	}
}

func register(user, pass string){
	var newUser login
	var pilihan int
		newUser.username = user
		newUser.password = pass
			fmt.Print("Apakah anda sudah isi data dengan benar? [1 untuk benar, 0 untuk ulangi]:")
			fmt.Scan(&pilihan)
				if pilihan == 1{
					if !cekUser(user, pass){
						cekRegister(user,newUser)
					}else{
						fmt.Println("Username sudah tersedia")
					}
				}else if pilihan == 0{
					pageRegister()
				}
}

func pageRegister(){
	fmt.Println("==== Silakan Isi Pendaftaran =====")
	register(input("Masukan Username :"), input("Masukan Password :"))
}

func cekData([]login){
	fmt.Println(data)
}

func deposit(saldo int)int{
	for i := range data{
		if data[i].username == akun.username{
			if saldo > 0{
			data[i].balance += saldo
			akun.balance = data[i].balance
			break
			}else if saldo <=0{
				fmt.Println("Harus nominal positif")
			}
		}
	}
	return akun.balance
}

func depositPage(){
    fmt.Println("Anda Memilih 1")
    fmt.Println("---Silahkan Isi Saldo Anda---")
    deposit(inputAngka("Masukkan Jumlah Saldo Anda : "))

}

func withdraw(saldo int)int{
	for i := range data{
		if data[i].username == akun.username{
			if cekWithdraw(saldo){
				data[i].balance -= saldo
				akun.balance = data[i].balance
				break
			}
		}
	}
	return akun.balance
}

func cekWithdraw(saldo int)bool{
	for i := range data{
		if data[i].balance >= saldo{
			return true
		}else if data[i].balance == 0{
			fmt.Println("Saldo anda tidak ada ")
		}else if data[i].balance <= saldo{
			fmt.Println("Saldo anda kurang")
		}
	}
	return false
}

func withdrawPage(){
	fmt.Println("Anda Memilih 2")
    fmt.Println("---Silahkan Isi Saldo Yang Akan Anda Ambil---")
    withdraw(inputAngka("Masukkan Jumlah Saldo Yang Akan Anda Ambil : "))
}


func pageUbahUsername(){
    fmt.Println("---Ubah Username---")
    ubahUsername(input("Masukkan username yang baru : "))
}

func ubahUsername(username string){
	for i := range data{ 	
		if data[i].username == akun.username{
			if cekUbahUsername1(username){
				data[i].username = username
				akun.username = data[i].username
				break
			}
		}
	}
}

func cekUbahUsername1(user string)bool{
	if !cekUbahUsername2(user){
		if batasUser1(user) && batasUser2(user){
			fmt.Println("Username berhasil diubah")
			return true
		}else if !batasUser1(user){
			fmt.Println("Panjang username harus lebih dari 3")
			pageUbahUsername()
		}else if !batasUser2(user){
			fmt.Println("Panjang username harus kurang dari 10")
			pageUbahUsername()
		}
	}else{
		fmt.Println("Username sudah ada")
	}
	return false
}

func cekUbahUsername2(user string)bool{
	for _, dataDaftar := range data{
		if dataDaftar.username == user{
			return true
		}
	}
	return false
}

func pageUbahPassword(){
    fmt.Println("---Ubah Password---")
    ubahPassword(input("Masukkan password yang baru : "))
}

func ubahPassword(pass string){
	for i := range data{
		if data[i].username == akun.username{
			data[i].password = pass
			akun.password = data[i].password
			break
		}
	}
}

func pageTransfer(){
	fmt.Println("==== Silakan Isi ====")
	transfer(input("Masukan Username yang akan ditransfer : "), 
	inputAngka("Saldo yang akan diransfer : "))
}

// func transfer(user string, saldo int){
// 	for i := range data{
// 		if data[i].username == akun.username{
// 			cekTransfer(user,saldo)
// 			data[i].username == user
// 			data[i].balance += saldo
// 		}else if data[i].username == akun.username{
// 			fmt.Println("Tidak bisa transfer ke akun yang sama ")
// 		}else{
// 			fmt.Println("Tidak ada data")
// 		}
// 	}
// }

func transfer(user string, saldo int){
	var biaya int = (saldo*3/100)
	for i:= range data{
		for data[i].username == akun.username{
			if transferAkun1(user,saldo) && transferAkun2(user,saldo){
				data[i].balance -= (saldo + biaya)
				akun.balance = data[i].balance
				fmt.Println("Transfer Berhasil, Dengan biaya admin sebesar : ",biaya )
				break
			}else {
				fmt.Println("data tidak ditemukan")
				break
			}
		}
	}
}

func transferAkun1(user string, saldo int)bool{
	for i := range data{
		for data[i].username == akun.username{
			if data[i].balance > saldo && saldo > 0{
				return true
			}else if data[i].balance == saldo{
				fmt.Println("Saldo tidak mencukupi untuk biaya admin")
				
			}else if data[i].balance <= saldo{
				fmt.Println("Saldo anda kurang")
				
			}else if saldo < 0{
				fmt.Println("Tidak bisa memasukan angka negatif")
				
			}
			break
		}
	}
	return false
}

func transferAkun2(user string, saldo int)bool{
	for i := range data{
		if data[i].username == user{
			data[i].balance += saldo
			return true
		}else if akun.username == user{
			fmt.Println("Tidak bisa transfer uang ke akun yang sama")
			break
		}
	}
	return false
}

func main(){
	var menuu int
	var statuslooping bool = true

	for statuslooping {
	menu()
	fmt.Print("Masukan Pilihan Anda : ")
	fmt.Scan(&menuu)
	if menuu == 1{
		login_page()
	}else if menuu == 2{
		fmt.Println("Anda Memilih 2")
		pageRegister()
	}else if menuu == 3{
		cekData(data)
	}else if menuu == 4{
		fmt.Println("Terima Kasih")
		break
	}
	}
}

func main2(){
	var menuu int
	var statuslooping bool = true

	for statuslooping {
	menu2()
	fmt.Print("Masukan Pilihan Anda : ")
	fmt.Scan(&menuu)
	if menuu == 1{
		depositPage()
		fmt.Println("Saldo anda :",akun.balance)
	}else if menuu == 2{
		withdrawPage()
		fmt.Println("Sisa saldo anda :", akun.balance)
	}else if menuu == 3{
		fmt.Println("Jumlah saldo anda :",akun.balance)
	}else if menuu == 4{
		main3()
	}else if menuu == 5{
		pageTransfer()
	}else if menuu == 6{
		fmt.Println("Anda berhasil logout")
		break
	}
	}
}

func main3(){
	var menuu int
	var statuslooping bool = true

	for statuslooping {
	menu3()
	fmt.Print("Masukan Pilihan Anda : ")
	fmt.Scan(&menuu)
	if menuu == 1{
		fmt.Println("Anda Memilih 1")
		pageUbahUsername()
	}else if menuu == 2{
		fmt.Println("Anda Memilih 2")
		pageUbahPassword()
		fmt.Println("Password berhasil diubah")
	}else if menuu == 3{
		fmt.Println("Username anda :",akun.username)
		fmt.Println("Password anda :",akun.password)
	}else if menuu == 4{
		fmt.Println("Anda berhasil keluar")
		break
	}
	}
}

func menu(){
	fmt.Println("Selamat Datang Di Aplikasi Menghitung")
	fmt.Println("1. Login")
	fmt.Println("2. Daftar")
	fmt.Println("3. Cek Data")
	fmt.Println("4. Keluar")
}

func menu2(){
	fmt.Println("Username anda :",akun.username)
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Akun")
	fmt.Println("5. Transfer")
	fmt.Println("6. Keluar")
}

func menu3(){
	fmt.Println("=====Akun=====")
	fmt.Println("1. Ubah username")
	fmt.Println("2. Ubah password")
	fmt.Println("3. Cek akun")
	fmt.Println("4. Keluar")
}
package main
import "fmt"

func main(){
	var choose,uang,saldo int
	var lagi string

	saldo=0
	menu()
	
	fmt.Print("Pilihan : ")
	fmt.Scan(&choose)
	fmt.Println(" ")


	if choose==1 {
		satu(uang,&saldo)

	}else if choose==2 {
		dua(uang,&saldo)

	}else if choose==3 {
		tiga(&saldo)

	}else {

	}
	
	fmt.Print("Kembali lagi ke menu ? (Y/T) : ")
	fmt.Scan(&lagi)
	fmt.Println(" ")

	for lagi=="Y" || lagi=="y" {
		menu()
	
		fmt.Print("Pilihan : ")
		fmt.Scan(&choose)
		fmt.Println(" ")


		if choose==1 {
			satu(uang,&saldo)
				
		}else if choose==2 {
			dua(uang,&saldo)
		
		}else if choose==3 {
			tiga(&saldo)

		}else {

		}
		
		fmt.Print("Kembali lagi ke menu ? (Y/T) : ")
		fmt.Scan(&lagi)
		fmt.Println(" ")
		
	}

	fmt.Println("Transaksi Selesai. \n")

}


func menu(){
	fmt.Println("Selamat Datang di Tabungan Online")
	fmt.Println("Pilih menu :")
	fmt.Println("1.Nabung")
	fmt.Println("2.Ambil")
	fmt.Println("3.Cek Saldo")
}

func satu(uang int, saldo *int) {
	fmt.Print("Masukan uang : ")
	fmt.Scan(&uang)
	
	fmt.Println("Berhasil")
	*saldo=*saldo+uang
	fmt.Println("Saldo : ",*saldo)
	
}

func dua(tarik int, saldo *int ){
	
	fmt.Println("Saldo awal : ",*saldo)
	fmt.Print("Jumlah ditarik : ")
	fmt.Scan(&tarik)
	if tarik<=*saldo {
		*saldo=*saldo-tarik
		fmt.Println("Berhasil.")
		fmt.Println("Saldo : ",*saldo)
	}else if tarik>*saldo {
		fmt.Println("Gagal.")
		fmt.Println("Saldo : ",*saldo)
	}
}

func tiga(saldo *int) {
	fmt.Println("Saldo : ",*saldo)
}


// defer digunakan untuk menjalankan perintah atau fungsi di akhir, setelah sebuah function selesai
// panic untuk throw error dan memberhentikan function
// recover digunakan untuk menghentikan panic dan mengambil data dari panic sehingga function dapat kembali berjalan

package main

import "fmt"

func sayHelloOk(name string) {
	if name == "Hello" {
		panic("name tidak boleh Hello")
	}
	catch := recover()
	if catch != nil {
		fmt.Println("ERROR", catch)
	}
	fmt.Println("Hello", name)
}

func main() {
	name := "Sahal"
	fmt.Println("Mulai")

	defer sayHelloOk(name)
	sayHelloOk("Hello Hi")
	fmt.Println("Selamat pagi")
}

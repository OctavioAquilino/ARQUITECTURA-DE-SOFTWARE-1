package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	path = "C:/Users/tayia/OneDrive/Documentos/UNIVERSIDAD/3 año/ARQUI DE SOFTWARE/Ejercicios practicos/prueba.txt"
	//f, _ := os.Create(path)
	f, _ := os.OpenFile(path, os.O_APPEND, 0666)
	fmt.Fprintln(f, "high")
	fmt.Fprintln(f, "holaaaaaa")
	fmt.Fprintln(f, "chau")
	fmt.Fprintln(f, "eeeeee")
	_ = f.Close()

	fmt.Println("Leyendoooo")

	x, _ := os.ReadFile(path)
	fmt.Print(string(x))
	//	_ = f.Close()
}

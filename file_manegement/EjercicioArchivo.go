package main

import (
	"fmt"
	"os"
)

func main() {
	f := CreateFile("./ArchivoNuevo")
	defer CloseFile(f)
	Escribir("Hola escribiendo skrrrr", f)
	Escribir("Hola escribiendo skrrrr", f)
	Escribir("Hola escribiendo skrrrr", f)

	x, _ := os.ReadFile("./ArchivoNuevo")
	fmt.Print(string(x))
	
	data := make([]byte, 100)
	count, _ := f.Read(data)

	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

func CreateFile(path string) *os.File {
	f, _ := os.Create(path)
	return f
}

func Escribir(texto string, f *os.File) {
	fmt.Fprintln(f, texto)
}

func CloseFile(f *os.File) {
	f.Close()
}

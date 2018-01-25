package main

import "fmt"

//Declarar un nuevo tipo de valor, que funcionarán como modelo
type Gorra struct {
	marca string
	color string
	precio float32
	plana bool
}

func main() {
	var gorra_negra = Gorra{
		marca: "Nike",
		color: "Negra",
		precio: 27.7,
		plana: false}

	fmt.Println(gorra_negra)

	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Jesus Lopez"
	var numero float32 = 1.222222
	var b bool = false
	//Otra forma de declarar variables
	pais := "Venezuela"

	const year int = 2018 //no se puede modificar
	fmt.Println("Hola mundo " + nombre)
	fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println(numero)
	fmt.Println(b)
	fmt.Println(pais)
	fmt.Println(year)
	/*
		Comandos de ayuda para go
		godoc //para ver la documentación de go
		go build archivo.go //para construir el .exe
		go fmt archivo.go //para indexar el código
	*/
}
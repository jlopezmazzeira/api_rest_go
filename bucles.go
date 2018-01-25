package main

//La librería "os" es para poder interactuar con la consola
//La librería "strconv" es para poder convertir un string a entero
import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	
	fmt.Println("Hola " + os.Args[1] + " Bienvenido al programa con Go")
	edad, _ := strconv.Atoi(os.Args[2])

	if (edad >= 18) && edad != 25 && edad != 99{
		fmt.Println("Eres mayor de edad")
	} else if edad == 25 {
		fmt.Println("Tienes 25 años")
	} else if edad == 99 {
		fmt.Println("Pronto morirás")
	} else {
		fmt.Println("Eres menor de edad")
	}


	if edad%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	tope := 26

	for i := 1; i <= tope; i++ {
		if i%2 == 0 {
			fmt.Println("El número es par: ", i)
		} else {
			fmt.Println("El número es impar: ", i)
		}
	}

	peliculas4 := []string{
		"Hola",
		"Mundo",
		"Go"}

	for i := 0; i < len(peliculas4); i++ {
		if i%2 == 0 {
			fmt.Println("La película "+peliculas4[i]+" es par: ", i)
		} else {
			fmt.Println("La película "+peliculas4[i]+" es impar: ", i)
		}
	}

	//Simulando un foreach
	for _, pelicula := range peliculas4 {
		fmt.Println(pelicula)
	}

	momento := time.Now()
	hoy := momento.Weekday()

	switch hoy {
		case 0:
			fmt.Println("Hoy es Domingo")
		case 1:
			fmt.Println("Hoy es Lunes")
		case 2:
			fmt.Println("Hoy es Martes")
		case 3:
			fmt.Println("Hoy es Miércoles")
		case 4:
			fmt.Println("Hoy es Jueves")
		case 5:
			fmt.Println("Hoy es Viernes")
		case 6:
			fmt.Println("Hoy es Sabádo")
		default:
			fmt.Println("Este día no existe")
	}
}

package main

import "fmt"

func main() {
	//Array 1
	var peliculas [3]string;
	peliculas[0] = "Hola";
	peliculas[1] = "Mundo";
	peliculas[2] = "Go";
	fmt.Println(peliculas);

	//Array 2
	peliculas2 := [3]string{
		"Hola",
		"Mundo",
		"Go"}
	fmt.Println(peliculas2);

	//Array 3
	var peliculas3 [3][2]string;
	peliculas3[0][0] = "Hola";
	peliculas3[0][1] = "Hola";
	peliculas3[1][0] = "Mundo";
	peliculas3[1][1] = "Mundo";
	peliculas3[2][0] = "Go";
	peliculas3[2][1] = "Go";
	fmt.Println(peliculas3);

	//Slice es un array dinámico.
	peliculas4 := []string{
		"Hola",
		"Mundo",
		"Go"}

	//Forma de añadir elementos a un array dinámico
		peliculas4 = append(peliculas4, "Programacion")

	//Tamaño de un slice
		fmt.Println(len(peliculas4));

	//Filtrar elementos
		fmt.Println(peliculas4[0:2]);

	fmt.Println(peliculas4);
}

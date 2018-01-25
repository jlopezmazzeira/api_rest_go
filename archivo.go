package main

//La librería "ioutil" permite leer archivos
import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Bienvenido al programa")

	//nuevo_texto := []byte(os.Args[1])
	nuevo_texto := os.Args[1]+"\n"
	//fmt.Println(string(nuevo_texto))
	
	//escribir := ioutil.WriteFile("fichero.txt", nuevo_texto, 0777)
	//showError(escribir)
	fichero, err := os.OpenFile("fichero.txt", os.O_APPEND|os.O_WRONLY, 0777)
	showError(err)

	escribir, err := fichero.WriteString(nuevo_texto)
	fmt.Println(escribir)
	showError(err)
	
	fichero.Close()

	texto, errorDeFichero := ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)

	fmt.Println(string(texto))
}

func showError(e error){
	if(e != nil){
		//Es equivalente al die de PHP, permite cortar la ejecución del programa si existe un error
		panic(e)
	}
}
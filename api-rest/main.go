package main

//La librería "http" permite crear nuestro servidor web
import (
	"net/http"
	"log"
)

func main() {
	router := NewRouter()

	/*
	Si no usamos la librería gorilla mux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	})*/

	//server := http.ListenAndServe(":8080", nil)
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

	/*Para instalar paquetes en Go debemos tener configurada las variables de entorno
	por eso debemos hacer lo siguiente:
		export GOPATH=$HOME/go

	ya con eso podemos instalar librerías
		go get -u github.com/gorilla/mux --> para crear rutas
		go get gopkg.in/mgo.v2 --> mongoDB
		go get gopkg.in/mgo.v2/bson
	*/
}
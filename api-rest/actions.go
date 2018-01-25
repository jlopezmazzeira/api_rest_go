package main

//La librería "http" permite crear nuestro servidor web
//La librería "github.com/gorilla/mux" nos permite crear un sistema de rutas
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*var movies = Movies{
		Movie{"Pelicula 1", 2000, "Desconocido"},
		Movie{"Pelicula 2", 2005, "Desconocido 0"},
		Movie{"Pelicula 3", 2002, "Desconocido 1"},
}*/

func getSession() *mgo.Session{
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

var collection = getSession().DB("curso_go").C("movies")

func responseMovie(status int, w http.ResponseWriter, result Movie) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

func responseMovies(status int, w http.ResponseWriter, results []Movie) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	var result []Movie
	//err := collection.Find(nil).All(&result)
	err := collection.Find(nil).Sort("-_id").All(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Result: ", result)
	}

	responseMovies(200, w, result)
	//fmt.Fprintf(w, "Listado de películas")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	//Para poder obtener los paramétros que han sido enviado por URL
	params := mux.Vars(r)

	//Forma de sacar los paramétros
	movie_id := params["id"]

	//Al colocar %s, le debemos pasar un segundo paramétro, para que lo sustituya
	//fmt.Fprintf(w, "Haz cargado la película número %s", movie_id)

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	result := Movie{}

	err:= collection.FindId(oid).One(&result)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(200, w, result)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie

	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movie_data)
	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseMovie(200, w, movie_data)
	//movies = append(movies, movie_data)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	//Para poder obtener los paramétros que han sido enviado por URL
	params := mux.Vars(r)

	//Forma de sacar los paramétros
	movie_id := params["id"]

	//Al colocar %s, le debemos pasar un segundo paramétro, para que lo sustituya
	//fmt.Fprintf(w, "Haz cargado la película número %s", movie_id)

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	decoder := json.NewDecoder(r.Body)
	
	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}

	err = collection.Update(document, change)

	if err != nil {
		panic(err)
		w.WriteHeader(404)
		return
	}

	responseMovie(200, w, movie_data)
}

type Message struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}

func MovieRemove(w http.ResponseWriter, r *http.Request) {
	//Para poder obtener los paramétros que han sido enviado por URL
	params := mux.Vars(r)

	//Forma de sacar los paramétros
	movie_id := params["id"]

	//Al colocar %s, le debemos pasar un segundo paramétro, para que lo sustituya
	//fmt.Fprintf(w, "Haz cargado la película número %s", movie_id)

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err:= collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	//result := Message{"success", "La película con el ID " +movie_id+ " ha sido eliminada correctamente"}
	
	message := new(Message)
	message.setStatus("success")
	message.setMessage("La película con el ID " +movie_id+ " ha sido eliminada correctamente")

	result := message
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}
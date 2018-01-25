package main

import "fmt"

func main() {
	var numero1 float32 = 6;
	var numero2 float32 = 4;
	calculadora(numero1, numero2);
	fmt.Println(retornarValores());
	fmt.Println(devolverTexto());
	fmt.Println(gorras(45, "$"));
	pantalon("Largo", "clasico", "sin bolsillo")
}

func retornarValores() (string, int){
	dato1 := "Hola"
	dato2 := 33;

	return dato1, dato2
}

func devolverTexto() (dato1 string, dato2 int){
	dato1 = "Hola"
	dato2 = 33;

	return
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32

	if (op == "+") {
		resultado = n1 + n2
	}

	if (op == "-") {
		resultado = n1 - n2
	}

	if (op == "*") {
		resultado = n1 * n2
	}

	if (op == "/") {
		resultado = n1 / n2
	}

	return resultado
}

func calculadora(n1 float32, n2 float32) {
	//Suma
	fmt.Println("La suma es:")
	fmt.Println(operacion(n1, n2, "+"))

	//Resta
	fmt.Println("La resta es:")
	fmt.Println(operacion(n1, n2, "-"))

	//Multiplicación
	fmt.Println("La multiplicación es:")
	fmt.Println(operacion(n1, n2, "*"))

	//División
	fmt.Println("La división es:")
	fmt.Println(operacion(n1, n2, "/"))
}

//Este tipo de funciones se conoce como Closures
func gorras(pedido float32, moneda string) (string, float32, string){
	
	precio := func() float32 {
		return pedido * 7
	}

	return "El monto del pedido es:", precio(), moneda
}

func pantalon(caracteristicas ...string) {
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}
}
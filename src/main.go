package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaración de constantes 
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi: ",pi)
	fmt.Println("pi2: ",pi2)

	// Declaración de variable enteras
	base := 12
	var altura int = 14
	var area int 

	fmt.Println(base, altura, area)

	// Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma", result)

	// Resta
	result = y - x
	fmt.Println("Resta", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación", result)

	// Divison
	result = x / y
	fmt.Println("Divison", result)

	// Módulo
	result = x % y
	fmt.Println("Módulo", result)

	// Incrementar
	x++
	fmt.Println("Incrementar", x)

	// Decrementar
	x--
	fmt.Println("Decrementar", x)

	// Area de un Rectagulo
	baseR := 10
	alturaR := 20
	fmt.Println("Area de un Rectangulo es:", baseR * alturaR)

	// Area de un Trapecio
	baseTM := 9.5
	baseTm := 3.5
	alturaT := 4.0
	fmt.Println("Area de un Trapecio es:", ((baseTM + baseTm)/2)*alturaT)

	// Area de un Circulo
	radioC := 5.0
	fmt.Println("Area de un Circulo es", math.Pi*(radioC*radioC))

	// -----------------------------------------

	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("HelloMessage: %T\n", helloMessage)
	fmt.Printf("Cursos: %T", cursos)

}
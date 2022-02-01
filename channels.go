package main

import "fmt"

func main() {
	// This is what is called a buffered channel
	// c := make(chan int)
	// c <- 1

	// fmt.Println(<-c)

	// 	Unbuffered channel: Espera una función o una rutina para recibir el mensaje,
	// es bloqueada por ser llamada en la misma función

	// Buffered channel: Se puede llamar de manera inmediata,
	// en el siguiente ejemplo 2 es el numero de canales que pueden ser usados

	c2 := make(chan int, 3)
	c2 <- 1
	c2 <- 2
	c2 <- 3
	c2 <- 4

	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println(<-c2)

}

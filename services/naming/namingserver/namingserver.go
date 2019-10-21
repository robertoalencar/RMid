package main

import (
	"Plataformas-Exercicio03/RobertoMiddleware/services/naming/invoker"
	"fmt"
)

func main() {

	fmt.Println("Naming server running!!")

	// control loop passed to invoker
	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}

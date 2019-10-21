package main

import (
	"RMid/services/naming/invoker"
	"fmt"
)

func main() {

	fmt.Println("Naming server running!!")

	// control loop passed to invoker
	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}

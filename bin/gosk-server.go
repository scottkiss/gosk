package main

import (
	"fmt"
	"github.com/scottkiss/gostas"
)

func main() {
	gostas.Mapping("/", "../root/publish")
	fmt.Println("running ...")
	gostas.Addr(":8088").Run()

}

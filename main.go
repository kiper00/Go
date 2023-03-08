package main

import (
	"flag"
	"fmt"
	mongodb "kiper/Go/mongoDB"
	"kiper/Go/postgre"
)

func main() {
	// local parse launch args to do something
	flag.Parse()
	args := flag.Args()
	fmt.Println("Test :" + args[0])
	switch args[0] {
	case "hello":
		fmt.Println("Hello!")
	case "mongo":
		mongodb.DemoMongoDB()
	case "postgre":
		postgre.DemoPostgre()
	}
}

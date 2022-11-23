/**
Package flag
~ package flag berisikan fungsionalitas untuk memparsing command line argument
*/

package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "put your database host")
	username := flag.String("username", "root", "put your database username")
	password := flag.String("password", "root", "put your database password")
	pin := flag.Int("pin", 1000, "put your pin 4 number")

	flag.Parse()

	fmt.Println("host: ", *host)
	fmt.Println("username: ", *username)
	fmt.Println("password: ", *password)
	fmt.Println("pin: ", *pin)
}

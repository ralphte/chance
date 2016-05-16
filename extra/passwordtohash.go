package main

import (
	"github.com/ralphte/chance/apps/passhash"
	"fmt"
)

func main() {
	pass := "Password!!"
	hash, err := passhash.HashString(pass)
	if err != nil {
		panic(err)
	}
	fmt.Println(hash)

}

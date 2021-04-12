package multimodgo

import (
	"fmt"

	"github.com/ankk13/gohash"
)

// Hello world funct
func Hello() string {
	str := "Hello, World."
	fmt.Println(str)
	return str
}

// HelloHash returns hash of helloWorld
func HelloHash() string {
	str := "Hello, World."
	client := gohash.Client{}
	hashedValue := client.SHA256(str)
	fmt.Println(hashedValue)
	return hashedValue
}

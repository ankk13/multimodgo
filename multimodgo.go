package multimodgo

import (
	"fmt"

	"github.com/ankk13/gohash"
	"github.com/ankk13/multimodgo/special"
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

// SpecialHello returns hash of helloWorld
func SpecialHello(name string) string {
	return special.Hello(name)
}

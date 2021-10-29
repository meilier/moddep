package moddep

import "fmt"

type Hello interface {
	Hello(s string)
}

type ModA int

func (m ModA) Hello(s string) {
	fmt.Println("hello", s)
}

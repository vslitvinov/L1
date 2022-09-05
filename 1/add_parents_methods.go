package main

import (
	"errors"
	"fmt"
	"strings"
)

// Дана структура Human (с произвольным набором полей и методов). 
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
// Human - describes human datatype
type Human struct {
	Name string
	Age  int
}

// String representation of human struct
func (h *Human) String() string {
	return fmt.Sprintf("%s %d years old", h.Name, h.Age)
}

// Rename - changes human name
func (h *Human) Rename(name string) error {
	if len(name) < 3 {
		return errors.New("name too short")
	}
	h.Name = strings.Title(strings.ToLower(name))
	return nil
}

// Action - holds all human's actions
type Action struct {
	Human
}

// NewHuman - human constructor
func NewAction(name string, age int) *Action {
	return &Action{
		Human{
			Name: name,
			Age:  age,
		},
	}
}

func main() {
	action := NewAction("Ivan", 22)

	fmt.Println(action)

	action.Rename("Dima")

	fmt.Println(action)
}

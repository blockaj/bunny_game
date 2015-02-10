package main

import (
	"fmt"
)

func main() {

}

//Creates a certain number of bunnies in linked-list
func createBunnies(root *Bunny, numberOfBunnies int) Bunny {
	root = NewBunny()
	ranger := &Bunny{}
	ranger = root // Will move through list

	for i := 0; i < numberOfBunnies - 1; i++ { 
		bunny := NewBunny()
		ranger.Next = bunny
		ranger = ranger.Next
		ranger.Next = nil
	}
	return *root
}

//Prints entire list of bunnies given root of linked-list
func printList(root *Bunny) {
	ranger := &Bunny{}
	ranger = root
	if ranger != nil {
		for ranger.Next != nil {

		}
	}
}

//Prints single bunny; used in func printList() for formatting purposes
func printBunny(inputBunny Bunny) {

}
package main

import (
	"fmt"
)

func main() {
	root := createBunnies(5)
	printList(root)
}

//Creates a certain number of bunnies in linked-list
func createBunnies(numberOfBunnies int) *Bunny {
	root := NewBunny()
	ranger := &Bunny{}
	ranger = root // Will move through list

	for i := 0; i < numberOfBunnies - 1; i++ { 
		bunny := NewBunny()
		ranger.Next = bunny
		ranger = ranger.Next
		ranger.Next = nil
	}
	return root
}

//Prints entire list of bunnies given root of linked-list
func printList(root *Bunny) {
	ranger := &Bunny{}
	ranger = root
	if ranger != nil {
		for ranger.Next != nil {
			printBunny(ranger)
			ranger = ranger.Next
		}
		printBunny(ranger)
	}
}

//Prints single bunny; used in func printList() for formatting purposes
func printBunny(inputBunny *Bunny) {
	fmt.Println(inputBunny.Name(), "is a", inputBunny.Age(), "year(s) old", inputBunny.Gender().String(), "bunny with", inputBunny.Color().String(), "fur.")
}
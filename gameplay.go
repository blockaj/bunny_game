package main

import (
	"fmt"
	"time"
	"reflect"
)

//Creates a certain number of bunnies in linked-list
func createBunnies(numberOfBunnies int) *Bunny {
	root := GenerateNewBunny()
	ranger := &Bunny{}
	ranger = root // Will move through list

	for i := 0; i < numberOfBunnies-1; i++ {
		bunny := GenerateNewBunny()
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
	fmt.Println(inputBunny.Name(), "is a", inputBunny.Age(), "year(s) old",
		inputBunny.Gender().String(), "bunny with",
		inputBunny.Color().String(), "fur.")
}

//Creates new bunnies based on current make-up of bunnies
func procreate(root *Bunny) *Bunny {
	numberOfBabies := root.FindNumberOfBabies()
	newBorns := createBunnies(numberOfBabies)
	root = root.AppendBunnyList(newBorns)
	return root
}

//Kill off bunnies, a.k.a remove node from linked-list
func killBunny(root *Bunny) *Bunny {
	ranger := root
	if ranger != nil {
		if ranger.Age() >= 10 && !ranger.mutant {
			*root = *ranger.Next
		} else if ranger.mutant && ranger.Age() == 50 {
			*root = *ranger.Next
		}
		for ranger != nil {
			rangerType := reflect.TypeOf(ranger.Next)
			fmt.Println(reflect.Zero(rangerType).String())
			if ranger.Next.ShouldBunnyDie() {
				ranger.Next = ranger.Next.Next
			}
			ranger = ranger.Next
		}
	}
	return root
}

//Defines rhythm of game
func turn(root *Bunny, turnNum int) {
	ranger := &Bunny{}
	if turnNum == 0 {
		fmt.Println("Starting creation…")
		root = createBunnies(5)
		printList(root)
	}
	if root != nil {
		fmt.Println("Starting regular cycle...")
		root = procreate(root)
		printList(root)
		root = killBunny(root)
		fmt.Println("Massacred bunny list")
		printList(root)
		ranger = root
		for ranger.Next != nil {
			ranger.SetAge(ranger.Age() + 1)
			ranger = ranger.Next
			printBunny(ranger)
		}
		printList(ranger)
		ranger.SetAge(ranger.Age() + 1)
		time.Sleep(1 * time.Second)
		printBunny(ranger)
		ranger = root
		for ranger.Next != nil {
			if ranger.mutant {
				if !ranger.Next.mutant {
					ranger.Next.mutant = true
				}
			}
		}
		turn(root, turnNum+1)
	}
}

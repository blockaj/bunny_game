package main

import (
	"fmt"
	"time"
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
	var listNumber int
	ranger := &Bunny{}
	ranger = root
	if ranger != nil {
		listNumber = 1
		for ranger.Next != nil {
			fmt.Printf("%d. %s \n", listNumber, ranger.String())
			ranger = ranger.Next
			listNumber++
		}
		fmt.Printf("%d. %s \n\n", listNumber, ranger.String())
	}
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
	if ranger.ShouldBunnyDie() {
		root = ranger.Next
	}
	
	for ranger.Next.Next != nil {
		if ranger.Next.ShouldBunnyDie() {
			fmt.Println("Killing", ranger.Next.String())
			ranger.Next = ranger.Next.Next
		}
		ranger = ranger.Next
	}
	return root
}

//Defines rhythm of game
func turn(root *Bunny, turnNum int) {
	ranger := &Bunny{}
	if turnNum == 0 {
		fmt.Println("Starting creation…")
		root = createBunnies(5)
	}
	if root != nil {
		fmt.Println("Starting regular cycle...")
		printList(root)
		root = procreate(root)
		time.Sleep(3 * time.Second)
		fmt.Println("After procreation")
		printList(root)
		root = killBunny(root)
		time.Sleep(3 * time.Second)
		fmt.Println("After killing bunnies")
		printList(root)
		ranger = root
		for ranger.Next != nil {
			ranger.SetAge(ranger.Age() + 1)
			ranger = ranger.Next
		}
		ranger.SetAge(ranger.Age() + 1)
		time.Sleep(3 * time.Second)
		ranger = root
		for ranger.Next != nil {
			if ranger.mutant {
				if !ranger.Next.mutant {
					ranger.Next.mutant = true
				}
			}
			ranger = ranger.Next
		}
		turn(root, turnNum+1)

	}
}

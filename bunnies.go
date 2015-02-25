/*
Contains everything related to any one bunny.
This includes various enums (color and sex) as well
as structure types.
*/
package main

import (
	"strconv"
	"math/rand"
	"time"
)

//enum for sex of bunny
type Sex int

const (
	MALE Sex = 1 + iota
	FEMALE
)

//returns stringified version of bunny sex
func (s Sex) String() string {
	var returnString string
	switch s {
	case MALE:
		returnString = "male"
	case FEMALE:
		returnString = "female"
	}
	return returnString
}

//enum for color of bunny
type Color int

const (
	BROWN Color = iota
	WHITE
	BLACK
	SPOTTED
)

//returns stringified version of bunny color
func (c Color) String() string {
	var returnString string
	switch c {
	case BROWN:
		returnString = "brown"
	case WHITE:
		returnString = "white"
	case BLACK:
		returnString = "black"
	case SPOTTED:
		returnString = "spotted"
	}
	return returnString
}

type Bunny struct {
	gender Sex
	color  Color
	age    int
	name   string
	mutant bool
	Next   *Bunny
}

//Gender getter
func (b *Bunny) Gender() Sex {
	return b.gender
}

//Color getter
func (b *Bunny) Color() Color {
	return b.color
}

//Age getter
func (b *Bunny) Age() int {
	return b.age
}

//Name getter
func (b *Bunny) Name() string {
	return b.name
}

//Set gender
func (b *Bunny) SetGender(inputGender Sex) {
	b.gender = inputGender
}

//Set color
func (b *Bunny) SetColor(inputColor Color) {
	b.color = inputColor
}

//Set age
func (b *Bunny) SetAge(inputAge int) {
	b.age = inputAge
}

//Set name
func (b *Bunny) SetName(inputName string) {
	b.name = inputName
}
func (b *Bunny) ShouldBunnyDie() bool {
	var returnVal bool
	if b.Age() == 10 && !b.mutant {
		returnVal = true
	} else if b.Age() == 50 && b.mutant {
		returnVal = true
	} else {
		returnVal = false
	}
	return returnVal
}
func (b *Bunny) AppendBunnyList(root *Bunny) *Bunny {
	ranger := b
	if ranger != nil {
		for ranger.Next != nil {
			ranger = ranger.Next
		}
		ranger.Next = root 
	}
	return b
}
func (b *Bunny) FindNumberOfBabies() int {
	ranger := b
	var viableMales, viableFemales int

	if ranger != nil {
		for ranger.Next != nil {
			if ranger.Gender() == MALE && ranger.Age() >= 2 {
				viableMales++
			}
			ranger = ranger.Next
		}
		if ranger.Gender() == MALE && ranger.Age() >= 2 {
			viableMales++
		}
	}

	ranger = b
	if ranger != nil {
		for ranger.Next != nil {
			if ranger.Gender() == FEMALE && ranger.Age() >= 2 {
				viableFemales++
			}
			ranger = ranger.Next
		}
		if ranger.Gender() == FEMALE && ranger.Age() >= 2 {
			viableFemales++
		}
	}
	numberOfBabies := viableMales * viableFemales
	return numberOfBabies
}

func GenerateNewBunny() *Bunny {
	maleNameOptions := []string{"Roger", "Donald", "Ralph", "Fluffy", "Snowball", 
							   "Arturio", "Dracula", "Barnabus",
							"Dolph", "Julio", "Voldemort"}
	femaleNameOptions := []string{"Penelope", "Artemis", "Bella", "Anna", "Sharyl", 
								 "Joanne", "Marissa", "Rita", "Nora", "Kelly", "Jesus"}
	bunnyMutantBool := false
	numberGenerator := rand.New(rand.NewSource(time.Now().UnixNano())) //A straight up number generator dude
	bunnyGender := Sex(numberGenerator.Int()%2 + 1)
	bunnyColor := Color(numberGenerator.Int() % 4)
	bunnyMutant := numberGenerator.Int() % 50
	if bunnyMutant == 1 {
		bunnyMutantBool = true
	}
	var bunnyName string
	if bunnyGender == MALE {
		bunnyName = maleNameOptions[numberGenerator.Int()%len(maleNameOptions)] //Index is a random number from 0-len(nameOptions)-1
	} else {
		bunnyName = femaleNameOptions[numberGenerator.Int()%len(femaleNameOptions)]
	}

	bunnyAge := numberGenerator.Int()%10 + 1

	//The new bunny with all the randomly generated features
	newBunny := &Bunny{
		color:  bunnyColor,
		gender: bunnyGender,
		age:    bunnyAge,
		name:   bunnyName,
		mutant: bunnyMutantBool,
	}

	return newBunny
}
func (b *Bunny) length() int {
	var length int
	if b != nil {
		for b.Next != nil {
			length++
			b = b.Next
		}
		length++
	}
	return length
}
//Prints single bunny; used in func printList() for formatting purposes
func (b *Bunny) String() string {
	bunnyString := b.Name() + ", " + strconv.Itoa(b.Age()) + " year(s) old, " +
		b.Gender().String() + ", and " +
		b.Color().String() + "."

	return bunnyString
}

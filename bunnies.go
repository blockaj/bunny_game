package main

import (
	"math/rand"
	"time"
)

type Sex int

const (
	MALE Sex = 1 + iota
	FEMALE
)

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

type Color int

const (
	BROWN Color = iota
	WHITE
	BLACK
	SPOTTED
)

func (c Color) String() string {
	var returnString string
	switch c {
	case BROWN:
		returnString = "Brown"
	case WHITE:
		returnString = "White"
	case BLACK:
		returnString = "Black"
	case SPOTTED:
		returnString = "Spotted"
	}
	return returnString
}

type Bunny struct {
	gender Sex
	color  Color
	age    int
	name   string
	Next   *Bunny
}

func (b *Bunny) Gender() Sex {
	return b.gender
}
func (b *Bunny) Color() Color {
	return b.color
}
func (b *Bunny) Age() int {
	return b.age
}
func (b *Bunny) Name() string {
	return b.name
}
func (b *Bunny) SetGender(inputGender Sex) {
	b.gender = inputGender
}
func (b *Bunny) SetColor(inputColor Color) {
	b.color = inputColor
}
func (b *Bunny) SetAge(inputAge int) {
	b.age = inputAge
}
func (b *Bunny) SetName(inputName string) {
	b.name = inputName
}
func NewBunny() Bunny {
	numberGenerator := rand.New(rand.NewSource(time.Now().UnixNano())) //A straight up number generator dude
	nameOptions := []string{"Roger", "Gracie", "Bella", "Fluffy", "Snowball", "Penelope", "Dracula"}
	bunnyGender := Sex(numberGenerator.Int()%2 + 1)
	bunnyColor := Color(numberGenerator.Int() % 4)
	bunnyName := nameOptions[numberGenerator.Int()%len(nameOptions)] //Index is a random number from 0-len(nameOptions)-1
	bunnyAge := numberGenerator.Int()%10 + 1
	newBunny := Bunny{color: bunnyColor,
		gender: bunnyGender,
		age:    bunnyAge,
		name:   bunnyName,
	}

	return newBunny
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const catFeedPerKg int = 7
const dogFeedPerKg int = 10
const cowFeedPerKg int = 25

type FarmAnimal interface {
	feedWeightPerMonth() int
	String()
}

//Basic animal struct
type Animal struct {
	Weight    int
	FeedPerKg int
	Name      string
}

func (animal Animal) feedWeightPerMonth() int {
	return animal.Weight * animal.FeedPerKg
}

func (animal Animal) String() {
	fmt.Printf("Me name is %s. My weight is %d kilos. I need %d kilos of feed per month.\n", animal.Name, animal.Weight, animal.feedWeightPerMonth())
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

type Cow struct {
	Animal
}

func typeFarmInfo(farm []FarmAnimal) (totalFarmFeedPerMonth int) {

	for _, a := range farm {
		a.String()
		totalFarmFeedPerMonth += a.feedWeightPerMonth()
	}
	return
}

//A function that creates animals on the farm. n - total number of each animals on the farm
func makeFarm(n int) (farm []FarmAnimal) {

	var animal FarmAnimal

	minCatWeight := 2
	maxCatWeight := 8

	minDogWeight := 5
	maxDogWeight := 10

	minCowWeight := 50
	maxCowWeight := 250

	minAnimalRange := 1
	maxAnimalRange := 3
	randShift := 1

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		newRandomAnimal := rand.Intn(maxAnimalRange-minAnimalRange+randShift) + minAnimalRange // Create new animal: 1 - cat, 2- dog, 3- cow
		switch newRandomAnimal {
		case 1:
			newCatWeight := rand.Intn(maxCatWeight-minCatWeight) + minCatWeight
			animal = Cat{Animal{Weight: newCatWeight, FeedPerKg: catFeedPerKg, Name: "Cat"}}

		case 2:
			newDogWeight := rand.Intn(maxDogWeight-minDogWeight) + minDogWeight
			animal = Dog{Animal{Weight: newDogWeight, FeedPerKg: dogFeedPerKg, Name: "Dog"}}

		case 3:
			newCowWeight := rand.Intn(maxCowWeight-minCowWeight) + minCowWeight
			animal = Cow{Animal{Weight: newCowWeight, FeedPerKg: cowFeedPerKg, Name: "Cow"}}

		}
		farm = append(farm, animal)
	}

	return
}

func main() {
	maxNumberOfAnimals := 15                                                          //Any value can be specified.
	farm := makeFarm(maxNumberOfAnimals)                                              //Making new farm with random num of cats, dogs and cows
	fmt.Printf("The farm needs %d kilograms of feed per month\n", typeFarmInfo(farm)) //Printing information about farm
}

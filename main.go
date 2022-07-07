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
	fmt.Stringer
}

type Cat struct {
	catWeight    int
	catFeedPerKg int
	catName      string
}

func (cat Cat) feedWeightPerMonth() int {
	return cat.catWeight * catFeedPerKg
}

func (cat Cat) String() string {
	return fmt.Sprintf("My name is %s. My weight is %d kilos. I need %d kilos of feed per month.", cat.catName, cat.catWeight, cat.feedWeightPerMonth())
}

type Dog struct {
	dogWeight    int
	dogFeedPerKg int
	dogName      string
}

func (dog Dog) feedWeightPerMonth() int {
	return dog.dogWeight * dog.dogFeedPerKg
}

func (dog Dog) String() string {
	return fmt.Sprintf("My name is %s. My weight is %d kilos. I need %d kilos of feed per month.", dog.dogName, dog.dogWeight, dog.feedWeightPerMonth())
}

type Cow struct {
	cowWeight    int
	cowFeedPerKg int
	cowName      string
}

func (cow Cow) feedWeightPerMonth() int {
	return cow.cowWeight * cow.cowFeedPerKg
}

func (cow Cow) String() string {
	return fmt.Sprintf("My name is %s. My weight is %d kilos. I need %d kilos of feed per month.", cow.cowName, cow.cowWeight, cow.feedWeightPerMonth())
}

func typeFarmInfo(farm []FarmAnimal) int {

	totalFarmFeedPerMonth := 0

	for _, a := range farm {
		fmt.Println(a.String())
		totalFarmFeedPerMonth += a.feedWeightPerMonth()
	}

	return totalFarmFeedPerMonth
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
			animal = Cat{newCatWeight, catFeedPerKg, "Cat"}

		case 2:
			newDogWeight := rand.Intn(maxDogWeight-minDogWeight) + minDogWeight
			animal = Dog{newDogWeight, dogFeedPerKg, "Dog"}

		case 3:
			newCowWeight := rand.Intn(maxCowWeight-minCowWeight) + minCowWeight
			animal = Cow{newCowWeight, cowFeedPerKg, "Cow"}

		}
		farm = append(farm, animal)
	}

	return
}

func main() {
	maxNumberOfAnimals := 15                                                           //Any value can be specified.
	farm := makeFarm(maxNumberOfAnimals)                                               //Making new farm with random num of cats, dogs and cows
	fmt.Printf("The farm needs %d kilograms of feed per month.\n", typeFarmInfo(farm)) //Printing information about farm
}

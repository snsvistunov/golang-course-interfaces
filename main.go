package main

import (
	"fmt"
	"math/rand"
	"time"
)

const catFeedPerKg int = 7
const dogFeedPerKg int = 10
const cowFeedPerKg int = 25
const minCatWeight int = 2
const maxCatWeight int = 8
const minDogWeight int = 5
const maxDogWeight int = 10
const minCowWeight int = 50
const maxCowWeight int = 250

type FarmAnimal interface {
	feedWeightPerMonth() int
	getAnimalName() string
	getAnimalWeight() int
	getIsEdible() bool
	fmt.Stringer
}

type Cat struct {
	catWeight    int
	catFeedPerKg int
	catName      string
	isEdible     bool
}

func (cat Cat) feedWeightPerMonth() int {
	return cat.catWeight * catFeedPerKg
}

func (cat Cat) String() string {
	return fmt.Sprintf("My name is %s. My weight is %d kilos. I need %d kilos of feed per month.", cat.catName, cat.catWeight, cat.feedWeightPerMonth())
}

func (cat Cat) getAnimalName() string {
	return cat.catName
}

func (cat Cat) getAnimalWeight() int {
	return cat.catWeight
}

func (cat Cat) getIsEdible() bool {
	return cat.isEdible
}

type Dog struct {
	dogWeight    int
	dogFeedPerKg int
	dogName      string
	isEdible     bool
}

func (dog Dog) feedWeightPerMonth() int {
	return dog.dogWeight * dog.dogFeedPerKg
}

func (dog Dog) String() string {
	return fmt.Sprintf("My name is %s. My weight is %d kilos. I need %d kilos of feed per month.", dog.dogName, dog.dogWeight, dog.feedWeightPerMonth())
}

func (dog Dog) getAnimalName() string {
	return dog.dogName
}

func (dog Dog) getAnimalWeight() int {
	return dog.dogWeight
}

func (dog Dog) getIsEdible() bool {
	return dog.isEdible
}

type Cow struct {
	cowWeight    int
	cowFeedPerKg int
	cowName      string
	isEdible     bool
}

func (cow Cow) feedWeightPerMonth() int {
	return cow.cowWeight * cow.cowFeedPerKg
}

func (cow Cow) String() string {
	return fmt.Sprintf("My name is %s. My weight is %d kilos. I need %d kilos of feed per month.", cow.cowName, cow.cowWeight, cow.feedWeightPerMonth())
}

func (cow Cow) getAnimalName() string {
	return cow.cowName
}

func (cow Cow) getAnimalWeight() int {
	return cow.cowWeight
}

func (cow Cow) getIsEdible() bool {
	return cow.isEdible
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

	minAnimalRange := 1
	maxAnimalRange := 3
	randShift := 1

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		newRandomAnimal := rand.Intn(maxAnimalRange-minAnimalRange+randShift) + minAnimalRange // Create new animal: 1 - cat, 2- dog, 3- cow
		switch newRandomAnimal {
		case 1:
			newCatWeight := rand.Intn(maxCatWeight-minCatWeight) + minCatWeight
			animal = Cat{newCatWeight, catFeedPerKg, "Cat", false}

		case 2:
			newDogWeight := rand.Intn(maxDogWeight-minDogWeight) + minDogWeight
			animal = Dog{newDogWeight, dogFeedPerKg, "Dog", false}

		case 3:
			newCowWeight := rand.Intn(maxCowWeight-minCowWeight) + minCowWeight
			animal = Cow{newCowWeight, cowFeedPerKg, "Cow", true}

		}
		farm = append(farm, animal)
	}

	return
}

//Check animal type

func checkAnimalType(animal FarmAnimal) error {
	var err error
	var animalType string
	switch animal.(type) {
	case Dog:
		animalType = "Dog"
	case Cat:
		animalType = "Cat"
	case Cow:
		animalType = "Cow"
	}
	if animal.getAnimalName() != animalType {
		err = fmt.Errorf("Animal name does not match its type. %s can't be %s", animalType, animal.getAnimalName())
	}
	return err
}

func checkAnimalWeight(animal FarmAnimal) error {
	var err error
	var minWeight int
	var maxWeight int
	switch animal.(type) {
	case Dog:
		minWeight = minDogWeight
		maxWeight = maxDogWeight
	case Cat:
		minWeight = minCatWeight
		maxWeight = maxCatWeight
	case Cow:
		minWeight = minCowWeight
		maxWeight = maxCowWeight
	}
	if animal.getAnimalWeight() < minWeight || animal.getAnimalWeight() > maxWeight {
		err = fmt.Errorf("Animal has incorrect weight. Current weight is %v kilos. Minimum weight is %v kilos, maximum weight is %v kilos", animal.getAnimalWeight(), minWeight, maxWeight)
	}
	return err
}

func checkIsEdeble(animal FarmAnimal) error {
	var err error
	var isEdeble bool
	switch animal.(type) {
	case Dog:
		isEdeble = false
	case Cat:
		isEdeble = false
	case Cow:
		isEdeble = true
	}
	if animal.getIsEdible() != isEdeble {
		err = fmt.Errorf("Invalid isEdible property")
	}
	return err
}

func checkFarm(farm []FarmAnimal) error {
	var err error
	for _, v := range farm {
		err = checkAnimalType(v)
		if err != nil {
			err = fmt.Errorf("Type validation failed:%w", err)
			return err
		}
		err = checkAnimalWeight(v)
		if err != nil {
			err = fmt.Errorf("Weight validation failed:%w", err)
			return err
		}
		err = checkIsEdeble(v)
		if err != nil {
			err = fmt.Errorf("Edeble validation failed:%w", err)
			return err
		}
	}
	return nil
}
func main() {
	maxNumberOfAnimals := 15                     //Any value can be specified.
	farm := makeFarm(maxNumberOfAnimals)         //Making new farm with random num of cats, dogs and cows
	farm = append(farm, Cow{1, 2, "Cow", false}) //Manual incorrect data input
	err := checkFarm(farm)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The farm needs %d kilograms of feed per month.\n", typeFarmInfo(farm)) //Printing information about farm

}

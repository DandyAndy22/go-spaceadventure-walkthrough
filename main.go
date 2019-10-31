package main

import "fmt"

func main() {
	printWelcome() 
	printGreeting(getName()) //input getName() instead piping, function is mselly if it does more than one thing. beyond it's defined capacity. 
	fmt.Println("Let's go on an adventure!") //extract to a function perhaps
	travel()
}

func printWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

func getName() string{
	var name string
	fmt.Println("What is your name?")
    fmt.Scan(&name)
    return name
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = getInput("Name the planet you would like to visit.")
		if choice == "Y" {
			travelToRandomPlanet()
		} else if choice == "N" {
			planetName := getInput("La di da")
			travelToPlanet(planetName)
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}

func getInput(prompt string) string {
	var response string
	fmt.Println(prompt)
	fmt.Scan(&response)
	return response
}

func travelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

// func getPlanetName() string {
// 	var name string
// 	fmt.Println("Name the planet you would like to visit.")
// 	fmt.Scan(&name)
// 	return name
//}

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}

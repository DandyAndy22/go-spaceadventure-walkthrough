package main

import "fmt"

func main() {
	printWelcome() 
	printGreeting(getInput("What is your name?")) //input getName() instead piping, function is mselly if it does more than one thing. beyond it's defined capacity. 
	adventureTime()
	travel()
}

func printWelcome() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 8 planets to explore.")
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func adventureTime() {
	fmt.Println("Let's go on an adventure!")
}

func travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = getInput("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			travelToRandomPlanet()
		} else if choice == "N" {
			travelToPlanet(getInput("Name the planet you would like to visit."))
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

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
		if planetName == "Neptune" {
			fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")

		}else if planetName == "Mars" {
			fmt.Println("Arrived at the Mars, the red planet. Watch out for the locals!")

		}else if planetName == "Venus" {
			fmt.Println("Arrived at Venus, the other one near Earth. This place is wild.")

		}else if planetName == "Mercury" {
			fmt.Println("Arrived at Mercury. Where you're hot and you're cold and you're yes and you're no..")
			
		}else if planetName == "Jupiter" {
			fmt.Println("Arrived at Jupiter. It's a real gas, man.")
			
		}else if planetName == "Uranus" {
			fmt.Println("Arrived at Uranus. Nothing special, ooph.")
			
		}else if planetName == "Earth" {
			fmt.Println("Arrived at Earth. The little blue dot, great surfing, cheese and tax shelters here.")
			
		}else if planetName == "Saturn" {
			fmt.Println("Arrived at Saturn. Watch out, you're liable to get mooned out here.")
			
		}else if planetName == "Pluto" {
			fmt.Println("Arrived at Pluto. Fake planet.")
			
		}else{
			fmt.Println("Sorry, I didn't get that.")
		}
}

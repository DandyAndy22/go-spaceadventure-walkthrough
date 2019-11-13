package spaceadventure

import "fmt"
//import "path of go directory in project"

func Start(planetarySystem PlanetarySystem) {
	printWelcome(planetarySystem) 
	printGreeting(getInput("What is your name?")) //input getName() instead piping, function is mselly if it does more than one thing. beyond it's defined capacity. 
	adventureTime()
	travel(destinationPrompt("Shall I randomly choose a planet for you to visit? (Y or N)"))
}

func printWelcome(planetarySystem PlanetarySystem) {
	fmt.Println("Welcome to the %s", planetarySystem.Name)
	fmt.Println("There are %d planets to explore.\n", len(planetarySystem.Planets))//Make a planetary System lengtho or number of planets attributes
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func adventureTime() {
	fmt.Println("Let's go on an adventure!")
}

func travel(randomDestination bool){
	if(randomDestination) {
		travelToRandomPlanet();
	}else {
		travelToPlanet(getInput("Name the planet you would like to visit."))
	}
}

func destinationPrompt(prompt string) bool {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = getInput(prompt)
		if choice == "Y" {
			return true
		} else if choice == "N" {
			return false
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
	return false
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
			fmt.Println("Arrived at Mars, the red planet. Watch out for the locals!")

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
			
		}
}
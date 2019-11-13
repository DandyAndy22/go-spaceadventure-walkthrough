package spaceadventure

type PlanetarySystem struct {
	Name string
	Planets []Planet
}

func (ps PlanetarySystem) NumberOfPlanets() int{
	return len(ps.Planets)
}
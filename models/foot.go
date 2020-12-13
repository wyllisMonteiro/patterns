package models

// FootPlayer Data model for a foot player
type FootPlayer struct {
	Name string
	Age  int
	Job  string
	Team string
}

// FootPlayers Stores all foot player
type FootPlayers []FootPlayer

// GetFootPlayers Get fake data
func GetFootPlayers() FootPlayers {
	return FootPlayers{
		FootPlayer{
			Name: "Cristiano Ronaldo",
			Age:  34,
			Job:  "Attaquant",
			Team: "Portugal",
		},
		FootPlayer{
			Name: "Lionel Messi",
			Age:  33,
			Job:  "Attaquant",
			Team: "Argentine",
		},
	}
}

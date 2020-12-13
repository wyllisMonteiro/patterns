package models

// BasketPlayer Data model for a basket player
type BasketPlayer struct {
	Name string
	Age  int
	Size float32
	Job  string
	Team string
}

// BasketPlayers Stores all foot player
type BasketPlayers []BasketPlayer

// GetBasketPlayers Get fake data
func GetBasketPlayers() BasketPlayers {
	return BasketPlayers{
		BasketPlayer{
			Name: "Lebron James",
			Age:  35,
			Size: 2.06,
			Job:  "Ailier",
			Team: "Lakers",
		},
		BasketPlayer{
			Name: "Kyrie Irving",
			Age:  28,
			Size: 1.88,
			Job:  "Meneur",
			Team: "Nets",
		},
	}
}

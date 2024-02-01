package main

type FootballPlayer struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

type FootballClub struct {
	Name     string           `json:"name"`
	Location string           `json:"location"`
	Players  []FootballPlayer `json:"players"`
}

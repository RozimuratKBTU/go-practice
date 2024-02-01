package main

var footballClubs = map[string]FootballClub{
	"manutd": {
		Name:     "Manchester United",
		Location: "Manchester",
		Players: []FootballPlayer{
			{Name: "Harry Maguire", Age: 30, Position: "Center. defender"},
			{Name: "Casemiro", Age: 31, Position: "Midfielder"},
			{Name: "Andre", Age: 27, Position: "Goalkeeper"},
		},
	},
	"liverpool": {
		Name:     "Liverpool",
		Location: "Liverpool",
		Players: []FootballPlayer{
			{Name: "Mohamed Salah", Age: 31, Position: "Forward"},
			{Name: "Virgil van Dijk", Age: 30, Position: "Defender"},
			{Name: "Allison Backer", Age: 31, Position: "Goalkeeper"},
		},
	},
	"chelsea": {
		Name:     "Chelsea",
		Location: "London",
		Players: []FootballPlayer{
			{Name: "Thiago Silva", Age: 38, Position: "Defender"},
			{Name: "Cole Palmer", Age: 21, Position: "Midfielder"},
			{Name: "Robert Sanchez", Age: 26, Position: "Goalkeeper"},
		},
	},
	"arsenal": {
		Name:     "Arsenal",
		Location: "London",
		Players: []FootballPlayer{
			{Name: "Pierre-Emerick Aubameyang", Age: 32, Position: "Forward"},
			{Name: "Bukayo Saka", Age: 20, Position: "Midfielder"},
		},
	},
	"tottenham": {
		Name:     "Tottenham Hotspur",
		Location: "London",
		Players: []FootballPlayer{
			{Name: "Richarlison", Age: 28, Position: "Forward"},
			{Name: "Heung-Min Son", Age: 29, Position: "Forward"},
		},
	},
	"astonvilla": {
		Name:     "Aston Villa",
		Location: "Birmingham",
		Players: []FootballPlayer{
			{Name: "Emeliano Martinez", Age: 26, Position: "Goalkeeper"},
			{Name: "Ollie Watkins", Age: 25, Position: "Forward"},
		},
	},
}

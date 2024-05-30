package driverteam

import consomeapi "github.com/RennanLopes/automacaoqueonetics/ConsomeApi"

// Estrutura DriverTeam
type DriverTeam struct {
	Description     string `json:"description"`
	Name            string `json:"name"`
	FiliationType   string `json:"filiationType"`
	Drivers         []int  `json:"drivers"`
	Representatives []int  `json:"representatives"`
	Organization    struct {
		ID int `json:"id"`
	} `json:"organization"`
}

// Envia os dados de Driverteam
func PostDriverTeam(driverTeam DriverTeam) int {
	return consomeapi.Post("http://desafio-api-trix.trixlog.com/driverteam/", driverTeam)
}

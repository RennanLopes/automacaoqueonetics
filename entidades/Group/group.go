package group

import consomeapi "github.com/RennanLopes/automacaoqueonetics/ConsomeApi"

// Estrutura Group
type Group struct {
	Disabled     bool   `json:"disabled"`
	IsAdmin      bool   `json:"isAdmin"`
	Name         string `json:"name"`
	Organization struct {
		ID int `json:"id"`
	} `json:"organization"`
}

// Envia os dados de Group
func PostGroup(group Group) int {
	return consomeapi.Post("http://desafio-api-trix.trixlog.com/group/", group)
}

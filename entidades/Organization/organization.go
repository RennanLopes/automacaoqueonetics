package organization

import consomeapi "github.com/RennanLopes/automacaoqueonetics/ConsomeApi"

// Estrutura de Organization
type Organization struct {
	Address           string `json:"address"`
	City              string `json:"city"`
	Country           string `json:"country"`
	CpfCnpj           string `json:"cpfcnpj"`
	Description       string `json:"description"`
	Dst               bool   `json:"dst"`
	Email             string `json:"email"`
	Gmt               int    `json:"gmt"`
	Name              string `json:"name"`
	HierarchicalLevel struct {
		ID int `json:"id"`
	} `json:"hierarchicalLevel"`
	ParentOrganization struct {
		ID int `json:"id"`
	} `json:"parentOrganization"`
}

// Envia os dados de Organization
func PostOrganization(org Organization) int {
	return consomeapi.Post("http://desafio-api-trix.trixlog.com/organization/", org)
}

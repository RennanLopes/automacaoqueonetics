package driver

import consomeapi "github.com/RennanLopes/automacaoqueonetics/ConsomeApi"

// Estrutura de Driver
type Driver struct {
	DriverTeam struct {
		ID int `json:"id"`
	} `json:"driverTeam"`
	HiringType        string `json:"hiringType"`
	License           string `json:"license"`
	LicenseRegister   string `json:"licenseRegister"`
	LicenseExpedition string `json:"licenseExpedition"`
	LicenseExpiration string `json:"licenseExpiration"`
	Licenses          []struct {
		License string `json:"license"`
	} `json:"licenses"`
	Name         string `json:"name"`
	FirstLicense string `json:"firstLicense"`
	BirthDate    string `json:"birthDate"`
	Registration string `json:"registration"`
	Status       string `json:"status"`
	Type         string `json:"type"`
}

// Envia os dados de Driver
func PostDriver(driver Driver) int {
	return consomeapi.Post("http://desafio-api-trix.trixlog.com/driver/", driver)
}

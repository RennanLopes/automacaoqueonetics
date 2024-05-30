package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	driver "github.com/RennanLopes/automacaoqueonetics/entidades/Driver"
	driverteam "github.com/RennanLopes/automacaoqueonetics/entidades/DriverTeam"
	group "github.com/RennanLopes/automacaoqueonetics/entidades/Group"
	organization "github.com/RennanLopes/automacaoqueonetics/entidades/Organization"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	// Abrir o arquivo CSV
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Ler o arquivo CSV
	reader := csv.NewReader(file)
	reader.Comma = ';'

	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Pular a primeira linha com cabeçalho
	headers := rows[0]
	data := rows[1:]

	for _, record := range data {
		fieldMap := map[string]string{}
		for i, header := range headers {
			fieldMap[header] = record[i]
		}

		// Criar Organização
		org := organization.Organization{
			Address:     fieldMap["address"],
			City:        fieldMap["city"],
			Country:     fieldMap["country"],
			CpfCnpj:     fieldMap["cpfcnpj"],
			Description: fieldMap["description"],
			Dst:         fieldMap["dst"] == "true",
			Email:       fieldMap["email"],
			Gmt:         parseInt(fieldMap["gmt"]),
			Name:        fieldMap["name"],
		}
		org.HierarchicalLevel.ID = 4
		org.ParentOrganization.ID = 2

		orgID := organization.PostOrganization(org)
		if orgID == 0 {
			log.Fatal("Falha so criar Organização")
		}

		// Criar DriverTeam
		driverTeam := driverteam.DriverTeam{
			Description:     fieldMap["driverTeamDescription"],
			Name:            fieldMap["driverTeamName"],
			FiliationType:   fieldMap["filiationType"],
			Drivers:         []int{},
			Representatives: parseIntArray(strings.Split(fieldMap["representatives"], ",")),
		}
		driverTeam.Organization.ID = orgID

		driverTeamID := driverteam.PostDriverTeam(driverTeam)
		if driverTeamID == 0 {
			log.Fatal("Falha ao criar DriverTeam")
		}

		// Criar Driver
		// Tratamento para as licenças
		licenseStrings := strings.Split(fieldMap["licenses"], ",")
		licenses := make([]struct{ License string }, len(licenseStrings))
		for i, licenseStr := range licenseStrings {
			licenses[i].License = strings.TrimSpace(licenseStr)
		}

		// Criar Driver
		drive := driver.Driver{
			HiringType:        fieldMap["hiringType"],
			License:           fieldMap["license"],
			LicenseRegister:   fieldMap["licenseRegister"],
			LicenseExpedition: fieldMap["licenseExpedition"],
			LicenseExpiration: fieldMap["licenseExpiration"],
			Licenses: []struct {
				License string `json:"license"`
			}(licenses),
			Name:         fieldMap["driverName"],
			FirstLicense: fieldMap["firstLicense"],
			BirthDate:    fieldMap["birthDate"],
			Registration: fieldMap["registration"],
			Status:       fieldMap["status"],
			Type:         fieldMap["driverType"],
		}
		drive.DriverTeam.ID = driverTeamID

		driverID := driver.PostDriver(drive)
		if driverID == 0 {
			log.Fatal("Falha ao criar Driver")
		}

		// Create Group
		groupp := group.Group{
			Disabled: fieldMap["groupDisabled"] == "true",
			IsAdmin:  fieldMap["groupIsAdmin"] == "true",
			Name:     fieldMap["groupName"],
		}
		groupp.Organization.ID = orgID

		groupID := group.PostGroup(groupp)
		if groupID == 0 {
			log.Fatal("Falha ao criar Group")
		}

		fmt.Printf("Organização criada com sucesso. Organização_ID (%d), DriverTeam_ID (%d), Driver_ID (%d), e Group_ID (%d)\n", orgID, driverTeamID, driverID, groupID)
	}
}

// Converte uma string para um inteiro.
func parseInt(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

// Converte um slice de strings para um slice de inteiros, utilizando parseInt para cada elemento.
func parseIntArray(values []string) []int {
	ints := make([]int, len(values))
	for i, v := range values {
		ints[i] = parseInt(v)
	}
	return ints
}

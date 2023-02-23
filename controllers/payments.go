package controllers

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"project-alta-store/models"
	"strconv"

	"github.com/labstack/echo"
)

func TestPaymentsAPI(c echo.Context) error {
	return c.String(http.StatusOK, "Payments API. API is Active")
}

type initalOrganizationRecord struct {
	Index         int
	ID            string
	Name          string
	Website       string
	Country       string
	Description   string
	Founded       int
	Industry      string
	NumOfEmployee int
}

func createOrganizationList(data [][]string) []initalOrganizationRecord {
	var organizationList []initalOrganizationRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec initalOrganizationRecord
			for j, field := range line {

				if j == 0 {
					index, _ := strconv.Atoi(field)
					rec.Index = index
				} else if j == 1 {
					rec.ID = field
				} else if j == 2 {
					rec.Name = field
				} else if j == 3 {
					rec.Website = field
				} else if j == 4 {
					rec.Country = field
				} else if j == 5 {
					rec.Description = field
				} else if j == 6 {
					yearFounded, _ := strconv.Atoi(field)
					rec.Founded = yearFounded
				} else if j == 7 {
					rec.Industry = field
				} else if j == 8 {
					valueNumOfEmployee, _ := strconv.Atoi(field)
					rec.NumOfEmployee = valueNumOfEmployee
				}
			}
			organizationList = append(organizationList, rec)
		}
	}
	return organizationList
}

func GetStatisticOrgByCountry(c echo.Context) error {

	f, err := os.Open("organizations-1000000.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert array to struct
	organizationList := createOrganizationList(data)

	j, _ := json.MarshalIndent(organizationList, "", "  ")

	var resOrganizationsByCountry []models.OrganizationsByCountry_response

	resAllOrg := models.Organization_response{
		Code:    200,
		Status:  "Success",
		Message: "Success",
		Data:    resOrganizationsByCountry,
	}
	return c.JSON(http.StatusOK, resAllOrg)

}

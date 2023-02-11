package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"project-alta-store/models"
	"strconv"

	"github.com/labstack/echo"
)

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

func OrganizationAPI(c echo.Context) error {

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

	// convert records to array of structs
	organizationList := createOrganizationList(data)

	j, _ := json.MarshalIndent(organizationList, "", "  ")

	var resOrganization []models.Organizations

	if err := json.Unmarshal([]byte(string(j)), &resOrganization); err != nil {
		panic(err)
	}

	res := models.Organization_response{
		Code:    200,
		Status:  "Success",
		Message: "Success",
		Data:    resOrganization,
	}
	return c.JSON(http.StatusOK, res)
}

func CreateOrganization(c echo.Context) error {
	var post_body models.Organization_post
	// var pts models.Organizations
	// var organizationStruct models.ListOrganization

	if e := c.Bind(&post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": e.Error(),
		})
	}
	if e := models.Validate.Struct(post_body); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "Error",
			"message": e.Error(),
		})
	}

	var organization models.Organizations
	organization.ID = post_body.ID
	organization.Name = post_body.Name
	organization.Website = post_body.Website
	organization.Country = post_body.Country
	organization.Description = post_body.Description
	organization.Founded = post_body.Founded
	organization.Industry = post_body.Industry
	organization.NumOfEmployee = post_body.NumOfEmployee

	file, err := os.Create("organizations-1000000.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	// Using WriteAll

	var head [9]string
	var row [9]string

	// values := [][]string{}

	row[1] = organization.ID
	row[2] = organization.Name
	row[3] = organization.Website
	row[4] = organization.Country
	row[5] = organization.Description
	row[6] = strconv.Itoa(organization.Founded)
	row[7] = organization.Industry
	row[8] = strconv.Itoa(organization.NumOfEmployee)

	head[1] = "Organization Id"
	head[2] = "Name"
	head[3] = "Website"
	head[4] = "Country"
	head[5] = "Description"
	head[6] = "Founded"
	head[7] = "Industry"
	head[8] = "Number of employees"

	// values = append(values, head)
	// values = append(values, row)

	fmt.Print("Data head ", head)
	fmt.Println()
	fmt.Print("Data row ", row)
	fmt.Println()
	// fmt.Print("Data values ", values)

	return c.JSON(http.StatusOK, models.Organization_response_single{
		Code:    200,
		Status:  "success",
		Message: "success add Organization",
	})
}

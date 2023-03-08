package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"project-alta-store/models"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/labstack/echo"
	// "github.com/go-gota/gota"
	// "github.com/kniren/gota"
)

func TestPaymentsAPI(c echo.Context) error {
	return c.JSON(http.StatusOK, "Payments API. API is Active")
}

func ReadCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func GetStatisticOrgByCountryAPI(c echo.Context) error {
	csvString, err := os.Open("organizations-1000000.csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// remember to close the file at the end of the program
	defer csvString.Close()

	var resStatsOrgByCountry []models.StatisticOrg

	df := dataframe.ReadCSV(csvString)
	selectedColumn := df.Select([]string{"Country"})
	groupByCountry := selectedColumn.GroupBy("Country")
	countByCountry := groupByCountry.Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"Country"})

	loadResultDf := countByCountry.Records()

	for i, v := range loadResultDf {
		if i != 0 {
			resStatsOrgByCountry = append(resStatsOrgByCountry, models.StatisticOrg{
				Country:       v[0],
				Country_COUNT: v[1],
			})
		}
	}

	finalDataStatsOrgByCountry := models.StatsOrgByCountry_response_single{
		Code:    200,
		Status:  "Success",
		Message: "Success",
		Data:    resStatsOrgByCountry,
	}
	return c.JSON(http.StatusOK, finalDataStatsOrgByCountry)

}

func GetTopTenOrgByNoeAPI(c echo.Context) error {
	csvString, err := os.Open("organizations-1000000.csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// remember to close the file at the end of the program
	defer csvString.Close()

	var resStatsTopTenOrgByNoe []models.StatisticOrgByNoe

	df := dataframe.ReadCSV(csvString)
	selectedColumn := df.Select([]string{"Name", "Number of employees"})

	sorted := selectedColumn.Arrange(
		dataframe.RevSort("Number of employees"), // Sort in ascending order
	)

	loadResultDf := sorted.Records()

	for i, v := range loadResultDf {
		fmt.Println(i, v)
		if i != 0 {
			noe, err := strconv.Atoi(v[1])

			if err != nil {
				fmt.Println("Error during conversion")
			}

			resStatsTopTenOrgByNoe = append(resStatsTopTenOrgByNoe, models.StatisticOrgByNoe{
				NameOrg:       v[0],
				NumOfEmployee: noe,
			})
		}
	}

	finalDataStatsTopTenOrgByNoe := models.StatsTopTenOrgByNoe_response_single{
		Code:    200,
		Status:  "Success",
		Message: "Success",
		Data:    resStatsTopTenOrgByNoe,
	}
	return c.JSON(http.StatusOK, finalDataStatsTopTenOrgByNoe)
}

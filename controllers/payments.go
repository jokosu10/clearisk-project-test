package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/labstack/echo"
	// "github.com/go-gota/gota"
	// "github.com/kniren/gota"
)

func TestPaymentsAPI(c echo.Context) error {
	return c.String(http.StatusOK, "Payments API. API is Active")
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
	// const iota = 7 // Untyped int.

	csvString, err := os.Open("organizations-1000000.csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// remember to close the file at the end of the program
	defer csvString.Close()

	df := dataframe.ReadCSV(csvString)
	selectedColumn := df.Select([]string{"Country", "Number of employees"})
	groupByCountry := selectedColumn.GroupBy("Country")                                                                         // Group by column "key1", and column "key2"
	countByCountry := groupByCountry.Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"Country"}) // Maximum value in column "values",  Minimum value in column "values2"

	loadResultDf := countByCountry.Records()

	return c.JSON(http.StatusOK, loadResultDf)
}


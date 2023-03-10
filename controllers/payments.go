package controllers

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"project-alta-store/models"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
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

func GetSumOfBalanceOfPaymentByPeriodAPI(c echo.Context) error {
	csvString, err := os.Open("balance-of-payments-september-2022.csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// remember to close the file at the end of the program
	defer csvString.Close()

	mean := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		return series.Floats(sum / float64(len(floats)))
	}

	df := dataframe.ReadCSV(csvString)

	if err != nil {
		log.Fatal(err)
	}

	var resSumBalanceOfPaymentsByPeriode []models.SumOfBalancePaymentsByPeriod
	selectedColumn := df.Select([]string{"Period"})

	finalSumBalanceOfPaymentsByPeriode := selectedColumn.Capply(mean)
	loadResultDf := finalSumBalanceOfPaymentsByPeriode.Records()

	for i, v := range loadResultDf {
		if i != 0 {
			resSumBalanceOfPaymentsByPeriode = append(resSumBalanceOfPaymentsByPeriode, models.SumOfBalancePaymentsByPeriod{
				SumOfBalanceOfPaymentByPeriod: v[0],
			})
		}
	}

	var res = models.Sumbalanceperiode_response_single{
		Code:    200,
		Status:  "Success",
		Message: "Success get data",
		Data:    resSumBalanceOfPaymentsByPeriode[0],
	}

	return c.JSON(http.StatusOK, res)
}

func GetSumOfBalanceOfPaymentByStatusAPI(c echo.Context) error {
	csvString, err := os.Open("balance-of-payments-september-2022.csv")

	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// remember to close the file at the end of the program
	defer csvString.Close()

	df := dataframe.ReadCSV(csvString)

	if err != nil {
		log.Fatal(err)
	}

	var resSumBalanceOfPaymentsByStatus []models.SumOfBalancePaymentsByStatus
	selectedColumn := df.Select([]string{"STATUS"}).GroupBy("STATUS")

	calcSumBalanceOfPaymentsByStatus := selectedColumn.Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"STATUS"}) // Maximum value in column "values",  Minimum value in column "values2"

	finalSumBalanceOfPaymentsByStatus := calcSumBalanceOfPaymentsByStatus.Records()

	// countByCountry := groupByCountry.Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"Country"})
	for i, v := range finalSumBalanceOfPaymentsByStatus {
		if i != 0 {
			resSumBalanceOfPaymentsByStatus = append(resSumBalanceOfPaymentsByStatus, models.SumOfBalancePaymentsByStatus{
				Status:       v[0],
				STATUS_COUNT: v[1],
			})
		}
	}

	var res = models.Sumbalancestatus_response_single{
		Code:    200,
		Status:  "Success",
		Message: "Success get data",
		Data:    resSumBalanceOfPaymentsByStatus,
	}

	return c.JSON(http.StatusOK, res)
}

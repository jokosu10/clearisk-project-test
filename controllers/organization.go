package controllers

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"project-alta-store/lib/utils"
	"project-alta-store/models"
	"strconv"
	"strings"

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

type Stringer interface {
	String() string
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

func GetOrganizations(c echo.Context) error {

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

	var resOrganization []models.Organizations

	if utils.StringIsNotNumber(c.QueryParam("id")) {
		id, _ := strconv.Atoi(c.QueryParam("id"))

		if err := json.Unmarshal([]byte(string(j)), &resOrganization); err != nil {
			panic(err)
		}

		for i, _ := range organizationList {
			var resOrgById = models.Organization_response_single{
				Code:    200,
				Status:  "Success",
				Message: "Find data",
				Data:    resOrganization[i],
			}

			if id == resOrgById.Data.Index {
				return c.JSON(http.StatusOK, resOrgById)
			}

		}

		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    400,
			Status:  "fail",
			Message: "invalid id supplied",
		})

	} else if len(c.QueryParam("id")) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{
			Code:    400,
			Status:  "fail",
			Message: "invalid id supplied",
		})
	} else {

		if err := json.Unmarshal([]byte(string(j)), &resOrganization); err != nil {
			panic(err)
		}

		resAllOrg := models.Organization_response{
			Code:    200,
			Status:  "Success",
			Message: "Success",
			Data:    resOrganization,
		}
		return c.JSON(http.StatusOK, resAllOrg)
	}
}

func addcol(fname string, column []string) error {
	// read the file
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}

	// add column
	l := len(lines)
	if len(column) < l {
		l = len(column)
	}
	for i := 0; i < l; i++ {
		lines[i] = append(lines[i], column[i])
	}

	// write the file
	f, err = os.Create(fname)
	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func CreateOrganization(c echo.Context) error {
	var post_body models.Organization_post

	f, err := os.Open("organizations-1000000.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	_, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

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

	var u models.Organizations
	u.Index = 8
	u.ID = post_body.ID
	u.Name = post_body.Name
	u.Website = post_body.Website
	u.Country = post_body.Country
	u.Description = post_body.Description
	u.Founded = post_body.Founded
	u.Industry = post_body.Industry
	u.NumOfEmployee = post_body.NumOfEmployee
	stringSlice := []string{strconv.FormatInt(int64(u.Index), 10), u.ID, u.Name, u.Website, u.Country, u.Description, strconv.FormatInt(int64(u.Founded), 10), u.Industry, strconv.FormatInt(int64(u.NumOfEmployee), 10)}
	stringByte := "\x00" + strings.Join(stringSlice, "\x20\x00") // x20 = space and x00 = null

	fmt.Println(stringSlice)
	fmt.Println(string([]byte(stringByte)))
	a := append(stringSlice, string([]byte(stringByte)))

	if err := addcol("organizations-1000000.csv", a); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "success add Organization",
	})
}

func DeleteOrganization(c echo.Context) error {

	f, err := os.Open("organizations-1000000.csv")
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()
	rows, err := readSample(f)
	organizationList := createOrganizationList(rows)

	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	// Using WriteAll
	var data [][]string
	for _, record := range organizationList {
		row := []string{record.ID, record.Name, record.Website, record.Country, record.Description, strconv.FormatInt(int64(record.Founded), 10), record.Industry, strconv.FormatInt(int64(record.NumOfEmployee), 10)}
		data = append(data, row)
	}
	w.WriteAll(data)

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "success delete Organization",
	})

}

func readSample(rs io.ReadSeeker) ([][]string, error) {
	// Skip first row (line)
	row1, err := bufio.NewReader(rs).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = rs.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Read remaining rows
	r := csv.NewReader(rs)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// func createKeyValuePairs(m map[string]string) string {
// 	b := new(bytes.Buffer)
// 	for key, value := range m {
// 		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
// 	}
// 	return b.String()
// }

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

// func (m models.Organizations) String() string {
// 	return fmt.Sprintf(
// 		"{ Index: %d, Organization Id : %s, Name : %s, Website : %s, Country : %s, Description : %s, Founded: %d, Industry : %s, Number of employees : %d }",
// 		m.ID, m.Name, m.Website, m.Country, m.Description, m.Founded, m.Industry, m.NumOfEmployee,
// 	)
// }

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
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	leg := len(data)

	fmt.Println("total no of rows:", leg)

	// convert array to struct
	// organizationList := createOrganizationList(data)
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

	// u := []models.Organizations{
	// 	{Index: 8, ID: post_body.ID, Name: post_body.Name, Website: post_body.Website, Country: post_body.Country, Description: post_body.Description, Founded: post_body.Founded, Industry: post_body.Industry, NumOfEmployee: post_body.NumOfEmployee},
	// }
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
	// index := make([]string, 0, 8)
	// id := post_body.ID
	// name := post_body.Name
	// website := post_body.Website
	// country := post_body.Country
	// description := post_body.Description
	// founded := make([]string, 0, post_body.Founded)
	// industry := post_body.Industry
	// noe := make([]string, 0, post_body.NumOfEmployee)
	// indexBro := leg + 1
	index := make([]string, 0, leg)
	id := make([]string, 0, leg)
	name := make([]string, 0, leg)
	website := make([]string, 0, leg)
	country := make([]string, 0, leg)
	description := make([]string, 0, leg)
	founded := make([]string, 0, leg)
	industry := make([]string, 0, leg)
	noe := make([]string, 0, leg)

	// fmt.Println("===============")
	// fmt.Println(u)

	// b, err := json.Marshal(u)
	// if err != nil {
	// 	panic(err)
	// }

	// var a interface{}
	// err = json.Unmarshal(b, &a)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// csvWriter := csv.NewWriter(f)
	for _, value := range data {
		if len(value) < 9 {
			continue // skip short records
		}

		// fmt.Println(len(value))
		index := append(index, strconv.FormatInt(int64(u.Index), 10))
		id := append(id, u.ID)
		name := append(name, u.Name)
		website := append(website, u.Website)
		country := append(country, u.Country)
		description := append(description, u.Description)
		founded := append(founded, strconv.FormatInt(int64(u.Founded), 10))
		industry := append(industry, u.Industry)
		noe := append(noe, strconv.FormatInt(int64(u.NumOfEmployee), 10))
		// fmt.Print("Field: %s\t Value: %v\n", typesOf.Field(i).Name, values.Field(i).Interface())
		// row := `index,`
		// fmt.Println(string{index})
		// fmt.Println(id)
		// fmt.Println(name)
		// ioutil.WriteFile("organizations-1000000.csv", []byte(row), 0666)

		// sIndex := strconv.FormatInt(int64(), 10)
		// sNoe := strconv.FormatInt(int64(record.NumOfEmployee), 10)
		// row := []string{strconv.FormatInt(int64(value), 10), value.ID, value.Name, value.Website, value.Country, value.Description, strconv.FormatInt(int64(value.Founded), 10), value.Industry, strconv.FormatInt(int64(value.NumOfEmployee), 10)}
		// // row := []string{value[0], value[1], value[2], value[3], value[4], value[5], value[6], value[7]}
		// // struct_v := fmt.Sprintf("%+v", row)
		// v, err := json.Marshal(row)
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Println(row)
		// fmt.Println("=================")
		// fmt.Println(string(v))

		// fmt.Println(string(out), err)
		// fmt.Println(struct_v)
		// var string resFinal = fmt.Println(struct_v)
		// w := csv.NewWriter(f)
		// fmt.Println(structStr) // { name: Yuto, age: 35 }
		// csvWriter.Write([]string{row[i]})
		// if err := w.Write(row); err != nil {
		//         log.Fatalln("error writing record to file", err)
		//     }
		//for i, v := range data {
		// x := []string{"column two", "a", "b", "c", "d"}
		// fmt.Println("==============start col=========")
		// fmt.Println([]string{string(createKeyValuePairs(u))})
		// fmt.Println(len(filedata))
		// fmt.Println(col)
		// fmt.Println([]string{string(createKeyValuePairs(u))})
		// fmt.Println([]string{string(v)})
		// fmt.Println(col)
		// fmt.Println("==============end col=========")
		// fmt.Println(v)
		// if len(col) < 6 {
		// 	continue // skip short records
		// }
		stringSlice := []string{index[0], id[0], name[0], website[0], country[0], description[0], founded[0], industry[0], noe[0]}
		stringByte := "\x00" + strings.Join(stringSlice, "\x20\x00") // x20 = space and x00 = null

		fmt.Println([]byte(stringByte))
		fmt.Println(string([]byte(stringByte)))
		ioutil.WriteFile("organizations-1000000.csv", []byte(stringByte), 0666)

		// if err := addcol(filepath, []string{string(v)}); err != nil {
		// 	panic(err)
		// }
		// var resOrgById = models.Organization_response_single{
		// 	Code:    200,
		// 	Status:  "Success",
		// 	Message: "Find data",
		// 	Data:    resOrganization[i],

	}
	// csvWriter.Flush()

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Code:    200,
		Status:  "success",
		Message: "success add Organization",
	})
}

package util

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"digital-house/pkg/web"

	"github.com/gin-gonic/gin"
)

func CheckError(sqlError error) error {
	/*switch {
	case strings.Contains(sqlError.Error(), "no rows in result set"):
		return fmt.Errorf("data not found")
	case strings.Contains(sqlError.Error(), "Duplicate entry"):
		err := strings.Split(sqlError.Error(), "'")
		msg := fmt.Sprint(err[3], " is unique, and ", err[1], " already registered")
		return fmt.Errorf(msg)
	case strings.Contains(sqlError.Error(), "Cannot add"):
		err := strings.Split(sqlError.Error(), "`")
		msg := fmt.Sprint(err[7], " is not registered on ", err[9])
		return fmt.Errorf(msg)
	case strings.Contains(sqlError.Error(), "sql: Scan error on column index 0, name \"locality_id\": converting NULL to string is unsupported"):
		return fmt.Errorf("")
	case strings.Contains(sqlError.Error(), "Cannot delete or update a parent row"):
		err := strings.Split(sqlError.Error(), "`")
		msg := fmt.Sprint("cannot delete the ", err[9], " row because it has a foreign key constraint on the ", err[3], " table")
		return fmt.Errorf(msg)
	}

	return sqlError
	func CheckError(sqlError error) error {*/
	switch {
	case strings.Contains(sqlError.Error(), "no rows in result set"):
		return fmt.Errorf("data not found")
	case strings.Contains(sqlError.Error(), "Duplicate entry"):
		err := strings.Split(sqlError.Error(), "'")
		msg := fmt.Sprint(err[3], " is unique, and ", err[1], " already registered")
		return fmt.Errorf(msg)
	case strings.Contains(sqlError.Error(), "1452"):
		err := strings.Split(sqlError.Error(), "`")
		msg := fmt.Sprint(err[7], " is not registered on ", err[9])
		return fmt.Errorf(msg)
	case strings.Contains(sqlError.Error(), "Cannot add or update a child row: a foreign key constraint fails"):
		msg := "foreign key constraint fails"
		return fmt.Errorf(msg)
	}
	return sqlError
}

func GetCurrentDateTime() string {
	currentTime := time.Now()
	return time.Date(currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		100,
		time.Local).Format("2006-01-02 15:04:05")
}

func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func IDChecker(ctx *gin.Context) (int, error) {
	var (
		id  int
		err error
	)
	switch {
	case ctx.Query("id") != "":
		id, err = strconv.Atoi(ctx.Query("id"))
	case ctx.Param("id") != "":
		id, err = strconv.Atoi(ctx.Param("id"))
	}
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "invalid ID"))
		return id, err
	}
	return id, nil
}

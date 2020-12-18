package models

import (
	"echo-framework/common"
	cm "echo-framework/common"
	"echo-framework/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

//var errMessage string

func FetchEmployees() (res Response, err error) {

	var errs *ex.AppError
	var cust cm.Employees
	var custObj []cm.Employees

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	con := db.CreateCon()

	sqlQuery := `SELECT
					IFNULL(LastName,'') LastName,
					IFNULL(FirstName,'') FirstName,
					IFNULL(Title,'') Title,
					IFNULL(Address,'') Address
				FROM employees`

	rows, err := con.Query(sqlQuery)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&cust.LastName, &cust.FirstName, &cust.Title, &cust.Address)

		if err != nil {
			errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
			errMessage = err.Error()
			return res, err
		}

		custObj = append(custObj, cust)

	}

	res.Status = http.StatusOK
	res.Message = "succses"
	res.Data = custObj

	return res, nil
}

//AddEmployee ...
func AddEmployee(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Employees)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `INSERT INTO employees (LastName,FirstName,Title,Address)
					 VALUES (?,?,?,?)`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.LastName, req.FirstName, req.Title, req.Address)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"Employees ADD": req.FirstName,
	}

	return res, nil
}

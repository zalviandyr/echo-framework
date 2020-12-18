package models

import (
	"echo-framework/common"
	"echo-framework/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

var errMessage string
var errs *ex.AppError
var cust common.Customers
var custObj []common.Customers

func FetchCustomers() (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	con := db.CreateCon()

	sqlQuery := `SELECT
					CustomerID,
					IFNULL(CompanyName,''),
					IFNULL(ContactName,'') ContactName,
					IFNULL(ContactTitle,'') ContactTitle,
					IFNULL(Address,'') Address,
					IFNULL(City,'') City,
					IFNULL(Country,'') Country,
					IFNULL(Phone,'') Phone ,
					IFNULL(PostalCode,'') PostalCode
				FROM customers `

	rows, err := con.Query(sqlQuery)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&cust.CustomerID, &cust.CompanyName, &cust.ContactName, &cust.ContactTitle, &cust.Address, &cust.City,
			&cust.Country, &cust.Phone, &cust.PostalCode)

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

//StoreCustomer ...
func StoreCustomer(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Customers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `INSERT INTO customers (CustomerID,CompanyName,ContactName,ContactTitle,Address,City,Country,Phone,PostalCode)
					 VALUES (?,?,?,?,?,?,?,?,?)`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CustomerID, req.CompanyName, req.ContactName, req.ContactTitle, req.Address,
		req.City, req.Country, req.Phone, req.PostalCode)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"CustomerID ADD": req.CustomerID,
	}

	return res, nil
}

//UpdateCustomer ...
func UpdateCustomer(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Customers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `UPDATE customers SET CompanyName = ?, ContactName = ?, ContactTitle = ? WHERE  CustomerID = '` + req.CustomerID + `'`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"row affected :": req.CustomerID,
	}

	return res, nil
}

//DeleteCustomer ...
func DeleteCustomer(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Customers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `DELETE FROM customers WHERE  CustomerID = ?`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CustomerID)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"row deleted :": req.CustomerID,
	}

	return res, nil
}

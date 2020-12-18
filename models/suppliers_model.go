package models

import (
	"echo-framework/common"
	"echo-framework/db"
	"net/http"

	"github.com/labstack/echo"
)

var suppliers common.Suppliers
var suppliersObj []common.Suppliers

func FetchSuppliers() (res Response, err error) {
	con := db.CreateCon()

	sqlQuery := `SELECT
					SupplierID,
					IFNULL(CompanyName, ''),
					IFNULL(ContactName, ''),
					IFNULL(ContactTitle, '')
				FROM suppliers`

	rows, err := con.Query(sqlQuery)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&suppliers.SupplierID, &suppliers.CompanyName, &suppliers.ContactName, &suppliers.ContactTitle)

		if err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = err.Error()
			res.Data = nil
			return res, err
		}

		suppliersObj = append(suppliersObj, suppliers)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = suppliersObj

	return res, nil
}

func StoreSupplier(e echo.Context) (res Response, err error) {
	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()
		res.Data = nil

		return res, err
	}

	con := db.CreateCon()
	sqlStatement := `INSERT INTO suppliers(CompanyName, ContactName, ContactTitle) VALUES(?,?,?)`
	stmt, err := con.Prepare(sqlStatement)
	_, err = stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle)

	if err != nil {
		res.Status = http.StatusConflict
		res.Message = err.Error()
		res.Data = nil
	} else {
		defer stmt.Close()
		res.Status = http.StatusCreated
		res.Message = "Success created"
		res.Data = nil
	}

	return res, nil
}

func UpdateSupplier(e echo.Context) (res Response, err error) {
	supplierID := e.Param("supplierID")
	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		res.Status = http.StatusBadRequest
		res.Message = err.Error()
		res.Data = nil

		return res, err
	}

	con := db.CreateCon()
	sqlStatement := `UPDATE suppliers SET CompanyName = ?, ContactName = ?, ContactTitle = ?
					WHERE SupplierID = ?`
	stmt, err := con.Prepare(sqlStatement)
	_, err = stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle, supplierID)

	if err != nil {
		res.Status = http.StatusConflict
		res.Message = err.Error()
		res.Data = nil
	} else {
		defer stmt.Close()
		res.Status = http.StatusCreated
		res.Message = "Success update"
		res.Data = nil
	}

	return res, nil
}

func DeleteSupplier(e echo.Context) (res Response, err error) {
	supplierID := e.Param("supplierID")

	con := db.CreateCon()
	sqlStatement := `DELETE FROM suppliers WHERE SupplierID = ?`
	stmt, err := con.Prepare(sqlStatement)
	_, err = stmt.Exec(supplierID)

	if err != nil {
		res.Status = http.StatusConflict
		res.Message = err.Error()
		res.Data = nil
	} else {
		defer stmt.Close()
		res.Status = http.StatusCreated
		res.Message = "Success delete"
		res.Data = nil
	}

	return res, nil
}

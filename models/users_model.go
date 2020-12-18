package models

import (
	"database/sql"
	"echo-framework/common"
	"echo-framework/db"
	"echo-framework/helpers"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

var user common.Users
var userObj []common.Users

//FetchUsers ...
func FetchUsers() (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	con := db.CreateCon()

	sqlQuery := `SELECT
					id,
					IFNULL(first_name,'') first_name,
					IFNULL(last_name,'') last_name,
					IFNULL(email,'') email,
					IFNULL(username,'') username,
					IFNULL(password,'') password
				FROM users `

	rows, err := con.Query(sqlQuery)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&user.ID, &user.NamaDepan, &user.NamaBelakang, &user.Email, &user.Username, &user.Password)

		if err != nil {
			errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
			errMessage = err.Error()
			return res, err
		}

		userObj = append(userObj, user)

	}

	res.Status = http.StatusOK
	res.Message = "succses"
	res.Data = userObj

	return res, nil
}

//StoreUser ...
func StoreUser(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Users)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `INSERT INTO users (nama_depan,nama_belakang,email,username,password)
					 VALUES (?,?,?,?,?)`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	newPassword, _ := helpers.HashPassword(req.Password)

	result, err := stmt.Exec(req.NamaDepan, req.NamaBelakang, req.Email, req.Username, newPassword)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"id ADD": req.Username,
	}

	return res, nil
}

//UpdateUser ...
func UpdateUser(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Users)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `UPDATE users SET nama_depan = ?, nama_belakang = ?, email = ? WHERE  id = ?`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.NamaDepan, req.NamaBelakang, req.Email, req.ID)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]int{
		"row affected :": req.ID,
	}

	return res, nil
}

//DeleteUser ...
func DeleteUser(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Users)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `DELETE FROM users WHERE  id = ?`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.ID)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]int{
		"row deleted :": req.ID,
	}

	return res, nil
}

//CheckUser ...
func CheckUser(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	var pwd string

	req := new(common.Users)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlQuery := `SELECT
					id,
					IFNULL(username,'') username,
					IFNULL(password,'') password
				FROM users WHERE username = ? `

	row := con.QueryRow(sqlQuery, req.Username).Scan(
		&user.ID, &user.Username, &pwd,
	)

	if row == sql.ErrNoRows {
		errs = ex.Errorc(http.StatusInternalServerError).Rem("No Result")
		res.Status = errs.ErrCode
		res.Message = errs.Remark
		return res, nil
	}

	if row != nil {
		errs = ex.Errorc(http.StatusInternalServerError).Rem("Query Error")
		res.Status = errs.ErrCode
		res.Message = "Query Error"
		return res, nil
	}

	match, err := helpers.CheckPasswordHash(req.Password, pwd)

	if !match {
		errs = ex.Errorc(http.StatusInternalServerError).Rem("password doesn't mactch")
		res.Status = errs.ErrCode
		res.Message = "password doesn't mactch"
		return res, nil
	}

	username := req.Username

	res.Status = 200
	res.Message = "Login Succeced"
	res.Data = map[string]string{
		"username": username,
	}

	return res, nil
}

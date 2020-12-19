## Installation

-   `git clone` this repository
-   Open with vscode
-   Exec `go env -w GO111MODULE=on` to activate go module
-   Exec `go mod vendor` for install all depedencies

## Run

```bash
> go run main.go
```

## Tutorial

1. Buat `struct` dengan nama `Suppliers` difile `common.go`
    ```go
    type Suppliers struct {
        SupplierID   string `json:"SupplierID,omitempty"`
        CompanyName  string `json:"CompanyName"`
        ContactName  string `json:"ContactName"`
        ContactTitle string `json:"ContactTitle"`
    }
    ```
2. Buat models dengan nama file `suppliers_model.go` di folder models. Fungsi dari models tersebut adalah untuk melakukan operasi CRUD pada database.
3. Buat fungsi dengan nama `FetchSuppliers`. Fungsi tersebut adalah untuk mengambil data dari database dan mengkonversi menjadi `struct` yang telah dibuat difile `common.go`

    ```go
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
    ```

4. Setelah kita membuat models, maka dari itu perlu `controller` untuk berinteraksi dengan `models` yang akan nantinya mengammbil fungsi-fungsi di`models`.
5. Buat file `suppliers_controller.go` difolder `controllers`.
6. Buat fungsi difile tersebut dengan nama `FetchAllSuppliers`. Fungsi tersebut adalah untuk mengakses fungsi yang ada pada models dan mengkonversikan kedalam `JSON` atau bisa disebut juga `view`.

    ```go
    func FetchAllSuppliers(c echo.Context) (err error) {
        result, err := models.FetchSuppliers()

        return c.JSON(http.StatusOK, result)
    }
    ```

7. Dengan adanya `controller` maka kita akan daftarkan ke `route` agar bisa diakses.
8. Buat file `suppliers_route.go`. Fungsi dari file ini adalah untuk me-route kan semua fungsi-fungsi yang ada pada `suppliers_controller`.
9. Buat fungsi dengan nama `SuppliersRoute`
    ```go
    func SuppliersRoute(g *echo.Group) {
        g.GET("/list", controllers.FetchAllSuppliers)
    }
    ```
10. Setelah `suppliers` memiliki route maka kita akan daftar `suppliers_route.go` kedalam main route, yaitu `routes.go`
11. Tinggal tambahkan baris dibawah dibagian bawah fungsi `Init`.
    ```go
    //SuppliersRoute ...
    SuppliersRoute(e.Group("/suppliers"))
    ```
12. Dan untuk mencek apakah telah berjalan, akses `127.0.0.1/suppliers/list`

## Request

-   List (GET)

    ```
    127.0.0.1:3000/suppliers/list
    ```

-   Add (POST)

    ```
    127.0.0.1:3000/suppliers/add
    ```

-   Update (PUT)

    ```
    127.0.0.1:3000/suppliers/update/:supplierID
    ```

-   Delete (DELETE)

    ```
    127.0.0.1:3000/suppliers/delete/:supplierID
    ```

## Response

-   List

    ```json
    {
        "status": 200,
        "message": "Success",
        "data": [
            {
                "SupplierID": "1",
                "CompanyName": "Exotic Liquids",
                "ContactName": "Charlotte Cooper",
                "ContactTitle": "Purchasing Manager"
            },
            ...
        ]
    }
    ```

-   Add

    ```json
    {
        "status": 201,
        "message": "Success created"
    }
    ```

-   Update

    ```json
    {
        "status": 201,
        "message": "Success update"
    }
    ```

-   Delete
    ```json
    {
        "status": 201,
        "message": "Success delete"
    }
    ```

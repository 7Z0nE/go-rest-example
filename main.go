package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	router := gin.Default()

	router.GET("/", getRoot)

	router.GET("/input/:parameterName", getInput)

	router.GET("/products", getProducts)

	router.Run(":8080")
}

func getInput(ctx *gin.Context) {
	input := ctx.Param("parameterName")

	fmt.Println(input)

	ctx.String(http.StatusOK, "Success, %s", input)
}

func getRoot(ctx *gin.Context) {
	ctx.String(http.StatusOK, "This is a REST API for the dummy db.")
}

func getProducts(ctx *gin.Context) {
	db, err := sql.Open("mysql", "dummy@/dummy")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Could not open database connection.")
		return
	}
	if db.Ping() != nil {
		ctx.String(http.StatusInternalServerError, "Could not open database connection.")
		return
	}
	defer db.Close()

	filterName := ctx.Query("name")
	filterVendor := ctx.Query("vendor")

	stmt, err := db.Prepare("SELECT productID, productName, Vendors.vendorID, vendorName FROM Products INNER JOIN Vendors ON Vendors.vendorID=Products.vendorID WHERE productName=? AND vendorName=?")

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to prepare query statement.")
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(filterName)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to run query")
		return
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {

		var id, vendorID int
		var name, vendorName string

		rows.Scan(&id, &name, &vendorID, &vendorName)

		products = append(products, Product{
			Id:   id,
			Name: name,
			Vendor: Vendor{
				Id:   vendorID,
				Name: vendorName}})
	}
	fmt.Println(products)

	ctx.JSON(http.StatusOK, products)
}

package main

import _ "github.com/denisenkom/go-mssqldb"
import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var db *sql.DB

var server = "localhost"
var port = 1433
var user = "sa"
var password = "Geometry54813790"
var database = "mie"

type Sample struct {
	SampleUNo  int
	SampleName sql.NullString
	StorePlace sql.NullString
}

func main() {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	//Create connection pool
	db, err = sql.Open("sqlserver", connString)

	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	//Read Sample
	result, err := ReadSample()
	if err != nil {
		log.Fatal("Error reading Sample:", err.Error())
	}

	fmt.Printf("Read %d row(s) successfully.\n", result)

	r := gin.Default()
	r.GET("/rick", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("你好羅元佑!"),
			"data":    result,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ReadSample() ([]Sample, error) {
	ctx := context.Background()

	// check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	tsql := fmt.Sprintf(`SELECT SampleUNo, SampleName, StorePlace FROM dbo.Sample;`)

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []Sample
	// Iterate through the result set.
	for rows.Next() {
		var each = Sample{}
		// Get values from row.
		err := rows.Scan(&each.SampleUNo, &each.SampleName, &each.StorePlace)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}

	return result, nil
}

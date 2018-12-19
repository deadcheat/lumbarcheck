package main

import (
	"fmt"
	"os"
	"time"

	"github.com/deadcheat/lumbarcheck/checker"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	dataSourceType = kingpin.Flag("type", `Kind of Data Source, listed below
 mysql connect to mysql database
 http connect to given URL with GET Method
`).Short('t').Default("mysql").String()
	ds = kingpin.Arg("Data Source", "Data Source Name you'll try to connect to").Required().String()

	startUpTextFormat = `Start checking using given information printed below,
DataBase      : %s
DataSourceName: %s
`

	failedTextFormat = `Failed to connect to datasource,
Error: %+v
`
	succeededText = `Succeeded to connect to datasource,
Time: %s
`
)

func main() {
	kingpin.Parse()

	fmt.Printf(startUpTextFormat, *dataSourceType, *ds)

	c, err := checker.Create(*dataSourceType)
	if err != nil {
		fmt.Printf(failedTextFormat, err)
		os.Exit(1)
	}
	timeToStart := time.Now()
	err = c.Check(*ds)
	timeToEnd := time.Now()
	if err != nil {
		fmt.Printf(failedTextFormat, err)
		os.Exit(1)
	}
	fmt.Printf(succeededText, timeToEnd.Sub(timeToStart).String())
}

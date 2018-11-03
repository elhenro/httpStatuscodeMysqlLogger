package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func MysqlAddScanEntry(url string, code string, rptime string ) (bool, error){
	c := Join(mysqlUser, ":", mysqlPass, "@/", "api")
	db, err := sql.Open("mysql", c)	
	defer db.Close()

	stmtIns, err := db.Prepare(Join("INSERT INTO testStatusRequests ( id, url, date, code, rptime ) VALUES( ?, ?, NOW(), ?, ? )"))
	defer stmtIns.Close() 

	_, err = stmtIns.Exec( mysqlGetNewHighestID("api","testStatusRequests"), url, code, rptime)
	
	return true, err
}


func mysqlGetNewHighestID(database string, row string) int {
	var q int
	c := Join(mysqlUser, ":", mysqlPass, "@/", database)
	db, err := sql.Open("mysql", c)
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()	
	results, err := db.Query( Join("SELECT id FROM ", row, " ORDER BY id DESC LIMIT 0, 1;"))
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var r ScanEntry
		err = results.Scan(&r.id)
		if err != nil {
			panic(err.Error())
		}
		q = r.id + 1
	}
	return q
}

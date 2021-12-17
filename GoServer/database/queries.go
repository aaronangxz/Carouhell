package database

import (
	"io/ioutil"
	"log"
)

const (
	DB_GET_LISTING_QUERY_WITH_CUSTOM_CONDITION = 0
	DB_LISTING_FIXED_QUERY                     = 1
	DB_GET_LISTING_LOGGEDIN_QUERY              = 2
)

var DB_query_name = map[int32]string{
	0: "get_listing_query_with_custom_condition.sql",
	1: "listing_fixed_query.sql",
	2: "get_listing_loggedin_query.sql",
}

var DB_query_value = map[string]int32{
	"get_listing_query_with_custom_condition.sql": 0,
	"listing_fixed_query.sql":                     1,
	"get_listing_loggedin_query.sql":              2,
}

func LoadSqlQuery(queryFile int32) string {

	// Read file
	file, err := ioutil.ReadFile("./database/" + DB_query_name[queryFile])
	if err != nil {
		log.Printf("Error during LoadSqlQuery: %v | %v", DB_query_name[queryFile], err.Error())
	}

	return string(file)
}

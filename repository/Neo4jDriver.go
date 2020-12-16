package repository

import (
	"log"
	"fmt"
	
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	
	"GoRest/configuration"
)

var driver neo4j.Driver

func CreateDriver(ConfigInfo *configuration.Neo4jConfiguration) {
	var err error
	driver, err = neo4j.NewDriver(ConfigInfo.Url, neo4j.BasicAuth(ConfigInfo.Username, ConfigInfo.Password, ""))
	if err != nil {
		log.Fatal(fmt.Errorf("Could not find database: %w", err))
	}
}

func CloseDriver() {
	if err := driver.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}


package configuration

import (
	"os"
)

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
}

var configuration Neo4jConfiguration = Neo4jConfiguration {
	Url: "bolt://localhost:7688",
	Username: "neo4j",
	Password: "admin",
}

func ParseConfiguration() *Neo4jConfiguration {
	return &Neo4jConfiguration{
		Url:      lookupEnvOrGetDefault("NEO4J_URI", configuration.Url),
		Username: lookupEnvOrGetDefault("NEO4J_USER", configuration.Username),
		Password: lookupEnvOrGetDefault("NEO4J_PASSWORD", configuration.Password),
	}
}

func lookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}

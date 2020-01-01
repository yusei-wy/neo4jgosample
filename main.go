package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"neo4jgosample/graphdb"

	"github.com/gin-gonic/gin"
)

type Secret struct {
	Neo4j Neo4j `json:"neo4j"`
}

type Neo4j struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const PORT = 8080

func main() {
	secret, err := getSecret()
	if err != nil {
		panic(err)
	}

	driver, err := graphdb.GetNeo4j(secret.Neo4j.Username, secret.Neo4j.Password)
	if err != nil {
		panic(err)
	}
	defer driver.Close()
	log.Println("connected")

	router := gin.Default()

	router.GET("/", handleHome)

	router.Run(fmt.Sprintf(":%d", PORT))
}

func handleHome(c *gin.Context) {
	c.String(http.StatusOK, "HELLO WORLD")
}

func getSecret() (*Secret, error) {
	row, err := ioutil.ReadFile("./secret/secret.json")
	if err != nil {
		return nil, err
	}

	var secret *Secret
	if err := json.Unmarshal(row, &secret); err != nil {
		return nil, err
	}

	return secret, nil
}

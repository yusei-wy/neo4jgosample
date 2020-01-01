package graphdb

import "github.com/neo4j/neo4j-go-driver/neo4j"

func GetNeo4(username string, password string) (neo4j.Driver, error) {
	var (
		driver neo4j.Driver
		err    error
	)

	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth(username, password, "")); err != nil {
		return nil, err // handle error
	}

	return driver, nil
}

func GetSession(driver neo4j.Driver) (neo4j.Session, error) {
	var (
		session neo4j.Session
		err     error
	)
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return nil, err
	}
	return session, nil
}

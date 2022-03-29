package graphdb

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

func NewNeo4jSession(uri, database, username, password string) (neo4j.Session, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead, DatabaseName: database})
	return session, nil
}

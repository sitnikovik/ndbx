package neo4j

import "errors"

// Auth holds the authentication credentials for the Neo4j database.
type Auth struct {
	// usr is the Neo4j database username
	usr string
	// pwd is the Neo4j database password
	pwd string
}

// NewAuth creates a new Auth instance
// with the given username and password.
func NewAuth(usr, pwd string) Auth {
	return Auth{
		usr: usr,
		pwd: pwd,
	}
}

// Username returns the Neo4j database username.
func (a Auth) Username() string {
	return a.usr
}

// Password returns the Neo4j database password.
func (a Auth) Password() string {
	return a.pwd
}

// Empty returns true if both username and password are empty.
func (a Auth) Empty() bool {
	return a.usr == "" && a.pwd == ""
}

// Validate validates the Auth credentials.
//
// If both username and password are empty, nil is returned
// because it is considered as no authentication.
func (a Auth) Validate() error {
	if a.Empty() {
		return nil
	}
	if a.usr == "" {
		return errors.New("empty username")
	}
	if a.pwd == "" {
		return errors.New("empty password")
	}
	return nil
}

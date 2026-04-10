package cassandra

// Auth represents the authentication credentials for Cassandra.
type Auth struct {
	// usr is username to use for Cassandra authentication.
	usr string
	// pwd is password to use for Cassandra authentication.
	pwd string
}

// NewAuth creates a new Auth instance.
func NewAuth(
	usr string,
	pwd string,
) Auth {
	return Auth{
		usr: usr,
		pwd: pwd,
	}
}

// Username returns the username to use for Cassandra authentication.
func (a Auth) Username() string {
	return a.usr
}

// Password returns the password to use for Cassandra authentication.
func (a Auth) Password() string {
	return a.pwd
}

// Empty defines whether the authentication credentials are empty.
func (a Auth) Empty() bool {
	return a.usr == "" && a.pwd == ""
}

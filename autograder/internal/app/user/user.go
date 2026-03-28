package user

// User represents a user in the autograder application.
type User struct {
	// fullName is the full name of the user.
	fullName string
	// username is the unique username of the user.
	username string
	// id is the unique identifier of the user.
	id ID
}

// NewUser creates and returns a new User instance.
func NewUser(
	id ID,
	username string,
	fullName string,
) User {
	return User{
		id:       id,
		fullName: fullName,
		username: username,
	}
}

// ID returns the unique identifier of the user.
func (u User) ID() ID {
	return u.id
}

// FullName returns the full name of the user.
func (u User) FullName() string {
	return u.fullName
}

// Username returns the unique username of the user.
func (u User) Username() string {
	return u.username
}

// Idendity returns the user's identity.
func (u User) Idendity() Identity {
	return NewIdentity(
		u.ID(),
		WithUsername(
			u.Username(),
		),
	)
}

// Hash returns hash representaion of the user.
func (u User) Hash() string {
	return u.Idendity().Hash()
}

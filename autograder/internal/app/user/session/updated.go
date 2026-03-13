package session

// Updated checks if the session has been updated since
// it was created by comparing the updatedAt and createdAt timestamps.
func (s Session) Updated() bool {
	a := s.Dates().UpdatedAt()
	b := s.Dates().CreatedAt()
	return a.After(b)
}

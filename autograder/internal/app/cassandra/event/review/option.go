package review

// Option represents a functional option
// to configure filtering the Reactions instance.
type Option func(*Reviews)

// WithLimit sets the limit to restrict
// the number of Reactions got from database.
func WithLimit(n int) Option {
	return func(r *Reviews) {
		r.limit = n
	}
}

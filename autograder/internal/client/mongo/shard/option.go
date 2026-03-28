package shard

// Option represents a functional option for configuring an Shard.
type Option func(s *Shard)

// WithCount sets count of records the shard.
func WithCount(n int) Option {
	return func(s *Shard) {
		s.count = n
	}
}

// WithOk sets the sign displays the working status of the Shard.
func WithOk(ok bool) Option {
	return func(s *Shard) {
		s.ok = ok
	}
}

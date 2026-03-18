package body

// Option represents a functional option for configuring the Body struct.
type Option func(b *Body)

// WithCategory sets the provided category to Body.
func WithCategory(cat string) Option {
	return func(b *Body) {
		b.category = cat
	}
}

// WithCity sets the provided city to Body.
func WithCity(city string) Option {
	return func(b *Body) {
		b.city = city
	}
}

// WithPrice sets the provided price to Body.
func WithPrice(price uint) Option {
	return func(b *Body) {
		b.price = price
	}
}

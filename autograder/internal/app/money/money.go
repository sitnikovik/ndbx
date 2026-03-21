package money

import "fmt"

// Money represents a cost to pay to get some.
type Money struct {
	// units is the unit part of the money (before comma).
	units uint64
	// nanos is the nanos part of the money (after comma).
	nanos uint8
}

// NewMoney creates a new Money instance.
func NewMoney(
	units uint64,
	nanos uint8,
) Money {
	return Money{
		units: units,
		nanos: nanos,
	}
}

// String returns a string representation on the money.
//
//	For example:
//	s := NewMoney(100.50).String() // 100.50
func (m Money) String() string {
	return fmt.Sprintf("%d.%d", m.units, m.nanos)
}

// Free defines if the money has no value.
func (m Money) Free() bool {
	return m.units == 0 && m.nanos == 0
}

// Units return the unit part of the money (before comma).
func (m Money) Units() uint64 {
	return m.units
}

// Nanos return the nanos part of the money (after comma).
func (m Money) Nanos() uint8 {
	return m.nanos
}

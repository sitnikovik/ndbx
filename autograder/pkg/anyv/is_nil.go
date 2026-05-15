package anyv

func (v Value) IsNil() bool {
	return v.raw == nil
}

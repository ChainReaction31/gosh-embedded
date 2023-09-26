package lib

type Device struct {
	uniqueID    string
	name        string
	description string
	location    [3]float64
}

func (d Device) isSystem() bool {
	return false
}

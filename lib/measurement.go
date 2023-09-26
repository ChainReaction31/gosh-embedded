package lib

type Measurement struct {
	name       string
	swe_type   string
	definition string
	label      string
	uom        string
	refFrame   string
	axisID     string
}

type VectorMeasurement struct {
	Measurement
	numCoords int
	coords    []Measurement
}

func (vec VectorMeasurement) addCoordinate(coord Measurement) {
	vec.coords = append(vec.coords, coord)
	vec.numCoords++
}

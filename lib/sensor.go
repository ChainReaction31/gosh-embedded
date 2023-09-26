package lib

type Sensor struct {
	Device
	numOutputs int
	outputs    []Measurement
}

func (s Sensor) addTimeStampUTC() {
	s.addTimeStamp(DEF_ISO8601, DEF_UTC)
}

func (s Sensor) addTimeStampOBC(uom string) {
	s.addTimeStamp(uom, "")
}

func (s Sensor) addTimeStamp(uom string, refFrame string) {
	measurement := Measurement{
		name:       "time",
		label:      "Sampling Time",
		definition: DEF_SAMPLING_TIME,
		uom:        uom,
		refFrame:   refFrame,
		swe_type:   TIME,
	}

	s.addMeasurement(measurement)
}

func (s Sensor) addLocationLLA(def string, label string) {
	vec := VectorMeasurement{
		Measurement: Measurement{
			name:       "loc",
			label:      label,
			definition: def,
			refFrame:   DEF_EPSG4979,
		},
		numCoords: 0,
		// this just makes an empty slice of Measurement(s)
		coords: make([]Measurement, 0),
	}
	vec.addCoordinate(Measurement{
		name:     "lat",
		label:    "Latitude",
		uom:      "deg",
		swe_type: QUANTITY,
		axisID:   "lat",
	})
	vec.addCoordinate(Measurement{
		name:     "lon",
		label:    "Longitude",
		uom:      "deg",
		swe_type: QUANTITY,
		axisID:   "lon",
	})
	vec.addCoordinate(Measurement{
		name:     "alt",
		label:    "Altitude",
		uom:      "m",
		swe_type: QUANTITY,
		axisID:   "alt",
	})
}

func (s Sensor) addMeasurement(measurement Measurement) {
	s.outputs = append(s.outputs, measurement)
	s.numOutputs++
}

func (s Sensor) isSystem() bool {
	return false
}

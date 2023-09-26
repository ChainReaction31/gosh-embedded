package lib

type System struct {
	Device
	sensors []Sensor
}

func (s System) isSystem() bool {
	return true
}

func (s System) addSensor(sensor ...Sensor) {
	s.sensors = append(s.sensors, sensor...)
}

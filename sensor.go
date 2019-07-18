package goev3

const (
	sensorAttrNumValues = "num_values"
	sensorAttrMode      = "mode"
	sensorAttrModes     = "modes"
)

const (
	irDriver = "lego-ev3-ir"
)

type SensorDescriptor struct {
	Name       string
	Path       string
	Address    string
	DriverName string
	Modes      []string
	Mode       string
	NumValues  int
}

func readSensorDescriptor(name string, nodePath string) (SensorDescriptor, error) {
	sd := SensorDescriptor{
		Name: name,
		Path: nodePath,
	}
	er := newAttrErrorReader(nodePath)
	sd.Address = er.readString(attrAddress)
	sd.DriverName = er.readString(attrDriverName)
	sd.Mode = er.readString(sensorAttrMode)
	sd.Modes = er.readStringSlice(sensorAttrModes)
	sd.NumValues = er.readInt(sensorAttrNumValues)
	return sd, er.error()
}

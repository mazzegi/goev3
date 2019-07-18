package goev3

const (
	colorReflectMode = "COL-REFLECT"
)

type ColorReflect struct {
	descriptor SensorDescriptor
}

func newColorReflect(d SensorDescriptor) (*ColorReflect, error) {
	cr := &ColorReflect{
		descriptor: d,
	}
	err := writeAttrString(d.Path, sensorAttrMode, colorReflectMode)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (cr *ColorReflect) Read() int {
	v, err := readAttrInt(cr.descriptor.Path, valueAttr(0))
	if err != nil {
		return 0
	}
	return v
}

package goev3

const (
	colorAmbientMode = "COL-AMBIENT"
)

type ColorAmbient struct {
	descriptor SensorDescriptor
}

func newColorAmbient(d SensorDescriptor) (*ColorAmbient, error) {
	cr := &ColorAmbient{
		descriptor: d,
	}
	err := writeAttrString(d.Path, sensorAttrMode, colorAmbientMode)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (cr *ColorAmbient) Read() int {
	v, err := readAttrInt(cr.descriptor.Path, valueAttr(0))
	if err != nil {
		return 0
	}
	return v
}

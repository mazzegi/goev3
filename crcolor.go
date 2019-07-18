package goev3

const (
	colorColorMode = "COL-COLOR"
)

type Color struct {
	descriptor SensorDescriptor
}

func newColor(d SensorDescriptor) (*Color, error) {
	cr := &Color{
		descriptor: d,
	}
	err := writeAttrString(d.Path, sensorAttrMode, colorColorMode)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

type ColorValue int

const (
	NoColor ColorValue = 0
	Black              = 1
	Blue               = 2
	Green              = 3
	Yellow             = 4
	Red                = 5
	White              = 6
	Brown              = 7
)

type ColorData struct {
	Value ColorValue
}

func (cd ColorData) String() string {
	switch cd.Value {
	case Black:
		return "black"
	case Blue:
		return "blue"
	case Green:
		return "green"
	case Yellow:
		return "yellow"
	case Red:
		return "red"
	case White:
		return "white"
	case Brown:
		return "brown"
	default:
		return "no-color"
	}
}

func (cr *Color) Read() ColorData {
	v, err := readAttrInt(cr.descriptor.Path, valueAttr(0))
	if err != nil {
		return ColorData{Value: NoColor}
	}
	return ColorData{Value: ColorValue(v)}
}

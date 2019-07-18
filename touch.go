package goev3

type Touch struct {
	descriptor SensorDescriptor
}

func newTouch(d SensorDescriptor) (*Touch, error) {
	t := &Touch{
		descriptor: d,
	}
	return t, nil
}

func (t *Touch) Touched() bool {
	v, err := readAttrInt(t.descriptor.Path, valueAttr(0))
	if err != nil {
		return false
	}
	switch v {
	case 1:
		return true
	default:
		return false
	}
}

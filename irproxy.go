package goev3

const (
	irProxyMode = "IR-PROX"
)

type IRProxy struct {
	descriptor SensorDescriptor
}

func newIRProxy(d SensorDescriptor) (*IRProxy, error) {
	irp := &IRProxy{
		descriptor: d,
	}
	err := writeAttrString(d.Path, sensorAttrMode, irProxyMode)
	if err != nil {
		return nil, err
	}
	return irp, nil
}

func (irp *IRProxy) Distance() int {
	v, err := readAttrInt(irp.descriptor.Path, valueAttr(0))
	cm := (float64(v) / 100.0) * 70.0
	if err != nil {
		return 0
	}
	return int(cm)
}

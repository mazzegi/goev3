package goev3

import "fmt"

const (
	irSeekMode = "IR-SEEK"
)

type IRSeek struct {
	descriptor SensorDescriptor
}

func newIRSeek(d SensorDescriptor) (*IRSeek, error) {
	irp := &IRSeek{
		descriptor: d,
	}
	err := writeAttrString(d.Path, sensorAttrMode, irSeekMode)
	if err != nil {
		return nil, err
	}
	return irp, nil
}

type IRSeekChannel struct {
	Heading  int //-25...25
	Distance int //0...100
}

type IRSeekData struct {
	Channels [4]IRSeekChannel
}

func (irsd IRSeekData) String() string {
	var s string
	for i, c := range irsd.Channels {
		s += fmt.Sprintf("(c: %d, h: %d, d: %d) ", i, c.Heading, c.Distance)
	}
	return s
}

//ReadData reads the channel data from the seek; for distance 100% is approximately 200cm, heading is -25 far left, 25 far right
func (irs *IRSeek) Seek() IRSeekData {
	data := IRSeekData{}
	for c := 0; c < 4; c++ {
		ch, _ := readAttrInt(irs.descriptor.Path, valueAttr(2*c))
		cd, _ := readAttrInt(irs.descriptor.Path, valueAttr(2*c+1))
		data.Channels[c].Heading = ch
		data.Channels[c].Distance = int((float64(cd) / 100.0) * 200.0)
	}
	return data
}

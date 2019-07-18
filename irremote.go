package goev3

import "fmt"

const (
	irRemoteMode = "IR-REMOTE"
)

type IRRemote struct {
	descriptor SensorDescriptor
}

func newIRRemote(d SensorDescriptor) (*IRRemote, error) {
	irr := &IRRemote{
		descriptor: d,
	}
	err := writeAttrString(d.Path, sensorAttrMode, irRemoteMode)
	if err != nil {
		return nil, err
	}
	return irr, nil
}

//IRRemoteCode is the returning chanel mode for ir remote mode
type IRRemoteCode int

//possible values for remote code
const (
	IRRCNone            IRRemoteCode = 0
	IRRCRedUp                        = 1
	IRRCRedDown                      = 2
	IRRCBlueUp                       = 3
	IRRCBlueDown                     = 4
	IRRCRedUpBlueUp                  = 5
	IRRCRedUpBlueDown                = 6
	IRRCRedDownBlueUp                = 7
	IRRCRedDownBlueDown              = 8
	IRRCBeaconOn                     = 9
	IRRCRedUpRedDown                 = 10
	IRRCBlueUpBlueDown               = 11
)

func (irrc IRRemoteCode) String() string {
	switch irrc {
	case IRRCRedUp:
		return "red-up"
	case IRRCRedDown:
		return "red-down"
	case IRRCBlueUp:
		return "blue-up"
	case IRRCBlueDown:
		return "blue-down"
	case IRRCRedUpBlueUp:
		return "red-up-blue-up"
	case IRRCRedUpBlueDown:
		return "red-up-blue-down"
	case IRRCRedDownBlueUp:
		return "red-down-blue-up"
	case IRRCRedDownBlueDown:
		return "red-down-blue-down"
	case IRRCBeaconOn:
		return "beacon-on"
	case IRRCRedUpRedDown:
		return "red-up-red-down"
	case IRRCBlueUpBlueDown:
		return "blue-up-blue-down"
	}
	return "none"
}

//IRRemoteData is the returning data structure for ir remote mode
type IRRemoteData struct {
	Channels [4]IRRemoteCode
}

func (irrd IRRemoteData) String() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", irrd.Channels[0], irrd.Channels[1], irrd.Channels[2], irrd.Channels[3])
}

func (irr *IRRemote) Read() IRRemoteData {
	data := IRRemoteData{}
	for c := 0; c < 4; c++ {
		code, _ := readAttrInt(irr.descriptor.Path, valueAttr(c))
		data.Channels[c] = IRRemoteCode(code)
	}

	return data
}

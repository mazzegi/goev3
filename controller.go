package goev3

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	attrAddress    = "address"
	attrDriverName = "driver_name"
	attrCommands   = "commands"
	attrModeAttr   = "mode"
)

//All out Addresses
const (
	OutA = "outA"
	OutB = "outB"
	OutC = "outC"
	OutD = "outD"
)

//All in Addresses
const (
	In1 = "in1"
	In2 = "in2"
	In3 = "in3"
	In4 = "in4"
)

type Controller struct {
	motorFS  string
	sensorFS string
	motors   []MotorDescriptor
	sensors  []SensorDescriptor
}

type ControllerOption func(ctrl *Controller) error

func WithMotorFS(fs string) ControllerOption {
	return func(api *Controller) error {
		api.motorFS = fs
		return nil
	}
}

func WithSensorFS(fs string) ControllerOption {
	return func(api *Controller) error {
		api.sensorFS = fs
		return nil
	}
}

func NewController(options ...ControllerOption) (*Controller, error) {
	ctrl := &Controller{
		motorFS:  "/sys/class/tacho-motor",
		sensorFS: "/sys/class/lego-sensor",
	}
	for _, option := range options {
		if err := option(ctrl); err != nil {
			return nil, errors.Wrap(err, "new-controller")
		}
	}
	ctrl.scanNodes(ctrl.motorFS, func(name, nodePath string) error {
		md, err := readMotorDescriptor(name, nodePath)
		if err != nil {
			return err
		}
		ctrl.motors = append(ctrl.motors, md)
		return nil
	})
	ctrl.scanNodes(ctrl.sensorFS, func(name, nodePath string) error {
		sd, err := readSensorDescriptor(name, nodePath)
		if err != nil {
			return err
		}
		ctrl.sensors = append(ctrl.sensors, sd)
		return nil
	})
	return ctrl, nil
}

func (ctrl *Controller) scanNodes(dir string, read func(name, nodePath string) error) error {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return errors.Wrapf(err, "read-dir failed (%s)", dir)
	}
	for _, fi := range fis {
		nodePath := filepath.Join(dir, fi.Name())
		err = read(fi.Name(), nodePath)
		if err != nil {
			return errors.Wrapf(err, "read motor descriptor (%s)", nodePath)
		}
	}
	return nil
}

func (ctrl *Controller) NewMotor(outAddr string) (*Motor, error) {
	for _, d := range ctrl.motors {
		if outAddr == d.Address {
			return newMotor(d), nil
		}
	}
	return nil, errors.Errorf("no motors at adress (%s)", outAddr)
}

func (ctrl *Controller) NewIRProxy(inAddr string) (*IRProxy, error) {
	for _, d := range ctrl.sensors {
		if d.Address == inAddr {
			if d.DriverName != irDriver {
				return nil, errors.Errorf("no ir-sensor at adress (%s) but (%s)", inAddr, d.DriverName)
			}
			return newIRProxy(d)
		}
	}
	return nil, errors.Errorf("no sensors at adress (%s)", inAddr)
}

func (ctrl *Controller) NewIRSeek(inAddr string) (*IRSeek, error) {
	for _, d := range ctrl.sensors {
		if d.Address == inAddr {
			if d.DriverName != irDriver {
				return nil, errors.Errorf("no ir-sensor at adress (%s) but (%s)", inAddr, d.DriverName)
			}
			return newIRSeek(d)
		}
	}
	return nil, errors.Errorf("no sensors at adress (%s)", inAddr)
}

func (ctrl *Controller) NewIRRemote(inAddr string) (*IRRemote, error) {
	for _, d := range ctrl.sensors {
		if d.Address == inAddr {
			if d.DriverName != irDriver {
				return nil, errors.Errorf("no ir-sensor at adress (%s) but (%s)", inAddr, d.DriverName)
			}
			return newIRRemote(d)
		}
	}
	return nil, errors.Errorf("no sensors at adress (%s)", inAddr)
}

package goev3

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

type MotorDescriptor struct {
	Name        string
	Path        string
	Address     string
	DriverName  string
	Commands    []string
	CountPerRot int
	MaxSpeed    int
}

type SensorDescriptor struct {
	Name       string
	Path       string
	Address    string
	DriverName string
	Modes      []string
	Mode       string
	NumValues  int
}

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
		md, err := ctrl.readMotorDescriptor(name, nodePath)
		if err != nil {
			return err
		}
		ctrl.motors = append(ctrl.motors, md)
		return nil
	})
	ctrl.scanNodes(ctrl.sensorFS, func(name, nodePath string) error {
		sd, err := ctrl.readSensorDescriptor(name, nodePath)
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

func (ctrl *Controller) readMotorDescriptor(name string, nodePath string) (MotorDescriptor, error) {
	md := MotorDescriptor{
		Name: name,
		Path: nodePath,
	}
	er := newAttrErrorReader(nodePath)
	md.Address = er.readString(attrAddress)
	md.DriverName = er.readString(attrDriverName)
	md.Commands = er.readStringSlice(attrCommands)
	md.CountPerRot = er.readInt(motorAttrCountPerRot)
	md.MaxSpeed = er.readInt(motorAttrMaxSpeed)
	return md, er.error()
}

func (ctrl *Controller) readSensorDescriptor(name string, nodePath string) (SensorDescriptor, error) {
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

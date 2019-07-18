package goev3

import "fmt"

const (
	batteryAttrCurrent    = "current_now"
	batteryAttrVoltage    = "voltage_now"
	batteryAttrVoltageMin = "voltage_min_design"
	batteryAttrVoltageMax = "voltage_max_design"
)

type Battery struct {
	CurrentMA  float64
	Voltage    float64
	VoltageMin float64
	VoltageMax float64
}

func readBattery(path string) (*Battery, error) {
	ar := newAttrErrorReader(path)
	b := &Battery{}
	b.CurrentMA = float64(ar.readInt(batteryAttrCurrent)) * 1e-3
	b.Voltage = float64(ar.readInt(batteryAttrVoltage)) * 1e-6
	b.VoltageMin = float64(ar.readInt(batteryAttrVoltageMin)) * 1e-6
	b.VoltageMax = float64(ar.readInt(batteryAttrVoltageMax)) * 1e-6
	return b, ar.error()
}

func (b Battery) String() string {
	return fmt.Sprintf("I=%.1f, U=%.1f (Umin=%.1f, Umax=%.1f)", b.CurrentMA, b.Voltage, b.VoltageMin, b.VoltageMax)
}

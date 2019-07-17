package goev3

import (
	"time"
)

const (
	runForeverCommand  = "run-forever"
	runToAbsPosCommand = "run-to-abs-pos"
	runToRelPosCommand = "run-to-rel-pos"
	runTimedCommand    = "run-timed"
	runDirectCommand   = "run-direct"
	stopCommand        = "stop"
	resetCommand       = "reset"
)

//List of all stop actions
const (
	StopActionCoast = "coast"
	StopActionBrake = "brake"
	StopActionHold  = "hold"
)

//List of all polarities actions
const (
	PolarityNormal   = "normal"
	PolarityInversed = "inversed"
)

//State wraps the state
type State string

//States wraps the state list
type States []State

//Possible values for state
const (
	StateIdle       State = ""
	StateRunning    State = "running"
	StateRamping    State = "ramping"
	StateHolding    State = "holding"
	StateOverloaded State = "overloaded"
	StateStalled    State = "stalled"
)

func (s State) String() string {
	switch s {
	case StateIdle:
		return "idle"
	}
	return string(s)
}

//StatesMoving contains all moving states
var StatesMoving = States{StateRunning, StateRamping}

//StatesNotMoving contains all non-moving states
var StatesNotMoving = States{StateIdle, StateHolding, StateOverloaded, StateStalled}

//ContainsOneOf return if states contains one of the specified ones
func (ts States) ContainsOneOf(s ...State) bool {
	for _, state := range ts {
		for _, cs := range s {
			if state == cs {
				return true
			}
		}
	}
	return false
}

//Motor is the base struct to control motors
type Motor struct {
	Descriptor MotorDescriptor
}

//newMotor creates a new motor at the given file path
func newMotor(md MotorDescriptor) *Motor {
	return &Motor{
		Descriptor: md,
	}
}

//ChangeSpeedSP changes the speed_sp to the specified value
func (tm *Motor) ChangeSpeedSP(val int) error {
	return writeAttrInt(tm.Descriptor.Path, motorAttrSpeedSP, val)
}

//SpeedSP return speed_sp
func (tm *Motor) SpeedSP() (int, error) {
	return readAttrInt(tm.Descriptor.Path, motorAttrSpeedSP)
}

//Speed return speed
func (tm *Motor) Speed() (int, error) {
	return readAttrInt(tm.Descriptor.Path, motorAttrSpeed)
}

//ChangePositionSP changes the position_sp to the specified value
func (tm *Motor) ChangePositionSP(val int) error {
	return writeAttrInt(tm.Descriptor.Path, motorAttrPositionSP, val)
}

//PositionSP returns position_sp
func (tm *Motor) PositionSP() (int, error) {
	return readAttrInt(tm.Descriptor.Path, motorAttrPositionSP)
}

//ChangeStopAction changes the stop action to the specified value
func (tm *Motor) ChangeStopAction(sa string) error {
	return writeAttrString(tm.Descriptor.Path, motorAttrStopAction, sa)
}

//StopAction returns stop_action
func (tm *Motor) StopAction() (string, error) {
	s, err := readAttrString(tm.Descriptor.Path, motorAttrStopAction)
	return s, err
}

//ChangePolarity changes the polarity to the specified value
func (tm *Motor) ChangePolarity(p string) error {
	return writeAttrString(tm.Descriptor.Path, motorAttrPolarity, p)
}

//Polarity returns the polarity
func (tm *Motor) Polarity() (string, error) {
	s, err := readAttrString(tm.Descriptor.Path, motorAttrPolarity)
	return s, err
}

//Position returns position
func (tm *Motor) Position() (int, error) {
	return readAttrInt(tm.Descriptor.Path, motorAttrPosition)
}

//ChangeRampUpSP changes the ramp_up_sp to the specified value
func (tm *Motor) ChangeRampUpSP(d time.Duration) error {
	return writeAttrInt(tm.Descriptor.Path, motorAttrRampUpSP, int(d.Seconds()*1000.0))
}

//RampUpSP returns ramp_up_sp
func (tm *Motor) RampUpSP() (time.Duration, error) {
	ms, err := readAttrInt(tm.Descriptor.Path, motorAttrRampUpSP)
	return time.Duration(ms) * time.Millisecond, err
}

//ChangeRampDownSP changes the ramp_down_sp to the specified value
func (tm *Motor) ChangeRampDownSP(d time.Duration) error {
	return writeAttrInt(tm.Descriptor.Path, motorAttrRampDownSP, int(d.Seconds()*1000.0))
}

//RampDownSP returns ramp_down_sp
func (tm *Motor) RampDownSP() (time.Duration, error) {
	ms, err := readAttrInt(tm.Descriptor.Path, motorAttrRampDownSP)
	return time.Duration(ms) * time.Millisecond, err
}

//ChangeTimeSP changes the time_sp to the specified value
func (tm *Motor) ChangeTimeSP(d time.Duration) error {
	return writeAttrInt(tm.Descriptor.Path, motorAttrTimeSP, int(d.Seconds()*1000.0))
}

//TimeSP returns time_sp
func (tm *Motor) TimeSP() (time.Duration, error) {
	ms, err := readAttrInt(tm.Descriptor.Path, motorAttrTimeSP)
	return time.Duration(ms) * time.Millisecond, err
}

//State returns state
func (tm *Motor) State() (States, error) {
	sl, err := readAttrStringSlice(tm.Descriptor.Path, motorAttrState)
	tmsl := States{}
	for _, s := range sl {
		tmsl = append(tmsl, State(s))
	}
	return tmsl, err
}

//Stop executes the stop command
func (tm *Motor) Stop() error {
	return writeAttrString(tm.Descriptor.Path, "command", stopCommand)
}

//Reset executes the reset command
func (tm *Motor) Reset() error {
	return writeAttrString(tm.Descriptor.Path, "command", resetCommand)
}

//RunForever executes the run-forever command
func (tm *Motor) RunForever() error {
	return writeAttrString(tm.Descriptor.Path, "command", runForeverCommand)
}

//RunToAbsPos executes the run-to-abs-pos command
func (tm *Motor) RunToAbsPos() error {
	return writeAttrString(tm.Descriptor.Path, "command", runToAbsPosCommand)
}

//RunToRelPos executes the run-to-rel-pos command
func (tm *Motor) RunToRelPos() error {
	return writeAttrString(tm.Descriptor.Path, "command", runToRelPosCommand)
}

//RunTimed executes the run-timed command
func (tm *Motor) RunTimed() error {
	return writeAttrString(tm.Descriptor.Path, "command", runTimedCommand)
}

//RunDirect executes the run-direct command
func (tm *Motor) RunDirect() error {
	return writeAttrString(tm.Descriptor.Path, "command", runDirectCommand)
}

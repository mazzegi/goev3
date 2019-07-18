package goev3

import (
	"time"

	"github.com/pkg/errors"
)

//Move collects several moves and actions of one motor
type Move struct {
	motor *Motor
	err   error
}

func (m *Move) Error() error {
	return m.err
}

//WithSpeed sets the speed for the next move
func (m *Move) WithSpeed(speed int) *Move {
	if m.err != nil {
		return m
	}
	m.err = m.motor.ChangeSpeedSP(speed)
	return m
}

//WithStopAction sets the stop-action
func (m *Move) WithStopAction(sa string) *Move {
	if m.err != nil {
		return m
	}
	m.err = m.motor.ChangeStopAction(sa)
	return m
}

//NewMove creates a new mover instance
func NewMove(motor *Motor) *Move {
	tm := &Move{
		motor: motor,
		err:   nil,
	}
	return tm
}

//ResetErrors clears all pending errors
func (m *Move) resetError() {
	m.err = nil
}

func (m *Move) waitForState(timeout time.Duration, waitState ...State) (States, error) {
	timer := time.NewTimer(timeout)
	ticker := time.NewTicker(20 * time.Millisecond)
	state, err := m.motor.State()
	if err != nil {
		return state, err
	}
	for {
		select {
		case <-timer.C:
			return state, errors.New("timeout")
		case <-ticker.C:
			state, err = m.motor.State()
			if err != nil {
				return state, errors.Wrap(err, "read-motor-state")
			}
			if state.ContainsOneOf(waitState...) {
				return state, nil
			}
		}
	}
}

//Some convenience functions

//ToRelPos moves to rel pos
func (m *Move) ToRelPos(timeoutMSec int, pos int) error {
	if m.err != nil {
		return m.err
	}
	m.err = m.motor.ChangePositionSP(pos)
	if m.err != nil {
		return m.err
	}
	m.err = m.motor.RunToRelPos()
	if m.err != nil {
		return m.err
	}
	_, m.err = m.waitForState(100*time.Millisecond, StatesMoving...)
	if m.err != nil {
		return m.err
	}
	_, m.err = m.waitForState(time.Duration(timeoutMSec)*time.Millisecond, StatesNotMoving...)
	if m.err != nil {
		return m.err
	}
	return m.err
}

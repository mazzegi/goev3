package goev3

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	attrAddress          = "address"
	attrDriverName       = "driver_name"
	attrCommands         = "commands"
	attrModeAttr         = "mode"
	motorAttrCountPerRot = "count_per_rot"
	motorAttrMaxSpeed    = "max_speed"
	motorAttrSpeedSP     = "speed_sp"
	motorAttrSpeed       = "speed"
	motorAttrPositionSP  = "position_sp"
	motorAttrPosition    = "position"
	motorAttrStopAction  = "stop_action"
	motorAttrPolarity    = "polarity"
	motorAttrRampUpSP    = "ramp_up_sp"
	motorAttrRampDownSP  = "ramp_down_sp"
	motorAttrTimeSP      = "time_sp"
	motorAttrState       = "state"
	sensorAttrNumValues  = "num_values"
	sensorAttrMode       = "mode"
	sensorAttrModes      = "modes"
)

func valueAttribute(idx int) string {
	return fmt.Sprintf("value%d", idx)
}

func readAttrRaw(path string, attr string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(path, attr))
}

func readAttrString(path string, attr string) (string, error) {
	b, err := readAttrRaw(path, attr)
	if err != nil {
		return "", errors.Wrapf(err, "read-attr-raw (%s) (%s)", path, attr)
	}
	return strings.Trim(string(b), "\n\r"), nil
}

func readAttrStringSlice(path string, attr string) ([]string, error) {
	s, err := readAttrString(path, attr)
	if err != nil {
		return nil, errors.Wrapf(err, "read-attr-string (%s) (%s)", path, attr)
	}
	return strings.Split(s, " "), nil
}

func readAttrInt(path string, attr string) (int, error) {
	s, err := readAttrString(path, attr)
	if err != nil {
		return 0, errors.Wrapf(err, "read-attr-string (%s) (%s)", path, attr)
	}
	return strconv.Atoi(s)
}

func writeAttrRaw(path string, attr string, b []byte) error {
	return ioutil.WriteFile(filepath.Join(path, attr), b, 0)
}

func writeAttrString(path string, attr string, s string) error {
	return writeAttrRaw(path, attr, []byte(s))
}

func writeAttrInt(path string, attr string, n int) error {
	return writeAttrString(path, attr, fmt.Sprintf("%d", n))
}

//Error Reader

type attrErrorReader struct {
	err  error
	path string
}

func newAttrErrorReader(path string) *attrErrorReader {
	return &attrErrorReader{
		err:  nil,
		path: path,
	}
}

func (r *attrErrorReader) error() error {
	return r.err
}

func (r *attrErrorReader) readString(attr string) string {
	if r.err != nil {
		return ""
	}
	s, err := readAttrString(r.path, attr)
	if err != nil {
		r.err = err
		return ""
	}
	return s
}

func (r *attrErrorReader) readStringSlice(attr string) []string {
	if r.err != nil {
		return nil
	}
	sl, err := readAttrStringSlice(r.path, attr)
	if err != nil {
		r.err = err
		return nil
	}
	return sl
}

func (r *attrErrorReader) readInt(attr string) int {
	if r.err != nil {
		return 0
	}
	n, err := readAttrInt(r.path, attr)
	if err != nil {
		r.err = err
		return 0
	}
	return n
}

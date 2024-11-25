package ulits

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type (
	LocalTime time.Time
	LocalDate time.Time
)

func NowTime() *LocalTime {
	n := LocalTime(time.Now())
	return &n
}

func NowDate() *LocalDate {
	n := LocalDate(time.Now())
	return &n
}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateTime))), nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	timeStr := strings.Trim(string(data), "\"")
	t1, err := time.ParseInLocation(time.DateTime, timeStr, time.Local)
	*t = LocalTime(t1)
	if err != nil {
		return err
	}
	return nil
}
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	if tlt.UnixMicro() == zeroTime.UnixMicro() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v any) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *LocalTime) String() string {
	return time.Time(*t).Format(time.DateTime)
}

func (t *LocalDate) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateOnly))), nil
}

func (t *LocalDate) UnmarshalJSON(data []byte) error {
	timeStr := strings.Trim(string(data), "\"")
	t1, err := time.ParseInLocation(time.DateOnly, timeStr, time.Local)
	*t = LocalDate(t1)
	if err != nil {
		return err
	}
	return nil
}
func (t LocalDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	if tlt.UnixMicro() == zeroTime.UnixMicro() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalDate) Scan(v any) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *LocalDate) String() string {
	return time.Time(*t).Format(time.DateOnly)
}

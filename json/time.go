package gjson

import (
	"time"
)

const secondTimeLayout = "2006-01-02 15:04:05"
type SecondTime struct {
	time.Time
}
func (t SecondTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(secondTimeLayout) + `"`), nil
}
func (t *SecondTime) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(`"` + secondTimeLayout + `"`, string(b))
	t.Time = v
	return err
}


const minuteTimeLayout = "2006-01-02 15:04"
type MinuteTime struct {
	time.Time
}
func (t MinuteTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(minuteTimeLayout) + `"`), nil
}
func (t *MinuteTime) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(`"` + minuteTimeLayout + `"`, string(b))
	t.Time = v
	return err
}

const hourTimeLayout = "2006-01-02 15"
type HourTime struct {
	time.Time
}
func (t HourTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(hourTimeLayout) + `"`), nil
}
func (t *HourTime) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(`"` + hourTimeLayout + `"`, string(b))
	t.Time = v
	return err
}

const dayTimeLayout = "2006-01-02"
type DayTime struct {
	time.Time
}
func (t DayTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(dayTimeLayout) + `"`), nil
}
func (t *DayTime) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(`"` + dayTimeLayout + `"`, string(b))
	t.Time = v
	return err
}

const monthTimeLayout = "2006-01"
type MonthTime struct {
	time.Time
}
func (t MonthTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(monthTimeLayout) + `"`), nil
}
func (t *MonthTime) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(`"` + monthTimeLayout + `"`, string(b))
	t.Time = v
	return err
}


const yearTimeLayout = "2006"
type YearTime struct {
	time.Time
}
func (t YearTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(yearTimeLayout) + `"`), nil
}
func (t *YearTime) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(`"` + yearTimeLayout + `"`, string(b))
	t.Time = v
	return err
}
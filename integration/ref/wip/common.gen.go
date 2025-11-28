package wip

import "time"

type DateWrapper struct {
	time.Time
}

func (w *DateWrapper) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(time.DateOnly, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	w.Time = t

	return nil
}

type TimeWrapper struct {
	time.Time
}

func (w *TimeWrapper) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(time.TimeOnly, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	w.Time = t

	return nil
}

type TimeWithTZWrapper struct {
	time.Time
}

func (w *TimeWithTZWrapper) UnmarshalJSON(b []byte) error {
	timeWithTZLayout := "15:04:05.000000Z07"
	t, err := time.Parse(timeWithTZLayout, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	w.Time = t

	return nil
}

type NoTimezoneWrapper struct {
	time.Time
}

func (w *NoTimezoneWrapper) UnmarshalJSON(b []byte) error {
	b[len(b)-1] = 'Z'
	t, err := time.Parse(time.RFC3339, string(b[1:]))
	w.Time = t

	return err
}

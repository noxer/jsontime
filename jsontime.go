package jsontime

import (
	"time"
)

// T represents a timestamp with second precision.
type T int64

// FromTime converts a time.Time to T. Be aware that the precision is only seconds.
func FromTime(t time.Time) T {
	return T(t.Unix())
}

// Time converts T into a time.Time.
func (t T) Time() time.Time {
	return time.Unix(int64(t), 0)
}

// MarshalJSON implements the json.Marshaler.
func (t T) MarshalJSON() ([]byte, error) {
	return t.Time().MarshalJSON()
}

// MarshalJSON implements the json.Unmarshaler.
func (t *T) UnmarshalJSON(data []byte) (err error) {
	var rt time.Time
	if err = rt.UnmarshalJSON(data); err != nil {
		return
	}
	*t = T(rt.Unix())
	return
}

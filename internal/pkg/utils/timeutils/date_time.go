package timeutils

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

const (
	dateTimeFormat = "2006-01-02 15:04:05"
)

type DateTime time.Time

func (d *DateTime) Scan(value interface{}) error {
	if value == nil {
		*d = DateTime(time.Time{})
		return nil
	}
	*d = DateTime(value.(time.Time))
	return nil
}

func (d DateTime) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	t, err := time.Parse(dateTimeFormat, dateStr)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(dateTimeFormat))
}

func (d DateTime) ToString() string {
	return time.Time(d).Format(dateTimeFormat)
}

package model

import "time"

type Date struct {
	time.Time
}

// MarshalCSV converts the internal date as CSV string
func (date *Date) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02"), nil
}

// UnmarshalCSV converts CSV string as the internal date
func (date *Date) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02", csv)
	return err
}

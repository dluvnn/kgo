package kgo

import (
	"errors"
	"strconv"
	"time"
)

const (
	BriefDateFormat = "20060102"
	FullDateFormat  = "02-01-2006"
)

var (
	// ErrInvalidDateStringValue ...
	ErrInvalidDateStringValue = errors.New("invalid date string value")
)

// Date ...
type Date int32

// DateNil ...
const DateNil Date = 0

// String converts date to string
func (d *Date) String() string {
	return strconv.Itoa(int(*d))
}

// ParseString converts string to date
func (d *Date) ParseString(s string) error {
	if len(s) == 0 {
		return ErrInvalidDateStringValue
	}
	if !VerifyBriefDate(s) {
		return ErrInvalidDateStringValue
	}
	x, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*d = Date(x)
	return nil
}

// Time convert date to time.Time
func (d *Date) Time() (time.Time, error) {
	return time.Parse(BriefDateFormat, d.String())
}

// Format converts date to string using format 'layout'
func (d *Date) Format(layout string) (string, error) {
	t, err := d.Time()
	if err != nil {
		return "", err
	}
	return t.Format(layout), nil
}

// FullFormat formats date using FullDateFormat layout
func (d *Date) FullFormat() (string, error) {
	return d.Format(FullDateFormat)
}

// Set ...
func (d *Date) Set(year, month, day int) {
	*d = Date(day + month*100 + year*10000)
}

// Add ...
func (d *Date) Add(years, months, days int) error {
	t, err := d.Time()
	if err != nil {
		return err
	}
	year, month, day := t.AddDate(years, months, days).Date()
	d.Set(year, int(month), day)
	return nil
}

// TimeToDate ...
func TimeToDate(t time.Time) Date {
	year, month, day := t.Date()
	return Date(day + (int)(month)*100 + year*10000)
}

// StringToDate ...
func StringToDate(s string) (Date, error) {
	d := DateNil
	err := d.ParseString(s)
	return d, err
}

// VerifyBriefDate ...
func VerifyBriefDate(s string) bool {
	_, err := time.Parse(BriefDateFormat, s)
	return err == nil
}

// ConvertDate ....
func ConvertDate(date, input_fmt, output_fmt string) string {
	t, err := time.Parse(input_fmt, date)
	if err != nil {
		Error(err)
		return ""
	}
	return t.Format(output_fmt)
}

func InterfaceToDate(x interface{}) (Date, bool) {
	i32, ok := x.(int32)
	if !ok {
		return DateNil, false
	}
	return Date(i32), true
}

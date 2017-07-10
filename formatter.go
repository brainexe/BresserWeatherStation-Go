package main

import "fmt"

type Formatter struct {
}

func NewFormatter() Formatter {
	return Formatter{}
}

func (f *Formatter) Format(r *Result) string {
	str := ""

	str += fmt.Sprintf("Temerature %0.1f°\n", r.temperature)
	str += fmt.Sprintf("Humidity %d%%\n", r.humidity)
	str += fmt.Sprintf("Rain %0.1fmm\n", r.rain)
	str += fmt.Sprintf("Wind %d (%d)\n", r.windSpeed, r.windDirection)

	return str
}

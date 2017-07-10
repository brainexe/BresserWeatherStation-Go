package main

type StationId string
type Temperature float32
type Humidity uint8
type Rain float32
type WindDirection uint8
type WindSpeed uint8
type RawData string

type Result struct {
	stationId     StationId
	temperature   Temperature
	humidity      Humidity
	windSpeed     WindSpeed
	windDirection WindDirection
	rain          Rain
	rawData       RawData
}

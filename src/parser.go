package main

import (
	"strconv"
	"fmt"
)

type parser struct {
}

func (p parser) parse (raw string) Result {
	hex := ""
	for i := 0; i < 66*4; i+=4 {
		tmp, _ := strconv.ParseInt(raw[i:i+4], 2, 64)
		hex += fmt.Sprintf("%x", tmp)
	}

	result := Result{}
	if hex[0:10] != "aaaaaaaaaa" {
		return result
	}

	for n := 0; n <= 25; n++ {
		if p.getDigit(hex, n+14) ^ 0xf != p.getDigit(hex, n+40) {
		//	return result
		}
		//	if ord(hex[14+n:15+n]) ^ 0xf != ord(hex[40+n:41+n]):
		//	return 4
		// todo
	}

	result.rawData = RawData(hex)
	result.stationId = p.getStationId(hex)
	result.temperature = p.getTemperature(hex)
	result.humidity = p.getHumidity(hex)
	result.rain = p.getRain(hex)
	result.windSpeed = p.getWindSpeed(hex)
	result.windDirection = p.getWindDirection(hex)

	return result
}

func (p parser) getStationId(hex string) StationId {
	return StationId(hex[10:14])
}

func (p parser) getTemperature(hex string) Temperature {
	temp_digit_2 := p.getDigit(hex, 54)
	temp_decimal := p.getDigit(hex, 55)
	temp_sign := p.getDigit(hex, 65)
	temp_digit_1 := p.getDigit(hex, 57)

	temperature := Temperature(float32(temp_digit_1*10) + float32(temp_digit_2) + (float32(temp_decimal) / float32(10)))
	if temp_sign != 0 {
		temperature = 0 - temperature
	}

	return temperature
}

func (p parser) getHumidity(hex string) Humidity {
	hum_digit_1 := p.getDigit(hex, 58)
	hum_digit_2 := p.getDigit(hex, 59)

	return Humidity(hum_digit_1 * 10 + hum_digit_2)
}

func (p parser) getDigit (hex string, digit int) uint8 {
	number, _ := strconv.ParseUint(hex[digit:digit+1], 16, 10)

	return uint8(number)
}

func (p parser) getRain (hex string) Rain {
	rain_digit_2 := p.getDigit(hex, 60)
	rain_decimal := p.getDigit(hex, 61)
	rain_digit_1 := p.getDigit(hex, 63)

	return Rain(float32(rain_digit_1 * 10 + rain_digit_2) + (float32(rain_decimal)/float32(10)))
}

func (p parser) getWindDirection (hex string) WindDirection {
	return WindDirection(p.getDigit(hex, 48))
}

func (p parser) getWindSpeed (hex string) WindSpeed {
	wind_digit_1 := p.getDigit(hex, 49)
	wind_digit_2 := p.getDigit(hex, 50)
	wind_decimal := p.getDigit(hex, 51)

	return WindSpeed(float32(wind_digit_1 * 10 + wind_digit_2) + (float32(wind_decimal)/float32(10)))
}



package main

import (
	"strings"
	"testing"
)

func TestFormat(t *testing.T) {
	result := &Result{}
	result.temperature = 21.3
	result.humidity = 51

	subject := NewFormatter()
	actual := subject.Format(result)

	if !strings.Contains(actual, "Temerature 21.3°\n") {
		t.Errorf("Missing temperature: %s", actual)
	}

	if !strings.Contains(actual, "Humidity 51%\n") {
		t.Errorf("Missing humidiy: %s", actual)
	}
}

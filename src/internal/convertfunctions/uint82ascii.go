package convertfunctions

import "github.com/brutella/can"

func Uint82AsciiToCan(input string) (can.Frame, error) {
	var returner [8]byte
	var i uint8 = 0
	for ; int(i) < len(input) && i < 8; i++ {
		returner[i] = input[i]
	}
	return can.Frame{Length: i, Data: returner}, nil
}

func Uint82AsciiToMqtt(input can.Frame) (string, error) {
	return string(input.Data[:input.Length]), nil
}

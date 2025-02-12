// Code generated from Pkl module `Control`. DO NOT EDIT.
package running

import (
	"encoding"
	"fmt"
)

type Running string

const (
	ON     Running = "ON"
	OFF    Running = "OFF"
	SHADOW Running = "SHADOW"
)

// String returns the string representation of Running
func (rcv Running) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Running)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Running.
func (rcv *Running) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "ON":
		*rcv = ON
	case "OFF":
		*rcv = OFF
	case "SHADOW":
		*rcv = SHADOW
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Running`, str)
	}
	return nil
}

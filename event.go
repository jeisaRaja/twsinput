package twsinput

import (
	"encoding/binary"
	"os"
)

type InputEvent struct {
	Time  int64
	Type  uint16
	Code  int16
	Value int32
}

func ReadEvent(file *os.File) (InputEvent, error) {
	var event InputEvent
	err := binary.Read(file, binary.LittleEndian, &event)
	return event, err
}

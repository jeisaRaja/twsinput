package twsinput

import (
	"fmt"
	"testing"
)

func TestReadEvent(t *testing.T) {
	file, err := OpenDevice("/dev/input/event30")
	if err != nil {
		t.Fatalf("error when opening file, %v", err)
	}
	ie, err := ReadEvent(file)
	if err != nil {
		t.Fatalf("there is an error: %v", err)
	}
	fmt.Println(ie)
}

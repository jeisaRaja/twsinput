package twsinput

import "fmt"

type TWSDevice struct {
	Path string
}

func (d *TWSDevice) Init() error{
  fmt.Println("Initialization TWS Device on path: ", d.Path)
  return nil
}

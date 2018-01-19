package serial

import (
	dr "../../../model/driver"
	"encoding/binary"
)

type BaseSerialDriver struct {
	BaseSerial

	endian binary.ByteOrder
}

func NewSerialDriver() {

}

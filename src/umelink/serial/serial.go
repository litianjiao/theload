package serial

import ("encoding/binary"
dr "../../../model/driver"
)


type BaseSerialDriver struct {
	BaseSerial

	endian binary.ByteOrder

}

func NewSerialDriver()  {

}
package serial

import (
	//dr "../../model/driver"
	"encoding/binary"
)



type BaseSerialDriver struct {
	BaseSerial
	endian binary.ByteOrder

}

func openSerialPort() (result bool) {
result=true
if
}
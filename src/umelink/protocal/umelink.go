package protocal

import "encoding/binary"

type UmelinkUnit struct {
	bytes []byte
	err string
}


func EmptyUmelinkUnitV1(msg_size int)*UmelinkUnit  {
	var unit =make([]byte,5+msg_size+3)
	binary.BigEndian.PutUint16(unit[0:],0xCCCC)     			//0XCCCC 头 2
	unit[2]=0x10	                                   			//len  1
	binary.BigEndian.PutUint16(unit[3:],0x0001)       			//CMD 2
	//copy(unit[5:],msg)										//*msg
	unit[5+msg_size:]=append(unit[5+msg_size:],0x00)		//XOR    1
	binary.BigEndian.PutUint16(unit[5+msg_size+1:],0xDDDD)	//结束符 2

	return &UmelinkUnit{
		bytes:unit,
		err:"",
	}

}

func (obj *UmelinkUnit)SetCmd(cmd uint16)*UmelinkUnit  {
	binary.BigEndian.PutUint16(obj.bytes[4:],cmd)
	return  obj
}

func (obj *UmelinkUnit) SetMessage(msg []byte)*UmelinkUnit  {
	copy(obj.bytes[5:],msg)
}

func (obj *UmelinkUnit) Size() int {
	return len(obj.bytes)
}

func (obj *UmelinkUnit) HeadTag() uint16 {
	return binary.BigEndian.Uint16(obj.bytes[0:])
}

func (obj *UmelinkUnit)Cmd()uint16  {
	return binary.BigEndian.Uint16(obj.bytes[3:])
}

func (obj *UmelinkUnit)Message()[]byte  {
	return obj.bytes[5:obj.Size()-3]
}

func (obj *UmelinkUnit) CheckXor() byte {
	return obj.bytes[obj.Size()-2]
}

func (obj *UmelinkUnit) EndTag() uint16 {
	return binary.BigEndian.Uint16(obj.bytes[obj.Size() - 2 :])
}

func (obj *UmelinkUnit) ToBytes() []byte {
	return obj.bytes
}


func CheckXor(xs []byte) bool {
	if size := len(xs); size < 3 {
		return false
	} else {
		xor_byte := xs[size - 1]
		xor_sum := xs[1]
		for _,x := range xs[2:size-1] {
			xor_sum ^= byte(x)
		}
		return xor_byte == xor_sum
	}
}

func XorBytes(xs []byte) byte {
	if size := len(xs); size == 0 {
		return 0
	} else {
		var xorSum byte = xs[0]
		for _, x := range xs[1:] {
			xorSum ^= byte(x)
		}
		return xorSum
	}
}




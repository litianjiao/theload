package serial

import (
	"io"
	"time"

	goSerial "github.com/tarm/goserial"
	"errors"
)




/*定义了操作串口的机制*/
type BaseSerial struct {
	fd             io.ReadWriteCloser    //read write or close   if err or needs reset the io,use close
	timeout        time.Duration        //超时时间
	recover_delay  time.Duration        //重试延时
	recover_trycnt uint                 //重试次数
	address        string
	config         *goSerial.Config     //COM_name 波特率
}



func (this *BaseSerial) Timeout(timeout time.Duration) {
	this.timeout = timeout
}
/**/
func (this *BaseSerial) Close() (err error) {
	if this.IsOpen() {
		err = this.fd.Close()
		this.fd = nil
	}
	return
}

func (this *BaseSerial) Interrupt() error {
	return this.Close()
}

func (this *BaseSerial) IsOpen() bool {
	return nil != this.fd
}

func (this *BaseSerial) Address() string {
	return this.address
}
/* open serial*/
func (this *BaseSerial) Open(address string, port uint, timeout time.Duration) bool {
	this.Close()
	config := &goSerial.Config{
		Name:        address,
		Baud:        int(port),
		ReadTimeout: timeout}
	fd, err := goSerial.OpenPort(config)
	if nil != err {
		return false
	}

	this.fd = fd
	this.timeout = timeout
	this.config = config

	return true
}

func (this *BaseSerial) Flush() error {
	return nil
}

func (this *BaseSerial) Read(buf []byte) (n int, err error) {
	if !this.IsOpen() {
		err = errors.New("Serial not open")
		return
	}
	n, err = this.fd.Read(buf)
	return
}

func (this *BaseSerial) Write(buf []byte) (int, error) {
	if !this.IsOpen() {
		return 0, errors.New("Serial not open")
	}
	return this.fd.Write(buf)
}

func (this *BaseSerial) Recover() bool {
	if nil == this.config || this.recover_delay < 1 {
		return false
	}

	var delay time.Duration
	trycnt := this.recover_trycnt

	for {
		fd, err := goSerial.OpenPort(this.config)
		if nil == err {
			this.Close()
			this.fd = fd
			return true
		}
		if trycnt--; trycnt < 0 {
			break
		}

		delay += this.recover_delay
		time.Sleep(delay)
	}
	return false
}

func (this *BaseSerial) NoDelay(bool) {

}

func (this *BaseSerial) RecoverDefault() {
	this.recover_delay = 5 * time.Second
	this.recover_trycnt = 256
}

func (this *BaseSerial) RecoverSetup(delay time.Duration, trycount ...uint) {
	this.recover_delay = delay
	if len(trycount) > 0 {
		this.recover_trycnt = trycount[0]
	}
}

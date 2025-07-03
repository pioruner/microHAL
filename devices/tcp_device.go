package devices

import (
	"net"
	"time"
)

type TCPDevice struct {
	Address string
	Conn    net.Conn
}

func (d *TCPDevice) Connect() error {
	conn, err := net.DialTimeout("tcp", d.Address, 5*time.Second)
	if err != nil {
		return err
	}
	d.Conn = conn
	return nil
}

func (d *TCPDevice) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := d.Conn.Read(buf)
	return buf[:n], err
}

func (d *TCPDevice) Write(data []byte) error {
	_, err := d.Conn.Write(data)
	return err
}

func (d *TCPDevice) Close() error {
	return d.Conn.Close()
}

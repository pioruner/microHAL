package devices

type Device interface {
	Connect() error
	Run() error
	SetUP() error
	Write([]byte) (int, error)
	Read([]byte) (int, error)
	Disconnect() error
}

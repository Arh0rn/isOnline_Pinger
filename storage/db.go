package storage

type DB interface {
	ConnectDB() error
	CloseDB() error
	GetUrls() ([]string, error)
	AddUrl(url string) error
	DeleteUrl(url string) error
	GetParameters() (Parameters, error)
	SetParameters(p Parameters) error
}

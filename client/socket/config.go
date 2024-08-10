package socket

type SocketConfig struct {
	Type string
	Host string
}

var Config = &SocketConfig{
	Type: "tcp",
	Host: "localhost:8080",
}

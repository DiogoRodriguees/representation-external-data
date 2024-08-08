package config

type SocketType struct {
	Type string
	Host string
}

var Socket = &SocketType{
	Type: "tcp",
	Host: "localhost:8080",
}

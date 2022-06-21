package config

type Mongo struct {
	Host        string
	Name        string
	Port        int
	Collections Collections
}

type Collections struct {
	Users string
}

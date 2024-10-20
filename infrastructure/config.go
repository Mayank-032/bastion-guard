package infrastructure

type database struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Schema string `json:"schema"`
}

type Config struct {
	Database database `json:"database"`
}

var Configurations Config

func InitConfig() error {
	return nil
}

package pkg

type ProxyArgs struct {
	ConfigFile string
	ReadJSON   bool
	ReadYAML   bool
}

type ToConfig struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Schema string `json:"schema" yaml:"schema"`
}

type PathConfig struct {
	Path string    `json:"path" yaml:"path"`
	To   *ToConfig `json:"to" yaml:"to"`
}

type ServerConfig struct {
	Name      string        `json:"name" yaml:"name"`
	Locations []*PathConfig `json:"locations" yaml:"locations"`
}

type ServersType struct {
	Servers []*ServerConfig `json:"servers" yaml:"servers"`
}

type readFuncType func(fileData []byte) error

var readFunc readFuncType
var config *ServersType

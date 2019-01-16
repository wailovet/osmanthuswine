package owstruct

type Config struct {
	Port        string `json:"port"`
	Host        string `json:"host"`
	CrossDomain string `json:"cross_domain"`
	MaxMemory   int64  `json:"post_max_memory"`
}

func (c *Config) ReadConfig(file string) {

}

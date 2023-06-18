package config

type Consul struct {
	Address string `mapstructure:"address" json:"address" yaml:"address"` // consul注册地址
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`          // consul注册端口
}

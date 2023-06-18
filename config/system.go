package config

type System struct {
	Name       string   `mapstructure:"name" json:"name" yaml:"name"`                   // consul注册名称
	Port       int      `mapstructure:"port" json:"port" yaml:"port"`                   // consul注册端口
	Tags       []string `mapstructure:"tags" json:"tags" yaml:"tags"`                   // consul注册tag集
	Timeout    string   `mapstructure:"timeout" json:"timeout" yaml:"timeout"`          // consul超时时间
	Interval   string   `mapstructure:"interval" json:"interval" yaml:"interval"`       // consul检查频率
	Deregister string   `mapstructure:"deregister" json:"deregister" yaml:"deregister"` // 检查失败多久删除实例
	ApiKey     string   `mapstructure:"apiKey" json:"apiKey" yaml:"apiKey"`             // 检查失败多久删除实例
}

package config

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                         // 服务器地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                         // 端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                   // 高级配置
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                   // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`             // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`             // 数据库密码
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                   // 全局表前缀，单独定义TableName则不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`             // 是否开启全局禁用复数，true表示开启
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`  // 数据库引擎，默认InnoDB
	MaxIdleConns int    `mapstructure:"maxIdleConns" json:"maxIdleConns" yaml:"maxIdleConns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"logMode" json:"logMode" yaml:"logMode"`                // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"logZap" json:"logZap" yaml:"logZap"`                   // 是否通过zap写入日志文件
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}

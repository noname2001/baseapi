package config

import "time"

type Server struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type App struct {
	RunMode      string        `mapstructure:"run-mode" json:"run-mode" yaml:"run-mode"`
	HTTPPort     int           `mapstructure:"http-port" json:"http-port" yaml:"http-port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout" json:"read-timeout" yaml:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout" json:"write-timeout" yaml:"write-timeout"`
	PageSize     int           `mapstructure:"pagesize" json:"pagesize" yaml:"page-size"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

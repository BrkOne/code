package db

import (
	"github.com/TeaWeb/code/teaconfigs/shared"
	"github.com/go-yaml/yaml"
	"github.com/iwind/TeaGo/Tea"
	"io/ioutil"
)

const (
	mysqlFilename = "mysql.conf"
)

// MySQL配置
type MySQLConfig struct {
	DSN string `yaml:"dsn" json:"dsn"`

	Addr     string `yaml:"addr" json:"addr"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	DBName   string `yaml:"dbName" json:"dbName"`
}

// 获取新MySQL配置对象
func NewMySQLConfig() *MySQLConfig {
	return &MySQLConfig{}
}

// 默认的MySQL配置
func DefaultMySQLConfig() *MySQLConfig {
	return &MySQLConfig{
		Addr:     "127.0.0.1:3306",
		Username: "root",
		DBName:   "teaweb",
	}
}

// 加载MySQL配置
func LoadMySQLConfig() (*MySQLConfig, error) {
	data, err := ioutil.ReadFile(Tea.ConfigFile(mysqlFilename))
	if err != nil {
		return nil, err
	}
	config := &MySQLConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// 组合DSN
func (this *MySQLConfig) ComposeDSN() string {
	return this.Username + ":" + this.Password + "@tcp(" + this.Addr + ")/" + this.DBName
}

// 保存
func (this *MySQLConfig) Save() error {
	shared.Locker.Lock()
	defer shared.Locker.WriteUnlockNotify()

	data, err := yaml.Marshal(this)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(Tea.ConfigFile(mysqlFilename), data, 0777)
}

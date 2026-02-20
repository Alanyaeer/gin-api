package config

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
)

// Config 定义全局配置结构体，与配置文件结构一一对应
// 结构化管理配置，比零散读取更规范、易维护
type Config struct {
	Redis   RedisConfig   `mapstructure:"redis"`   // mapstructure 用于Viper解析yaml到结构体
	MySQL MySQLConfig `mapstructure:"mysql"`
	Agent  AgentConfig   `mapstructure:"agent"`
	App AppConfig `mapstructure:"app"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Dsn     string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Addr string `mapstructure:"addr"`
	DB   int    `mapstructure:"db"`
}

type AgentConfig struct {
	APIKey string `mapstructure:"api_key"`
	Model string `mapstructure:"model"`
	BaseURL string `mapstructure:"base_url"`
}
type AppConfig struct {
	Port int `mapstructure:"port"`
}


var Cfg Config

func InitConfig(configPath string,configName string,configType string) error{
		// 1. 初始化 Viper
	v := viper.New() // 使用独立的Viper实例，避免全局Viper冲突（规范做法）
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	v.AddConfigPath(configPath)
	if err := v.ReadInConfig(); err != nil {
		slog.Error("Failed to read config file", "error", err)
		return err
	}
	if err := v.Unmarshal(&Cfg); err != nil {
		slog.Error("Failed to unmarshal config", "error", err)
		return err
	}
	setDefaults(v)
		// 3. 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return errors.New("配置文件未找到: " + err.Error())
		}
		return errors.New("读取配置文件失败: " + err.Error())
	}

	// 4. 将配置解析到结构体中
	if err := v.Unmarshal(&Cfg); err != nil {
		return errors.New("解析配置到结构体失败: " + err.Error())
	}
	slog.Info(fmt.Sprintf("Loaded config:\n%+v\n", Cfg))
	// 5. 配置校验（可选，规范做法：确保关键配置非空）
	if err := validateConfig(); err != nil {
		return errors.New("配置校验失败: " + err.Error())
	}
	return nil
}

// setDefaults 设置配置默认值
func setDefaults(v *viper.Viper) {
	// MySQL 配置默认值
	v.SetDefault("mysql.dsn", "root:123456@tcp(127.0.0.1:13306)/go_gin?charset=utf8mb4&parseTime=True&loc=Local")
}

// validateConfig 校验关键配置是否合法
func validateConfig() error {
	if Cfg.MySQL.Dsn == "" {
		return errors.New("mysql.dsn 不能为空")
	}
	if Cfg.Redis.Addr == "" {
		return errors.New("redis.addr 不能为空")
	}
	if Cfg.Agent.APIKey == "" {
		return errors.New("agent.api_key 不能为空")
	}
	if Cfg.Agent.Model == "" {
		return errors.New("agent.model 不能为空")
	}
	if Cfg.Agent.BaseURL == "" {
		return errors.New("agent.base_url 不能为空")
	}
	return nil
}
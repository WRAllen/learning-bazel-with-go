package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"port"`
}

var Global *Config

// Load 加载并解析 config.yaml
func Load() (*Config, error) {
	viper.SetConfigName("config")    // 文件名（不带扩展名）
	viper.SetConfigType("yaml")      // 类型
	viper.AddConfigPath("./config/") // 你也可以改成绝对路径或其他位置

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	Global = &cfg
	return &cfg, nil
}

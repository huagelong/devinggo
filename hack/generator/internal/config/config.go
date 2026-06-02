package config

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/os/gfile"
)

// GeneratorConfig generator.yaml 配置
type GeneratorConfig struct {
	Module string           `yaml:"module"`
	Tables []GeneratorTable `yaml:"tables"`
}

// GeneratorTable 单表生成配置
type GeneratorTable struct {
	Table       string `yaml:"table"`
	Business    string `yaml:"business"`
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
	Date        string `yaml:"date"`
}

// LoadGeneratorConfig 加载 generator.yaml
func LoadGeneratorConfig(path string) (*GeneratorConfig, error) {
	if !gfile.Exists(path) {
		return nil, fmt.Errorf("配置文件不存在: %s", path)
	}

	content := gfile.GetContents(path)
	if content == "" {
		return nil, fmt.Errorf("配置文件为空: %s", path)
	}

	config := &GeneratorConfig{}
	if err := gyaml.DecodeTo([]byte(content), config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	if config.Module == "" {
		return nil, fmt.Errorf("配置文件缺少 module 字段")
	}

	if len(config.Tables) == 0 {
		return nil, fmt.Errorf("配置文件缺少 tables 配置")
	}

	return config, nil
}

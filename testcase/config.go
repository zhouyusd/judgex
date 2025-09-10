package testcase

import (
	"github.com/zhouyusd/judgex/checker"
)

type Config struct {
	Checker   *CheckerConfig   `json:"checker"`   // 检查器配置
	Generator *GeneratorConfig `json:"generator"` // 数据生成器配置
	Standard  *StandardConfig  `json:"standard"`  // 标准答案配置
	Testcases []*Testcase      `json:"testcases"` // 测试点列表
}

func (c *Config) Validate() error {
	return nil
}

type CheckerConfig struct {
	Type checker.Type `json:"type"` // 检查器类型

	// Type 为testlib时，以下字段必填
	SrcName string `json:"src_name"` // 源文件名
	ExeName string `json:"exe_name"` // 可执行文件名
	Code    string `json:"code"`     // 检查器代码
}

type GeneratorConfig struct{}

type StandardConfig struct {
}

type Testcase struct {
	InputName        string `json:"input_name"`
	OutputName       string `json:"output_name"`
	OutputSize       int64  `json:"output_size"`
	OutputMd5        string `json:"output_md5"`
	OutputTrimMd5    string `json:"output_trim_md5"` // 去除开头结尾空白字符后的输出文件的MD5
	OutputNoSpaceMd5 string `json:"output_no_space_md5"`
}

package config

type Local struct {
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`          // 文件夹名
	RootPath string `mapstructure:"root_path" json:"root_path" yaml:"root_path"` // 根目录
}

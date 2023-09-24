package config

type WorkerPool struct {
	MaxCapicity int `mapstructure:"maxCapicity" json:"maxCapicity" yaml:"maxCapicity"`
}

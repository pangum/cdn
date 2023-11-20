package config

type D struct {
	// 密钥
	Key string `json:"key" yaml:"key" xml:"key" toml:"key" validate:"required"`
	// 签名字段
	Signature string `default:"sign" json:"signature" yaml:"signature" xml:"signature" toml:"signature"`
	// 时间字段
	Timestamp string `default:"t" json:"timestamp" yaml:"timestamp" xml:"timestamp" toml:"timestamp"`
}

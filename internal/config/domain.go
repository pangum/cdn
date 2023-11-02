package config

type Domain struct {
	// 主机
	Host string `json:"host" yaml:"host" xml:"host" toml:"host" validate:"required,hostname|hostname_port"`
	// 模式
	Scheme string `default:"https" json:"scheme" yaml:"scheme" xml:"scheme" toml:"scheme" validate:"oneof=http https"`
	// 匹配
	Pattern string `json:"pattern" yaml:"pattern" xml:"pattern" toml:"pattern"`
	// 匹配列表
	Patterns []string `json:"patterns" yaml:"patterns" xml:"patterns" toml:"patterns"`
	// 创世云
	// nolint: lll
	Chuangcache *Chuangcache `json:"chuangcache" yaml:"chuangcache" xml:"chuangcache" toml:"chuangcache"  validate:"required_without_all=Tencent"`
	// 腾讯云
	// nolint: lll
	Tencent *Tencent `json:"Tencent" yaml:"Tencent" xml:"Tencent" toml:"Tencent"  validate:"required_without_all=Chuangcache"`
}

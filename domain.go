package cdn

type domain struct {
	// 主机
	Host string `json:"host" yaml:"host" xml:"host" toml:"host"`
	// 模式
	Scheme string `default:"https" json:"scheme" yaml:"scheme" xml:"scheme" toml:"scheme" validate:"oneof=http https"`
	// 匹配
	Pattern string `json:"pattern" yaml:"pattern" xml:"pattern" toml:"pattern"`
	// 匹配列表
	Patterns []string `json:"patterns" yaml:"patterns" xml:"patterns" toml:"patterns"`
	// 创世云
	// nolint: lll
	Chuangcache *chuangcache `json:"chuangcache" yaml:"chuangcache" xml:"chuangcache" toml:"chuangcache"  validate:"required_without_all=Tencent"`
	// 腾讯云
	Tencent *tencent `json:"tencent" yaml:"tencent" xml:"tencent" toml:"tencent"  validate:"required_without_all=Chuangcache"`
}
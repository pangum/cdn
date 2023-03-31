package cdn

type config struct {
	// 域名
	Domain *domain `json:"domain" yaml:"domain" xml:"domain" toml:"domain" validate:"required_without=Domains"`
	// 域名
	Domains []*domain `json:"domains" yaml:"domains" xml:"domains" toml:"domains" validate:"required_without=Domain"`
}

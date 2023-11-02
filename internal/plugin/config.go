package plugin

import (
	"github.com/pangum/cdn/internal/config"
)

type Config struct {
	// 域名
	Domain *config.Domain `json:"domain" yaml:"domain" xml:"domain" toml:"domain" validate:"required_without=Domains"`
	// 域名
	Domains []*config.Domain `json:"domains" yaml:"domains" xml:"domains" toml:"domains" validate:"required_without=Domain"`
}

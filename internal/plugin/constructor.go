package plugin

import (
	"github.com/goexl/cdn"
	"github.com/pangum/pangu"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(config *pangu.Config) (client *cdn.Client, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		client, err = c.new(wrapper.CDN)
	}

	return
}

func (c *Constructor) new(config *Config) (client *cdn.Client, err error) {
	builder := cdn.New()
	if nil != config.Domain {
		config.Domains = append(config.Domains, config.Domain)
	}
	for _, domain := range config.Domains {
		if "" != domain.Pattern {
			domain.Patterns = append(domain.Patterns, domain.Pattern)
		}
		_domain := builder.Domain().Host(domain.Host).Scheme(domain.Scheme).Pattern(domain.Patterns...)
		if nil != domain.Chuangcache {
			_domain = _domain.Chuangcache().C(domain.Chuangcache.C).Build()
		} else if nil != domain.Tencent {
			signer := _domain.Tencent()
			if "" != domain.Tencent.A {
				signer.A(domain.Tencent.A)
			} else if "" != domain.Tencent.B {
				signer.B(domain.Tencent.B)
			} else if "" != domain.Tencent.C {
				signer.C(domain.Tencent.C)
			} else if nil != domain.Tencent.D {
				signer.D(domain.Tencent.D.Key, domain.Tencent.D.Signature, domain.Tencent.D.Timestamp)
			}
			_domain = signer.Build()
		}
		builder = _domain.Build()
	}
	client = builder.Build()

	return
}

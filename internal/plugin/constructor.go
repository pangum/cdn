package plugin

import (
	"github.com/goexl/cdn"
	"github.com/pangum/pangu"
)

type Constructor struct {
	// 构造方法
}

func (c *Constructor) New(config *pangu.Config) (client *cdn.CDN, err error) {
	wrapper := new(Wrapper)
	if ge := config.Build().Get(wrapper); nil != ge {
		err = ge
	} else {
		client, err = c.new(wrapper.CDN)
	}

	return
}

func (c *Constructor) new(config *Config) (client *cdn.CDN, err error) {
	builder := cdn.New()
	if nil != config.Domain {
		config.Domains = append(config.Domains, config.Domain)
	}
	for _, _domain := range config.Domains {
		if "" != _domain.Pattern {
			_domain.Patterns = append(_domain.Patterns, _domain.Pattern)
		}
		if nil != _domain.Chuangcache {
			chuangecache := builder.Chuangcache().Host(_domain.Host).Scheme(_domain.Scheme).Pattern(_domain.Patterns...)
			signer := chuangecache.Signer()
			if "" != _domain.Chuangcache.C {
				signer.A(_domain.Chuangcache.C)
			}
			signer.Build()
			chuangecache.Build()
		} else if nil != _domain.Tencent {
			tencent := builder.Tencent().Host(_domain.Host).Scheme(_domain.Scheme).Pattern(_domain.Patterns...)
			signer := tencent.Signer()
			if "" != _domain.Tencent.A {
				signer.A(_domain.Tencent.A)
			} else if "" != _domain.Tencent.B {
				signer.B(_domain.Tencent.B)
			} else if "" != _domain.Tencent.C {
				signer.C(_domain.Tencent.C)
			} else if nil != _domain.Tencent.D {
				signer.D(_domain.Tencent.D.Key, _domain.Tencent.D.Signature, _domain.Tencent.D.Timestamp)
			}
			signer.Build()
			tencent.Build()
		}
	}
	client = builder.Build()

	return
}

package cdn

import (
	"github.com/goexl/cdn"
	"github.com/pangum/pangu"
)

func newCDN(config *pangu.Config) (cd *CDN, err error) {
	wrap := new(wrapper)
	if err = config.Load(wrap); nil != err {
		return
	}

	builder := cdn.New()
	conf := wrap.CDN
	if nil != conf.Domain {
		conf.Domains = append(conf.Domains, conf.Domain)
	}
	for _, _domain := range conf.Domains {
		if "" != _domain.Pattern {
			_domain.Patterns = append(_domain.Patterns, _domain.Pattern)
		}
		if nil != _domain.Chuangcache {
			_chuangecache := builder.Chuangcache().Host(_domain.Host).Scheme(_domain.Scheme).Pattern(_domain.Patterns...)
			signer := _chuangecache.Signer()
			if "" != _domain.Chuangcache.A {
				signer.A(_domain.Chuangcache.A)
			}
			signer.Build()
			_chuangecache.Build()
		} else if nil != _domain.Tencent {
			_tencent := builder.Tencent().Host(_domain.Host).Scheme(_domain.Scheme).Pattern(_domain.Patterns...)
			signer := _tencent.Signer()
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
			_tencent.Build()
		}
	}

	return
}

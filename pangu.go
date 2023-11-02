package cdn

import (
	"github.com/pangum/cdn/internal/plugin"
	"github.com/pangum/pangu"
)

func init() {
	ctor := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		ctor.New,
	).Build().Build().Apply()
}

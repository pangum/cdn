package cdn

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newCDN,
	)
}

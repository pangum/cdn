package cdn

type wrapper struct {
	CDN *config `json:"cdn" yaml:"cdn" xml:"cdn" toml:"cdn" validate:"required"`
}

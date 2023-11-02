package plugin

type Wrapper struct {
	CDN *Config `json:"cdn" yaml:"cdn" xml:"cdn" toml:"cdn" validate:"required"`
}

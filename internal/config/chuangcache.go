package config

type Chuangcache struct {
	C string `json:"c" yaml:"c" xml:"c" toml:"c" validate:"required,min=6,max=40"`
}

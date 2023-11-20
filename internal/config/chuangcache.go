package config

type Chuangcache struct {
	A string `json:"a" yaml:"a" xml:"a" toml:"a" validate:"required,min=6,max=40"`
}

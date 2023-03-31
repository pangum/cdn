package cdn

type tencent struct {
	A string `json:"a" yaml:"a" xml:"a" toml:"a" validate:"required_without_all=B C D"`
	B string `json:"b" yaml:"b" xml:"b" toml:"b" validate:"required_without_all=A C D"`
	C string `json:"c" yaml:"c" xml:"c" toml:"c" validate:"required_without_all=A B D"`
	D *d     `json:"d" yaml:"d" xml:"d" toml:"d" validate:"required_without_all=A B C"`
}

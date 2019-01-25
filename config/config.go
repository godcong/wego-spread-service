package config

type Database struct {
}

type Configure struct {
	Database Database `toml:"database"`
}

func InitLoader(path string) *Configure {
	return &Configure{}
}

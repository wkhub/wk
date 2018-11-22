package mixer

type MixConfig struct {
	Copy   []string
	Ignore []string
}

type Config struct {
	Mix    MixConfig
	Params Parameters
}

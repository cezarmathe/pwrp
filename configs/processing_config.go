package configs

// ProcessingConfig is a configuration data structure for project processing
type ProcessingConfig struct {
	Name        string                 `toml:"name, omitempty"`
	Owner       string                 `toml:"owner, omitempty"`
	OwnerEmail  string                 `toml:"owner_email, omitempty"`
	URL         string                 `toml:"url, omitempty"`
	FrontMatter map[string]FrontMatter `toml:"front_matter, omitempty"`
	Text        string                 `toml:"text, omitempty"`
	Tag         string                 `toml:"tag, omitempty"`
}

// FrontMatter is a data structure that contains the value for the front matter,
// along with the name of a generator shell script
type FrontMatter struct {
	Value     string `toml:"value, omitempty"`
	Generator string `toml:"generator, omitempty"`
}

package articles

type (
	// Model for articles.
	Model struct {
		Articles []*Article `yaml:"books,omitempty"`
	}

	// Article for our site.
	Article struct {
		Title       string `yaml:"title,omitempty"`
		Description string `yaml:"description,omitempty"`
		Slug        string `yaml:"slug,omitempty"`
	}
)

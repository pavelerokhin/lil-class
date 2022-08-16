package utl

type ArticlePreviewSelectors struct {
	ArticleHyperlink string
	Container        string
	Date             string
	Subtitle         string
	Tag              string
	Title            string

	// for controlling the output
	ArticleHyperlinkQuery string `yaml:"-"`
	ContainerQuery        string `yaml:"-"`
	DateQuery             string `yaml:"-"`
	LinkQuery             string `yaml:"-"`
	SubtitleQuery         string `yaml:"-"`
	TagQuery              string `yaml:"-"`
	TitleQuery            string `yaml:"-"`
}

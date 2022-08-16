package utl

type HTMLPreviewSelectors struct {
	ArticleURL string
	Container  string
	Link       string
	Subtitle   string
	Tag        string
	Title      string

	// for controlling the output
	ArticleURLQuery string `yaml:"-"`
	ContainerQuery  string `yaml:"-"`
	LinkQuery       string `yaml:"-"`
	SubtitleQuery   string `yaml:"-"`
	TagQuery        string `yaml:"-"`
	TitleQuery      string `yaml:"-"`
}

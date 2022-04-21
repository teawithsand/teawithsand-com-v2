package trans

type Middleware = func(lang, key, val string) (string, string, string)

type Builder struct {
	middlewares  []Middleware
	translations map[string]map[string]string
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) AddMiddleware(mw Middleware) *Builder {
	b.middlewares = append(b.middlewares, mw)
	return b
}

func (b *Builder) AddTranslation(lang, key, value string) *Builder {
	if b.translations == nil {
		b.translations = make(map[string]map[string]string)
	}
	_, ok := b.translations[lang]
	if !ok {
		b.translations[lang] = make(map[string]string)
	}

	for _, mw := range b.middlewares {
		lang, key, value = mw(lang, key, value)
	}

	m := b.translations[lang]
	m[key] = value

	return b
}

func (b *Builder) Lang(lang string) *LanguageBuilder {
	if b.translations == nil {
		b.translations = make(map[string]map[string]string)
	}

	_, ok := b.translations[lang]
	if !ok {
		b.translations[lang] = make(map[string]string)
	}
	return &LanguageBuilder{
		Builder: b,
		Lang:    lang,
	}
}

func (b *Builder) MustBuild() Translations {
	return Translations(b.translations)
}

package trans

import "strings"

type LanguageBuilder struct {
	Builder *Builder
	Lang    string

	PrefixStack []string
}

func (lb *LanguageBuilder) Put(key string, value string) *LanguageBuilder {
	keyparts := make([]string, 0, len(lb.PrefixStack)+1)
	keyparts = append(keyparts, lb.PrefixStack...)
	keyparts = append(keyparts, key)

	realKey := strings.Join(keyparts, ".")
	lb.Builder.AddTranslation(lb.Lang, realKey, value)
	return lb
}

func (lb *LanguageBuilder) PushPrefix(prefix string) *LanguageBuilder {
	lbc := *lb
	lbc.PrefixStack = append(lb.PrefixStack, prefix)
	return &lbc
}

func (lb *LanguageBuilder) PopPushPrefix(prefix string) *LanguageBuilder {
	return lb.PopPrefix().PushPrefix(prefix)
}

func (lb *LanguageBuilder) WithPrefix(prefix string, cb func(lb *LanguageBuilder) *LanguageBuilder) *LanguageBuilder {
	lb = lb.PushPrefix(prefix)
	lb = cb(lb)
	lb = lb.PopPrefix()
	return lb
}

func (lb *LanguageBuilder) PopPrefix() *LanguageBuilder {
	lbc := *lb
	lbc.PrefixStack = lb.PrefixStack[:len(lb.PrefixStack)-1]
	return &lbc
}

func (lb *LanguageBuilder) ClearPrefix() *LanguageBuilder {
	lbc := *lb
	lbc.PrefixStack = nil
	return &lbc
}

func (lb *LanguageBuilder) Done() *Builder {
	return lb.Builder
}

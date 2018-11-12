package data

type LanguageSet struct {
	Language     string            `json:"language,omitempty"`
	Translations map[string]string `json:"translations,omitempty"`
}

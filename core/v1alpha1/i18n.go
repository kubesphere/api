package v1alpha1

type LanguageCode string
type LocaleString string
type Locales map[LanguageCode]LocaleString

const DefaultLanguageCode = "default"

# Vertimas
A simple Golang translation handling library. "Vertimas" is a Lithuanian word meaning translation.

## Usage
This library assumes that translations will be stored in .json files where filename is canonical language name and file will contain key/value pairs.

Example content of translations/web/en.json:
```json
{
    "greating": "Hello",
    "farewell": "Bye"
}
```

Example content of translations/web/lt.json:
```json
{
    "greating": "Labas",
    "farewell": "Viso gero"
}
```

```go
import (
    "github.com/v3nom/vertimas"
	"golang.org/x/text/language"
)

// Creates instance, preloads translations from "translations/web/" folder and sets initial language to first language in the array (English)
instance, err := vertimas.CreateInstance("translations/web/", []language.Tag{
    language.English,
    language.Lithuanian,
})

// Get translation for key
instance.GetTranslation("greating") // returns "Hello"

// Get all translations for current language
instance.GetTranslations() // returns map of translations for current language

// Tries to find given language in supported language list, sets it as current language and returns best match
instance.SetLanguage("lt") // returns matched language

// Tries to find language from HTTP Accept-Language header in supported language list, sets it as current language and returns best match
instance.SetLanguageFromRequest(r) // returns matched language

// Gets current language
instance.GetLanguage() // returns current language
```

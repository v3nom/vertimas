# Vertimas [![Build Status](https://travis-ci.com/v3nom/vertimas.svg?branch=master)](https://travis-ci.com/v3nom/vertimas)
A simple Golang translation handling library. "Vertimas" is a Lithuanian word meaning translation.

## Usage
This library assumes that translations will be stored in .json files where filename is canonical language name and file will contain key/value pairs.

Example content of translations/web/en.json:
```json
{
    "greating": "Hello",
    "farewell": "Bye",
    "greeting_name": "Hello, {name}"
}
```

Example content of translations/web/lt.json:
```json
{
    "greating": "Labas",
    "farewell": "Viso gero",
    "greeting_name": "Labas, {name}"
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

// Get translation by key and replacing placeholders with dynamic values
instance.GetParametrizedTranslation("greeting_name", map[string]string{
	"name": "Tomas",
})

// Get all translations for current language
instance.GetTranslations() // returns map of translations for current language

// Tries to find given language in supported language list and creates new instance without reloading translation files
lithuanianVertimas := instance.SetLanguage("lt") // returns new instance where current language is updated

// Tries to find language from HTTP Accept-Language header in supported language list and creates new instance without reloading translation files
newInstance := instance.SetLanguageFromRequest(r) // returns new instance with current language set from HTTP request

// Gets current language
instance.GetLanguage() // returns current language
```

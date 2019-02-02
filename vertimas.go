package vertimas

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/text/language"
)

// Instance vertimas instance
type Instance struct {
	languageMap        *map[string]map[string]string
	languageMatcher    *language.Matcher
	supportedLanguages *[]language.Tag
	currentLanguage    language.Tag
}

// SetLanguage sets current language
func (v *Instance) SetLanguage(lang string) *Instance {
	_, i := language.MatchStrings(*v.languageMatcher, lang)
	return v.updateLanguage((*v.supportedLanguages)[i])
}

// SetLanguageFromRequest sets current language from http.Request
func (v *Instance) SetLanguageFromRequest(r *http.Request) *Instance {
	acceptedLanguages := r.Header.Get("Accept-Language")
	_, i := language.MatchStrings(*v.languageMatcher, acceptedLanguages)
	return v.updateLanguage((*v.supportedLanguages)[i])
}

// GetLanguage gets current language
func (v Instance) GetLanguage() language.Tag {
	return v.currentLanguage
}

// GetTranslation gets translation from current language
func (v Instance) GetTranslation(name string) string {
	return (*v.languageMap)[v.currentLanguage.String()][name]
}

var parametrizedRegex = regexp.MustCompile(`\{([a-z0-9]*)\}`)

// GetParametrizedTranslation gets translation from current language with {token} replaced from values map
func (v Instance) GetParametrizedTranslation(name string, values map[string]string) string {
	translation := v.GetTranslation(name)
	tokens := parametrizedRegex.FindAllStringSubmatch(translation, -1)
	for _, v := range tokens {
		translation = strings.Replace(translation, v[0], values[v[1]], 1)
	}
	return translation
}

// GetTranslations gets translation map for current language
func (v Instance) GetTranslations() map[string]string {
	return (*v.languageMap)[v.currentLanguage.String()]
}

func (v Instance) updateLanguage(lang language.Tag) *Instance {
	return &Instance{
		languageMap:        v.languageMap,
		supportedLanguages: v.supportedLanguages,
		languageMatcher:    v.languageMatcher,
		currentLanguage:    lang,
	}
}

// CreateInstance creates vertimas instance and loads .json translation files from given path
func CreateInstance(path string, supported []language.Tag) (*Instance, error) {
	languageMap := map[string]map[string]string{}
	matcher := language.NewMatcher(supported)

	instance := &Instance{
		languageMap:        &languageMap,
		supportedLanguages: &supported,
		languageMatcher:    &matcher,
		currentLanguage:    supported[0],
	}
	for _, v := range supported {
		b, err := ioutil.ReadFile(path + v.String() + ".json")
		if err != nil {
			return nil, err
		}

		data := make(map[string]string, 0)
		err = json.Unmarshal(b, &data)
		if err != nil {
			return nil, err
		}

		languageMap[v.String()] = data
	}
	return instance, nil
}

// CreateInstanceWithPanic creates vertimas instance with panic if file loading fails
func CreateInstanceWithPanic(path string, supported []language.Tag) *Instance {
	i, err := CreateInstance(path, supported)
	if err != nil {
		panic(err)
	}
	return i
}

package vertimas

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/language"
)

// Instance vertimas instance
type Instance struct {
	languageMap        map[string]map[string]string
	languageMatcher    language.Matcher
	supportedLanguages []language.Tag
	currentLanguage    language.Tag
}

// SetLanguage sets current language
func (v *Instance) SetLanguage(lang string) language.Tag {
	_, i := language.MatchStrings(v.languageMatcher, lang)
	v.currentLanguage = v.supportedLanguages[i]
	return v.currentLanguage
}

// SetLanguageFromRequest sets current language from http.Request
func (v *Instance) SetLanguageFromRequest(r *http.Request) language.Tag {
	acceptedLanguages := r.Header.Get("Accept-Language")
	_, i := language.MatchStrings(v.languageMatcher, acceptedLanguages, "en")
	v.currentLanguage = v.supportedLanguages[i]
	return v.currentLanguage
}

// GetLanguage gets current language
func (v Instance) GetLanguage() language.Tag {
	return v.currentLanguage
}

// GetTranslation gets translation from current language
func (v Instance) GetTranslation(name string) string {
	return v.languageMap[v.currentLanguage.String()][name]
}

// GetTranslations gets translation map for current language
func (v Instance) GetTranslations() map[string]string {
	return v.languageMap[v.currentLanguage.String()]
}

// CreateInstance creates vertimas instance and loads .json translation files from given path
func CreateInstance(path string, supported []language.Tag) (*Instance, error) {
	instance := &Instance{
		languageMap:        make(map[string]map[string]string),
		supportedLanguages: supported,
		languageMatcher:    language.NewMatcher(supported),
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

		instance.languageMap[v.String()] = data
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

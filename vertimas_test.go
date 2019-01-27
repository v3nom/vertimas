package vertimas

import (
	"net/http"
	"testing"

	"golang.org/x/text/language"
)

func TestInstanceCreation(t *testing.T) {
	instance := createTestInstance()

	if instance.GetLanguage() != language.English {
		t.Error("Should initialize current language to the first one in the array")
	}

	greeting := instance.GetTranslation("greeting")
	if greeting != "Hello" {
		t.Error("Expected English greeting")
	}
}

func TestCreateWithPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	CreateInstanceWithPanic("/testdata2", []language.Tag{
		language.English,
		language.Lithuanian,
	})
}

func TestCreateWithError(t *testing.T) {
	_, err := CreateInstance("/testdata2", []language.Tag{
		language.English,
		language.Lithuanian,
	})
	if err == nil {
		t.Error("Expected file read error")
	}
}

func TestCreateWithJSONError(t *testing.T) {
	_, err := CreateInstance("testdata/", []language.Tag{
		language.Danish,
	})
	if err == nil {
		t.Error("Expected JSON error")
	}
}

func TestSetLanguage(t *testing.T) {
	instance := createTestInstance()
	lithuanianVertimas := instance.SetLanguage("lt")

	if instance.GetLanguage() != language.English {
		t.Error("Original instance should not change")
	}

	if lithuanianVertimas.GetLanguage() != language.Lithuanian {
		t.Error("New instance should have language set to Lithuanian")
	}

	greeting := instance.GetTranslation("greeting")
	if greeting != "Hello" {
		t.Error("Expected English greeting")
	}

	greeting = lithuanianVertimas.GetTranslation("greeting")
	if greeting != "Labas" {
		t.Error("Expected Lithuanian greeting")
	}
}

func TestSetLanguageFromRequest(t *testing.T) {
	instance := createTestInstance()
	r := &http.Request{}
	r.Header = http.Header{}
	r.Header.Add("Accept-Language", "en;q=0.8,lt;q=0.9")

	lithuanianVertimas := instance.SetLanguageFromRequest(r)

	if lithuanianVertimas.GetLanguage() != language.Lithuanian {
		t.Error("New instance should have language set to Lithuanian")
	}

	greeting := lithuanianVertimas.GetTranslation("greeting")
	if greeting != "Labas" {
		t.Error("Expected Lithuanian greeting")
	}
}

func TestSetDefaultLanguage(t *testing.T) {
	instance := createTestInstance()
	defaultVertimas := instance.SetLanguage("da")

	if defaultVertimas.GetLanguage() != language.English {
		t.Error("Original instance should not change")
	}
}

func TestGetTranslations(t *testing.T) {
	instance := createTestInstance()
	translation := instance.GetTranslations()

	greeting := translation["greeting"]
	farewell := translation["farewell"]

	if greeting != "Hello" {
		t.Error("Expected English greeting")
	}
	if farewell != "Bye" {
		t.Error("Expected English greeting")
	}
}

func createTestInstance() *Instance {
	return CreateInstanceWithPanic("testdata/", []language.Tag{
		language.English,
		language.Lithuanian,
	})
}

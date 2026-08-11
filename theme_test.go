package authsignal

import (
	"encoding/json"
	"testing"
)

func TestFontWeightMarshalsAsAString(t *testing.T) {
	face := FontFace{Url: "https://cdn.example.com/regular.woff2", Weight: "400"}

	jsonBody, err := json.Marshal(face)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"url\":\"https://cdn.example.com/regular.woff2\",\"weight\":\"400\"}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestFontWeightIsOmittedWhenUnset(t *testing.T) {
	face := FontFace{Url: "https://cdn.example.com/regular.woff2"}

	jsonBody, err := json.Marshal(face)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"url\":\"https://cdn.example.com/regular.woff2\"}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// The API coerces a single weight to a number and leaves a range as a string, so a read has to take both.
func TestFontWeightUnmarshalsFromANumberOrAString(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		expected FontWeight
	}{
		{name: "number", body: "{\"url\":\"https://cdn.example.com/a.woff2\",\"weight\":400}", expected: "400"},
		{name: "range", body: "{\"url\":\"https://cdn.example.com/a.woff2\",\"weight\":\"100 900\"}", expected: "100 900"},
		{name: "absent", body: "{\"url\":\"https://cdn.example.com/a.woff2\"}", expected: ""},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var face FontFaceResponse

			if err := json.Unmarshal([]byte(testCase.body), &face); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			if face.Weight != testCase.expected {
				t.Fatalf("bad weight. expected: %v. got : %v", testCase.expected, face.Weight)
			}
		})
	}
}

func TestFontWeightRejectsAValueThatIsNeitherStringNorNumber(t *testing.T) {
	var face FontFaceResponse

	if err := json.Unmarshal([]byte("{\"url\":\"https://cdn.example.com/a.woff2\",\"weight\":true}"), &face); err == nil {
		t.Fatalf("expected an error for a boolean weight")
	}
}

func TestThemeTypographyMarshalsBothRoles(t *testing.T) {
	theme := Theme{
		Typography: SetValue(Typography{
			Text: SetValue(Typeface{
				Faces: SetValue([]FontFace{
					{Url: "https://cdn.example.com/regular.woff2", Weight: "400"},
					{Url: "https://cdn.example.com/variable.woff2", Weight: "100 900"},
				}),
			}),
			Display: SetValue(Typeface{FontUrl: SetValue("https://cdn.example.com/display.woff2")}),
		}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"typography\":{\"text\":{\"faces\":[{\"url\":\"https://cdn.example.com/regular.woff2\",\"weight\":\"400\"},{\"url\":\"https://cdn.example.com/variable.woff2\",\"weight\":\"100 900\"}]},\"display\":{\"fontUrl\":\"https://cdn.example.com/display.woff2\"}}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestSwitchesMarshalWhenSetToFalse(t *testing.T) {
	theme := Theme{
		Links:   SetValue(Links{Underline: SetValue(false)}),
		Shadows: SetValue(Shadows{Enabled: SetValue(false)}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"links\":{\"underline\":false},\"shadows\":{\"enabled\":false}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestSwitchesCanBeClearedWithNull(t *testing.T) {
	theme := Theme{
		Links:   SetValue(Links{Underline: SetNull(false)}),
		Shadows: SetValue(Shadows{Enabled: SetNull(false)}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"links\":{\"underline\":null},\"shadows\":{\"enabled\":null}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// The API omits a switch the tenant never set, which has to read back as unset rather than as off.
func TestSwitchesReadAbsentApartFromFalse(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		expected *bool
	}{
		{name: "absent", body: "{}", expected: nil},
		{name: "empty switch", body: "{\"links\":{},\"shadows\":{}}", expected: nil},
		{name: "off", body: "{\"links\":{\"underline\":false},\"shadows\":{\"enabled\":false}}", expected: boolPointer(false)},
		{name: "on", body: "{\"links\":{\"underline\":true},\"shadows\":{\"enabled\":true}}", expected: boolPointer(true)},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var theme ThemeResponse

			if err := json.Unmarshal([]byte(testCase.body), &theme); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			assertBoolPointer(t, "underline", theme.Links.Underline, testCase.expected)
			assertBoolPointer(t, "shadows enabled", theme.Shadows.Enabled, testCase.expected)
		})
	}
}

func boolPointer(value bool) *bool {
	return &value
}

func assertBoolPointer(t *testing.T, name string, actual *bool, expected *bool) {
	t.Helper()

	if expected == nil {
		if actual != nil {
			t.Fatalf("bad %v. expected: unset. got : %v", name, *actual)
		}
		return
	}

	if actual == nil {
		t.Fatalf("bad %v. expected: %v. got : unset", name, *expected)
	}

	if *actual != *expected {
		t.Fatalf("bad %v. expected: %v. got : %v", name, *expected, *actual)
	}
}

// The API rejects an exitPosition under darkMode, so only the theme container can carry one.
func TestExitPositionMarshalsOnTheThemeContainerOnly(t *testing.T) {
	theme := Theme{
		Container: SetValue(Container{Padding: SetValue(int64(8)), ExitPosition: SetValue("bottom")}),
		DarkMode:  SetValue(DarkMode{Container: SetValue(ModeContainer{Padding: SetValue(int64(8))})}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"container\":{\"padding\":8,\"exitPosition\":\"bottom\"},\"darkMode\":{\"container\":{\"padding\":8}}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// A null clears the stored value, so an unset exit position has to be absent rather than null.
func TestExitPositionIsOmittedWhenUnset(t *testing.T) {
	theme := Theme{Container: SetValue(Container{Padding: SetValue(int64(8))})}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"container\":{\"padding\":8}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

func TestExitPositionReadsFromTheThemeContainer(t *testing.T) {
	var theme ThemeResponse

	if err := json.Unmarshal([]byte("{\"container\":{\"exitPosition\":\"bottom\"}}"), &theme); err != nil {
		t.Fatalf("failed to unmarshal json: %v", err)
	}

	if theme.Container.ExitPosition != "bottom" {
		t.Fatalf("bad exit position. expected: bottom. got : %v", theme.Container.ExitPosition)
	}
}

func TestFacesCanBeClearedWithNull(t *testing.T) {
	typeface := Typeface{Faces: SetNull([]FontFace{})}

	jsonBody, err := json.Marshal(typeface)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"faces\":null}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

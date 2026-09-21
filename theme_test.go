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

func TestThemeTypographyMarshalsEveryRole(t *testing.T) {
	theme := Theme{
		Typography: SetValue(Typography{
			Text: SetValue(Typeface{
				Faces: SetValue([]FontFace{
					{Url: "https://cdn.example.com/regular.woff2", Weight: "400"},
					{Url: "https://cdn.example.com/variable.woff2", Weight: "100 900"},
				}),
			}),
			Display: SetValue(Typeface{FontUrl: SetValue("https://cdn.example.com/display.woff2")}),
			Button:  SetValue(Typeface{Faces: SetValue([]FontFace{{Url: "https://cdn.example.com/button.woff2", Weight: "500"}})}),
		}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"typography\":{\"text\":{\"faces\":[{\"url\":\"https://cdn.example.com/regular.woff2\",\"weight\":\"400\"},{\"url\":\"https://cdn.example.com/variable.woff2\",\"weight\":\"100 900\"}]},\"display\":{\"fontUrl\":\"https://cdn.example.com/display.woff2\"},\"button\":{\"faces\":[{\"url\":\"https://cdn.example.com/button.woff2\",\"weight\":\"500\"}]}}}"

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
		{name: "off", body: "{\"links\":{\"underline\":false},\"shadows\":{\"enabled\":false}}", expected: pointer(false)},
		{name: "on", body: "{\"links\":{\"underline\":true},\"shadows\":{\"enabled\":true}}", expected: pointer(true)},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var theme ThemeResponse

			if err := json.Unmarshal([]byte(testCase.body), &theme); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			assertPointer(t, "underline", theme.Links.Underline, testCase.expected)
			assertPointer(t, "shadows enabled", theme.Shadows.Enabled, testCase.expected)
		})
	}
}

func pointer[T any](value T) *T {
	return &value
}

func assertPointer[T comparable](t *testing.T, name string, actual *T, expected *T) {
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

func TestShadowsMarshalOnBothColourModes(t *testing.T) {
	theme := Theme{
		Shadows:  SetValue(Shadows{Enabled: SetValue(true)}),
		DarkMode: SetValue(DarkMode{Shadows: SetValue(Shadows{Enabled: SetValue(false)})}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"shadows\":{\"enabled\":true},\"darkMode\":{\"shadows\":{\"enabled\":false}}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// A null clears the stored value, so an unset dark mode switch has to be absent rather than null.
func TestDarkModeShadowsAreOmittedWhenUnset(t *testing.T) {
	theme := Theme{DarkMode: SetValue(DarkMode{PrimaryColor: SetValue("#111111")})}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"darkMode\":{\"primaryColor\":\"#111111\"}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// The dark mode switch reads back like the theme one: absent stays unset, false reads as off.
func TestDarkModeShadowsReadAbsentApartFromFalse(t *testing.T) {
	testCases := []struct {
		name     string
		body     string
		expected *bool
	}{
		{name: "absent", body: "{\"darkMode\":{}}", expected: nil},
		{name: "empty switch", body: "{\"darkMode\":{\"shadows\":{}}}", expected: nil},
		{name: "off", body: "{\"darkMode\":{\"shadows\":{\"enabled\":false}}}", expected: pointer(false)},
		{name: "on", body: "{\"darkMode\":{\"shadows\":{\"enabled\":true}}}", expected: pointer(true)},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var theme ThemeResponse

			if err := json.Unmarshal([]byte(testCase.body), &theme); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			assertPointer(t, "dark mode shadows enabled", theme.DarkMode.Shadows.Enabled, testCase.expected)
		})
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

func TestAxisPaddingMarshalsAlongsidePadding(t *testing.T) {
	theme := Theme{
		Container: SetValue(Container{
			Padding:           SetValue(int64(8)),
			PaddingHorizontal: SetValue(int64(24)),
			PaddingVertical:   SetValue(int64(16)),
		}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"container\":{\"padding\":8,\"paddingHorizontal\":24,\"paddingVertical\":16}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// Zero is a padding a tenant can set, so it has to reach the wire rather than read as unset.
func TestAxisPaddingMarshalsWhenSetToZero(t *testing.T) {
	theme := Theme{
		Container: SetValue(Container{
			PaddingHorizontal: SetValue(int64(0)),
			PaddingVertical:   SetValue(int64(0)),
		}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"container\":{\"paddingHorizontal\":0,\"paddingVertical\":0}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// A null clears the axis override and falls the container back to the shared padding.
func TestAxisPaddingCanBeClearedWithNull(t *testing.T) {
	theme := Theme{
		Container: SetValue(Container{
			PaddingHorizontal: SetNull(int64(0)),
			PaddingVertical:   SetNull(int64(0)),
		}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"container\":{\"paddingHorizontal\":null,\"paddingVertical\":null}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

// A null clears the stored value, so an unset axis has to be absent rather than null.
func TestAxisPaddingIsOmittedWhenUnset(t *testing.T) {
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

// The API omits an axis the tenant never set, which has to read back as unset rather than as zero.
func TestAxisPaddingReadsAbsentApartFromZero(t *testing.T) {
	testCases := []struct {
		name       string
		body       string
		horizontal *int64
		vertical   *int64
	}{
		{name: "absent", body: "{}", horizontal: nil, vertical: nil},
		{name: "empty container", body: "{\"container\":{}}", horizontal: nil, vertical: nil},
		{name: "padding only", body: "{\"container\":{\"padding\":8}}", horizontal: nil, vertical: nil},
		{name: "zero", body: "{\"container\":{\"paddingHorizontal\":0,\"paddingVertical\":0}}", horizontal: pointer(int64(0)), vertical: pointer(int64(0))},
		{name: "set", body: "{\"container\":{\"paddingHorizontal\":24,\"paddingVertical\":16}}", horizontal: pointer(int64(24)), vertical: pointer(int64(16))},
		{name: "one axis only", body: "{\"container\":{\"paddingHorizontal\":24}}", horizontal: pointer(int64(24)), vertical: nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var theme ThemeResponse

			if err := json.Unmarshal([]byte(testCase.body), &theme); err != nil {
				t.Fatalf("failed to unmarshal json: %v", err)
			}

			assertPointer(t, "padding horizontal", theme.Container.PaddingHorizontal, testCase.horizontal)
			assertPointer(t, "padding vertical", theme.Container.PaddingVertical, testCase.vertical)
		})
	}
}

// A per-axis override is theme-wide, so the dark mode container carries the shared padding only.
func TestModeContainerDoesNotMarshalAxisPadding(t *testing.T) {
	theme := Theme{
		Container: SetValue(Container{PaddingHorizontal: SetValue(int64(24))}),
		DarkMode:  SetValue(DarkMode{Container: SetValue(ModeContainer{Padding: SetValue(int64(8))})}),
	}

	jsonBody, err := json.Marshal(theme)
	if err != nil {
		t.Fatalf("failed to marshal json")
	}

	expectedJson := "{\"container\":{\"paddingHorizontal\":24},\"darkMode\":{\"container\":{\"padding\":8}}}"

	if string(jsonBody) != expectedJson {
		t.Fatalf("bad json. expected: %v. got : %v", expectedJson, string(jsonBody))
	}
}

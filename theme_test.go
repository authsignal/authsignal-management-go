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

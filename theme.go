package authsignal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Borders struct {
	ButtonBorderRadius    NullableJsonInput[int64] `json:"buttonBorderRadius,omitempty"`
	ButtonBorderWidth     NullableJsonInput[int64] `json:"buttonBorderWidth,omitempty"`
	CardBorderRadius      NullableJsonInput[int64] `json:"cardBorderRadius,omitempty"`
	CardBorderWidth       NullableJsonInput[int64] `json:"cardBorderWidth,omitempty"`
	InputBorderRadius     NullableJsonInput[int64] `json:"inputBorderRadius,omitempty"`
	InputBorderWidth      NullableJsonInput[int64] `json:"inputBorderWidth,omitempty"`
	ContainerBorderRadius NullableJsonInput[int64] `json:"containerBorderRadius,omitempty"`
}

type Colors struct {
	ButtonPrimaryText         NullableJsonInput[string] `json:"buttonPrimaryText,omitempty"`
	ButtonPrimaryBorder       NullableJsonInput[string] `json:"buttonPrimaryBorder,omitempty"`
	ButtonSecondaryText       NullableJsonInput[string] `json:"buttonSecondaryText,omitempty"`
	ButtonSecondaryBackground NullableJsonInput[string] `json:"buttonSecondaryBackground,omitempty"`
	ButtonSecondaryBorder     NullableJsonInput[string] `json:"buttonSecondaryBorder,omitempty"`
	CardBackground            NullableJsonInput[string] `json:"cardBackground,omitempty"`
	CardBorder                NullableJsonInput[string] `json:"cardBorder,omitempty"`
	InputBackground           NullableJsonInput[string] `json:"inputBackground,omitempty"`
	InputBorder               NullableJsonInput[string] `json:"inputBorder,omitempty"`
	Link                      NullableJsonInput[string] `json:"link,omitempty"`
	HeadingText               NullableJsonInput[string] `json:"headingText,omitempty"`
	BodyText                  NullableJsonInput[string] `json:"bodyText,omitempty"`
	ContainerBackground       NullableJsonInput[string] `json:"containerBackground,omitempty"`
	ContainerBorder           NullableJsonInput[string] `json:"containerBorder,omitempty"`
	Divider                   NullableJsonInput[string] `json:"divider,omitempty"`
	Icon                      NullableJsonInput[string] `json:"icon,omitempty"`
	Loader                    NullableJsonInput[string] `json:"loader,omitempty"`
	Positive                  NullableJsonInput[string] `json:"positive,omitempty"`
	Critical                  NullableJsonInput[string] `json:"critical,omitempty"`
	Information               NullableJsonInput[string] `json:"information,omitempty"`
	Hover                     NullableJsonInput[string] `json:"hover,omitempty"`
	Focus                     NullableJsonInput[string] `json:"focus,omitempty"`
}

// ExitPosition sits on Container only. Where the exit control sits is theme-wide, so DarkMode takes
// ModeContainer instead: the API rejects an exitPosition under darkMode.
type Container struct {
	ContentAlignment NullableJsonInput[string] `json:"contentAlignment,omitempty"`
	Padding          NullableJsonInput[int64]  `json:"padding,omitempty"`
	LogoAlignment    NullableJsonInput[string] `json:"logoAlignment,omitempty"`
	LogoPosition     NullableJsonInput[string] `json:"logoPosition,omitempty"`
	LogoHeight       NullableJsonInput[int64]  `json:"logoHeight,omitempty"`
	ExitPosition     NullableJsonInput[string] `json:"exitPosition,omitempty"`
}

type ModeContainer struct {
	ContentAlignment NullableJsonInput[string] `json:"contentAlignment,omitempty"`
	Padding          NullableJsonInput[int64]  `json:"padding,omitempty"`
	LogoAlignment    NullableJsonInput[string] `json:"logoAlignment,omitempty"`
	LogoPosition     NullableJsonInput[string] `json:"logoPosition,omitempty"`
	LogoHeight       NullableJsonInput[int64]  `json:"logoHeight,omitempty"`
}

type PageBackground struct {
	BackgroundColor    NullableJsonInput[string] `json:"backgroundColor,omitempty"`
	BackgroundImageUrl NullableJsonInput[string] `json:"backgroundImageUrl,omitempty"`
}

// MaxFontFacesPerTypeface mirrors the cap the Management API enforces on each typeface's face list.
const MaxFontFacesPerTypeface = 6

// FontWeight is a single weight ("400") or an ascending range ("100 900"). The API stores a single
// weight as a number and a range as a string, so it can come back as either JSON type.
type FontWeight string

func (w FontWeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(w))
}

func (w *FontWeight) UnmarshalJSON(data []byte) error {
	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		*w = FontWeight(asString)
		return nil
	}

	var asNumber json.Number
	if err := json.Unmarshal(data, &asNumber); err != nil {
		return fmt.Errorf("font weight must be a string or a number, got %s", data)
	}

	*w = FontWeight(asNumber.String())

	return nil
}

type FontFace struct {
	Url    string     `json:"url"`
	Weight FontWeight `json:"weight,omitempty"`
}

type Typeface struct {
	Faces NullableJsonInput[[]FontFace] `json:"faces,omitempty"`
	// Deprecated: the single-url shape. Prefer Faces, which carries a weight per file.
	FontUrl NullableJsonInput[string] `json:"fontUrl,omitempty"`
}

// Typography sits on Theme only. A typeface is shared by both colour modes, so DarkMode has none.
type Typography struct {
	Text    NullableJsonInput[Typeface] `json:"text,omitempty"`
	Display NullableJsonInput[Typeface] `json:"display,omitempty"`
}

// Links and Shadows sit on Theme only. Neither is a design token, so DarkMode has neither.
type Links struct {
	Underline NullableJsonInput[bool] `json:"underline,omitempty"`
}

type Shadows struct {
	Enabled NullableJsonInput[bool] `json:"enabled,omitempty"`
}

type DarkMode struct {
	Borders        NullableJsonInput[Borders]        `json:"borders,omitempty"`
	Colors         NullableJsonInput[Colors]         `json:"colors,omitempty"`
	Container      NullableJsonInput[ModeContainer]  `json:"container,omitempty"`
	PageBackground NullableJsonInput[PageBackground] `json:"pageBackground,omitempty"`
	LogoUrl        NullableJsonInput[string]         `json:"logoUrl,omitempty"`
	WatermarkUrl   NullableJsonInput[string]         `json:"watermarkUrl,omitempty"`
	FaviconUrl     NullableJsonInput[string]         `json:"faviconUrl,omitempty"`
	PrimaryColor   NullableJsonInput[string]         `json:"primaryColor,omitempty"`
}

type Theme struct {
	Borders        NullableJsonInput[Borders]        `json:"borders,omitempty"`
	Colors         NullableJsonInput[Colors]         `json:"colors,omitempty"`
	Container      NullableJsonInput[Container]      `json:"container,omitempty"`
	PageBackground NullableJsonInput[PageBackground] `json:"pageBackground,omitempty"`
	Typography     NullableJsonInput[Typography]     `json:"typography,omitempty"`
	Links          NullableJsonInput[Links]          `json:"links,omitempty"`
	Shadows        NullableJsonInput[Shadows]        `json:"shadows,omitempty"`
	LogoUrl        NullableJsonInput[string]         `json:"logoUrl,omitempty"`
	WatermarkUrl   NullableJsonInput[string]         `json:"watermarkUrl,omitempty"`
	FaviconUrl     NullableJsonInput[string]         `json:"faviconUrl,omitempty"`
	PrimaryColor   NullableJsonInput[string]         `json:"primaryColor,omitempty"`
	DarkMode       NullableJsonInput[DarkMode]       `json:"darkMode,omitempty"`
	Name           NullableJsonInput[string]         `json:"name,omitempty"`
}

type BordersResponse struct {
	ButtonBorderRadius    int64 `json:"buttonBorderRadius"`
	ButtonBorderWidth     int64 `json:"buttonBorderWidth"`
	CardBorderRadius      int64 `json:"cardBorderRadius"`
	CardBorderWidth       int64 `json:"cardBorderWidth"`
	InputBorderRadius     int64 `json:"inputBorderRadius"`
	InputBorderWidth      int64 `json:"inputBorderWidth"`
	ContainerBorderRadius int64 `json:"containerBorderRadius"`
}

type ColorsResponse struct {
	ButtonPrimaryText         string `json:"buttonPrimaryText"`
	ButtonPrimaryBorder       string `json:"buttonPrimaryBorder"`
	ButtonSecondaryText       string `json:"buttonSecondaryText"`
	ButtonSecondaryBackground string `json:"buttonSecondaryBackground"`
	ButtonSecondaryBorder     string `json:"buttonSecondaryBorder"`
	CardBackground            string `json:"cardBackground"`
	CardBorder                string `json:"cardBorder"`
	InputBackground           string `json:"inputBackground"`
	InputBorder               string `json:"inputBorder"`
	Link                      string `json:"link"`
	HeadingText               string `json:"headingText"`
	BodyText                  string `json:"bodyText"`
	ContainerBackground       string `json:"containerBackground"`
	ContainerBorder           string `json:"containerBorder"`
	Divider                   string `json:"divider"`
	Icon                      string `json:"icon"`
	Loader                    string `json:"loader"`
	Positive                  string `json:"positive"`
	Critical                  string `json:"critical"`
	Information               string `json:"information"`
	Hover                     string `json:"hover"`
	Focus                     string `json:"focus"`
}

type ContainerResponse struct {
	ContentAlignment string `json:"contentAlignment"`
	Padding          int64  `json:"padding"`
	LogoAlignment    string `json:"logoAlignment"`
	LogoPosition     string `json:"logoPosition"`
	LogoHeight       int64  `json:"logoHeight"`
	ExitPosition     string `json:"exitPosition"`
}

type ModeContainerResponse struct {
	ContentAlignment string `json:"contentAlignment"`
	Padding          int64  `json:"padding"`
	LogoAlignment    string `json:"logoAlignment"`
	LogoPosition     string `json:"logoPosition"`
	LogoHeight       int64  `json:"logoHeight"`
}

type PageBackgroundResponse struct {
	BackgroundColor    string `json:"backgroundColor"`
	BackgroundImageUrl string `json:"backgroundImageUrl"`
}

type FontFaceResponse struct {
	Url    string     `json:"url"`
	Weight FontWeight `json:"weight"`
}

type TypefaceResponse struct {
	Faces   []FontFaceResponse `json:"faces"`
	FontUrl string             `json:"fontUrl"`
}

type TypographyResponse struct {
	Text    TypefaceResponse `json:"text"`
	Display TypefaceResponse `json:"display"`
}

// A pointer, because the API omits a switch the tenant never set and `false` is a value they set.
type LinksResponse struct {
	Underline *bool `json:"underline"`
}

type ShadowsResponse struct {
	Enabled *bool `json:"enabled"`
}

type DarkModeResponse struct {
	Borders        BordersResponse        `json:"borders"`
	Colors         ColorsResponse         `json:"colors"`
	Container      ModeContainerResponse  `json:"container"`
	PageBackground PageBackgroundResponse `json:"pageBackground"`
	LogoUrl        string                 `json:"logoUrl"`
	WatermarkUrl   string                 `json:"watermarkUrl"`
	FaviconUrl     string                 `json:"faviconUrl"`
	PrimaryColor   string                 `json:"primaryColor"`
}

type ThemeResponse struct {
	Borders        BordersResponse        `json:"borders"`
	Colors         ColorsResponse         `json:"colors"`
	Container      ContainerResponse      `json:"container"`
	PageBackground PageBackgroundResponse `json:"pageBackground"`
	Typography     TypographyResponse     `json:"typography"`
	Links          LinksResponse          `json:"links"`
	Shadows        ShadowsResponse        `json:"shadows"`
	LogoUrl        string                 `json:"logoUrl"`
	WatermarkUrl   string                 `json:"watermarkUrl"`
	FaviconUrl     string                 `json:"faviconUrl"`
	PrimaryColor   string                 `json:"primaryColor"`
	DarkMode       DarkModeResponse       `json:"darkMode"`
	Name           string                 `json:"name"`
}

func (c Client) GetTheme() (*ThemeResponse, int, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/theme", c.Host), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var theme ThemeResponse
	err = json.Unmarshal(body, &theme)
	if err != nil {
		return nil, statusCode, err
	}

	return &theme, statusCode, nil
}

func (c Client) UpdateTheme(theme Theme) (*ThemeResponse, int, error) {
	updateBody, err := json.MarshalIndent(theme, "", "\t")
	if err != nil {
		return nil, 0, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/theme", c.Host), bytes.NewReader(updateBody))
	if err != nil {
		return nil, 0, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, statusCode, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, statusCode, err
	}

	var updatedTheme ThemeResponse
	err = json.Unmarshal(body, &updatedTheme)
	if err != nil {
		return nil, statusCode, err
	}

	return &updatedTheme, statusCode, nil
}

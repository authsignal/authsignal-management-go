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

type Container struct {
	ContentAlignment NullableJsonInput[string] `json:"contentAlignment,omitempty"`
	Position         NullableJsonInput[string] `json:"position,omitempty"`
	Padding          NullableJsonInput[int64]  `json:"padding,omitempty"`
	LogoAlignment    NullableJsonInput[string] `json:"logoAlignment,omitempty"`
	LogoPosition     NullableJsonInput[string] `json:"logoPosition,omitempty"`
	LogoHeight       NullableJsonInput[int64]  `json:"logoHeight,omitempty"`
}

type PageBackground struct {
	BackgroundColor    NullableJsonInput[string] `json:"backgroundColor,omitempty"`
	BackgroundImageUrl NullableJsonInput[string] `json:"backgroundImageUrl,omitempty"`
}

type Display struct {
	FontUrl NullableJsonInput[string] `json:"fontUrl,omitempty"`
}

type Typography struct {
	Display NullableJsonInput[Display] `json:"display,omitempty"`
}

type DarkMode struct {
	Borders        NullableJsonInput[Borders]        `json:"borders,omitempty"`
	Colors         NullableJsonInput[Colors]         `json:"colors,omitempty"`
	Container      NullableJsonInput[Container]      `json:"container,omitempty"`
	PageBackground NullableJsonInput[PageBackground] `json:"pageBackground,omitempty"`
	Typography     NullableJsonInput[Typography]     `json:"typography,omitempty"`
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
	Position         string `json:"position"`
	Padding          int64  `json:"padding"`
	LogoAlignment    string `json:"logoAlignment"`
	LogoPosition     string `json:"logoPosition"`
	LogoHeight       int64  `json:"logoHeight"`
}

type PageBackgroundResponse struct {
	BackgroundColor    string `json:"backgroundColor"`
	BackgroundImageUrl string `json:"backgroundImageUrl"`
}

type DisplayResponse struct {
	FontUrl string `json:"fontUrl"`
}

type TypographyResponse struct {
	Display DisplayResponse `json:"display"`
}

type DarkModeResponse struct {
	Borders        BordersResponse        `json:"borders"`
	Colors         ColorsResponse         `json:"colors"`
	Container      ContainerResponse      `json:"container"`
	PageBackground PageBackgroundResponse `json:"pageBackground"`
	Typography     TypographyResponse     `json:"typography"`
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
	LogoUrl        string                 `json:"logoUrl"`
	WatermarkUrl   string                 `json:"watermarkUrl"`
	FaviconUrl     string                 `json:"faviconUrl"`
	PrimaryColor   string                 `json:"primaryColor"`
	DarkMode       DarkModeResponse       `json:"darkMode"`
	Name           string                 `json:"name"`
}

func (c Client) GetTheme() (*ThemeResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/theme", c.Host), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var theme ThemeResponse
	err = json.Unmarshal(body, &theme)
	if err != nil {
		return nil, err
	}

	return &theme, nil
}

func (c Client) UpdateTheme(theme Theme) (*ThemeResponse, error) {
	updateBody, err := json.MarshalIndent(theme, "", "\t")
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%s/theme", c.Host), bytes.NewReader(updateBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	body, err := c.makeRequest(request, c.ApiSecret)
	if err != nil {
		return nil, err
	}

	var updatedTheme ThemeResponse
	err = json.Unmarshal(body, &updatedTheme)
	if err != nil {
		return nil, err
	}

	return &updatedTheme, nil
}

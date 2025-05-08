package model

import (
	"github.com/JohnnyJa/AdServer/BidHandler/internal/requests"
	"github.com/google/uuid"
	"strings"
)

type Creative struct {
	ID           uuid.UUID
	MediaURL     string
	Width        int
	Height       int
	CreativeType string
}

func (c Creative) IsSettingsMatched(imp requests.Imp) bool {
	if imp.Banner != nil {
		return c.IsMatchedBannerSettings(imp)
	} else {
		return false
	}
}

func (c Creative) IsMatchedBannerSettings(imp requests.Imp) bool {
	if !c.IsBanner() {
		return false
	}

	if c.HasMediaURL() {
		return false
	}

	if !c.IsSizeMatched(imp) {
		return false
	}
	return true
}

func (c Creative) HasMediaURL() bool {
	return c.MediaURL == ""
}

func (c Creative) IsBanner() bool {
	return strings.EqualFold(c.CreativeType, "banner")
}

func (c Creative) IsSizeMatched(imp requests.Imp) bool {
	return imp.Banner.W >= c.Width && imp.Banner.H >= c.Height
}

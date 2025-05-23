package model

import (
	"github.com/JohnnyJa/AdServer/BidHandler/internal/requests"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/semanticTargetingService"
	"github.com/google/uuid"
	"math/rand"
	"strings"
)

type Profile struct {
	Id               uuid.UUID
	Name             string
	BidPrice         float32
	ProfileTargeting map[string]string
	Creatives        map[uuid.UUID]*Creative
}

type ProfileWithMatchedCreative struct {
	Id               uuid.UUID
	Name             string
	BidPrice         float32
	ProfileTargeting map[string]string
	Creative         *Creative
}

func (p *Profile) IsEligible(imp requests.Imp, service semanticTargetingService.SemanticTargetingService) bool {
	return p.IsBidPriceHigherThanBidFloor(imp) && p.IsTargetingMatched(imp, service)
}

func (p *Profile) IsBidPriceHigherThanBidFloor(imp requests.Imp) bool {
	return p.BidPrice >= imp.BidFloor
}

func (p *Profile) IsTargetingMatched(imp requests.Imp, service semanticTargetingService.SemanticTargetingService) bool {
	if targeting, exist := p.ProfileTargeting["ImpTargeting"]; exist {
		if imp.Ext == nil && imp.Ext.Targeting == nil {
			return false
		}

		return service.IsSimilar(strings.ToLower(imp.Ext.Targeting.Keyword), strings.ToLower(targeting))
	}

	return true
}

func (p *Profile) FindMatchedCreative(imp requests.Imp) *ProfileWithMatchedCreative {
	matchedCreatives := make([]*Creative, 0)
	for _, creative := range p.Creatives {
		if creative.IsSettingsMatched(imp) {
			matchedCreatives = append(matchedCreatives, creative)
		}
	}

	if len(matchedCreatives) == 0 {
		return nil
	}

	matchedCreative := chooseRandomCreative(matchedCreatives)
	return &ProfileWithMatchedCreative{
		Id:               p.Id,
		Name:             p.Name,
		BidPrice:         p.BidPrice,
		ProfileTargeting: p.ProfileTargeting,
		Creative:         matchedCreative,
	}
}

func chooseRandomCreative(matchedCreatives []*Creative) *Creative {
	index := rand.Intn(len(matchedCreatives))
	return matchedCreatives[index]
}

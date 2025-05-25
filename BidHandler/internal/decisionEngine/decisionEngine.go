package decisionEngine

import (
	"context"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/grpcClients"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/mapper"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/model"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/requests"
	"github.com/JohnnyJa/AdServer/BidHandler/internal/semanticTargetingService"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DecisionEngine interface {
	GetWinners(ctx context.Context, request requests.BidRequest) (*requests.BidResponse, error)
}

type decisionEngine struct {
	profileService grpcClients.ProfilesClient
	logger         *logrus.Logger

	request            requests.BidRequest
	profilesByPackage  map[uuid.UUID][]uuid.UUID
	profilesByUUID     map[uuid.UUID]*model.Profile
	semanticService    semanticTargetingService.SemanticTargetingService
	profileStateClient grpcClients.ProfileStateClient
}

func NewDecisionEngine(logger *logrus.Logger, profileService grpcClients.ProfilesClient, profileStateClient grpcClients.ProfileStateClient, semanticService semanticTargetingService.SemanticTargetingService) DecisionEngine {
	return &decisionEngine{
		logger:             logger,
		profileService:     profileService,
		profileStateClient: profileStateClient,
		semanticService:    semanticService,
	}
}

func (d *decisionEngine) GetWinners(ctx context.Context, request requests.BidRequest) (*requests.BidResponse, error) {
	d.request = request

	winners := make(map[uuid.UUID]*model.ProfileWithMatchedCreative)

	for _, imp := range request.Imp {
		err := d.getProfiles(ctx, imp.TagID)
		if err != nil {
			return &requests.BidResponse{}, err
		}

		winner := d.makeDecision(ctx, imp)

		if winner != nil {
			winners[uuid.MustParse(imp.ID)] = winner
		}
	}

	if len(winners) == 0 {
		return nil, nil
	}

	return mapper.NewBidResponse(request, winners), nil
}

func (d *decisionEngine) makeDecision(ctx context.Context, imp requests.Imp) *model.ProfileWithMatchedCreative {
	eligibleProfiles := d.findEligibleProfiles(ctx, imp)
	winnerProfile := PerformAuction(eligibleProfiles)
	return winnerProfile
}

func PerformAuction(profiles map[uuid.UUID]*model.ProfileWithMatchedCreative) *model.ProfileWithMatchedCreative {
	var winnerProfile *model.ProfileWithMatchedCreative
	highestPrice := float32(0.0)
	for _, profile := range profiles {
		if profile.BidPrice > highestPrice {
			winnerProfile = profile
			highestPrice = profile.BidPrice
		}
	}
	return winnerProfile
}

func (d *decisionEngine) getProfiles(ctx context.Context, id uuid.UUID) error {
	profilesByPackages, profilesByUUID, err := d.profileService.GetProfilesMapsByZone(ctx, id)
	if err != nil {
		return err
	}

	d.profilesByUUID = mapper.ProtoToProfilesByUUID(profilesByUUID)
	d.profilesByPackage = mapper.ProtoToProfilesByPackages(profilesByPackages)
	return nil
}

func (d *decisionEngine) findEligibleProfiles(ctx context.Context, imp requests.Imp) map[uuid.UUID]*model.ProfileWithMatchedCreative {
	result := make(map[uuid.UUID]*model.ProfileWithMatchedCreative)
	for id, profile := range d.profilesByUUID {
		if d.IsProfileActive(ctx, id) && profile.IsEligible(imp, d.semanticService) {
			matchedProfile := profile.FindMatchedCreative(imp)
			if matchedProfile != nil {
				result[id] = matchedProfile
			}
		}
	}
	return result
}

func (d *decisionEngine) IsProfileActive(ctx context.Context, id uuid.UUID) bool {
	state, err := d.profileStateClient.GetProfileState(ctx, id)
	if err != nil {
		logrus.Error(err)
		return false
	}

	return state == 0
}

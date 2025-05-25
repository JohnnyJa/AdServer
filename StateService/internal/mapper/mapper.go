package mapper

import (
	"github.com/JohnnyJa/AdServer/StateService/internal/grpcClients/proto"
	"github.com/google/uuid"
)

func ProtoProfileWithLimitsToMap(limits []*proto.ProfilesWithLimits) map[uuid.UUID]int {
	res := make(map[uuid.UUID]int)
	for _, profileWithLimits := range limits {
		res[uuid.MustParse(profileWithLimits.Id)] = int(profileWithLimits.ViewsLimits)
	}
	return res
}

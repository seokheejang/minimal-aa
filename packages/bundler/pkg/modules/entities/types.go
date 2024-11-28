package entities

import (
	"github.com/ethereum/go-ethereum/common"
)

type ReputationOverride struct {
	Address     common.Address `json:"address"`
	OpsSeen     int            `json:"opsSeen"`
	OpsIncluded int            `json:"opsIncluded"`
}

type ReputationConstants struct {
	MinUnstakeDelay                int
	MinStakeValue                  int64
	SameSenderMempoolCount         int
	SameUnstakedEntityMempoolCount int
	ThrottledEntityMempoolCount    int
	ThrottledEntityLiveBlocks      int
	ThrottledEntityBundleCount     int
	MinInclusionRateDenominator    int
	ThrottlingSlack                int
	BanSlack                       int
}

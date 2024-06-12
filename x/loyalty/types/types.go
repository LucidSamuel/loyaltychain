package types

import (
    "github.com/cosmos/gogoproto/proto"
)

// EventType and AttributeKey types
type EventType string
type AttributeKey string

// Event type and attribute key constants
const (
    EventTypeRedeemReward EventType = "RedeemReward"
    AttributeKeyRewardItem AttributeKey = "reward"
    AttributeKeyOwner      AttributeKey = "owner"
    AttributeKeyPoints     AttributeKey = "points"
)

// Reward defines the reward structure
type Reward struct {
    Item   string `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
    Points int32  `protobuf:"varint,2,opt,name=points,proto3" json:"points,omitempty"`
}

func (r *Reward) Reset()         { *r = Reward{} }
func (r *Reward) String() string { return proto.CompactTextString(r) }
func (*Reward) ProtoMessage()    {}

// Explicitly use the proto package to avoid "imported and not used" error
var _ proto.Message = &Reward{}

// QueryPointsBalanceRequest is the request type for the Query/PointsBalance RPC method.
type QueryPointsBalanceRequest struct {
    Owner string `json:"owner"`
}

// QueryPointsBalanceResponse is the response type for the Query/PointsBalance RPC method.
type QueryPointsBalanceResponse struct {
    Balance int32 `json:"balance"`
}

// QueryRewardDetailsRequest is the request type for the Query/RewardDetails RPC method.
type QueryRewardDetailsRequest struct {
    RewardItem string `json:"reward_item"`
}

// QueryRewardDetailsResponse is the response type for the Query/RewardDetails RPC method.
type QueryRewardDetailsResponse struct {
    Points int32 `json:"points"`
}
 
package loyalty

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "loyaltychain/api/loyaltychain/loyalty"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "PointAll",
					Use:       "list-point",
					Short:     "List all Point",
				},
				{
					RpcMethod:      "Point",
					Use:            "show-point [id]",
					Short:          "Shows a Point",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePoint",
					Use:            "create-point [index] [owner] [balance]",
					Short:          "Create a new Point",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "balance"}},
				},
				{
					RpcMethod:      "UpdatePoint",
					Use:            "update-point [index] [owner] [balance]",
					Short:          "Update Point",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "owner"}, {ProtoField: "balance"}},
				},
				{
					RpcMethod:      "DeletePoint",
					Use:            "delete-point [index]",
					Short:          "Delete Point",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

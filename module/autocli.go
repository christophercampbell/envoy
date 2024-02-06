package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	envoyv1 "github.com/polygon/envoy/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: envoyv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetLock",
					Use:       "get-lock by name",
					Short:     "Get the current state of the named lock",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
					},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: envoyv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateLock",
					Use:       "create a named lock for an envoy address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "name"},
						{ProtoField: "envoy"},
						{ProtoField: "at_block"},
						{ProtoField: "num_blocks"},
					},
				},
			},
		},
	}
}

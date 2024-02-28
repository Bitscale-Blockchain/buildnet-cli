package dex

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "bitscale/api/bitscale/dex"
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
					RpcMethod:      "DexGetPool",
					Use:            "dex-get-pool [pool-id]",
					Short:          "Query dexGetPool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "poolId"}},
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
					RpcMethod:      "DexSwap",
					Use:            "dex-swap [asset-one] [asset-two] [amount-in] [amount-out]",
					Short:          "Send a dexSwap tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "assetOne"}, {ProtoField: "assetTwo"}, {ProtoField: "amountIn"}, {ProtoField: "amountOut"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

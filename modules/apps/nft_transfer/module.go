package nft_transfer

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/ibc-go/modules/apps/nft_transfer/keeper"
	"github.com/cosmos/ibc-go/modules/apps/nft_transfer/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModule represents the AppModule for this module
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

func (a AppModule) InitGenesis(context sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	panic("implement me")
}

func (a AppModule) ExportGenesis(context sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	panic("implement me")
}

func (a AppModule) RegisterInvariants(registry sdk.InvariantRegistry) {
	panic("implement me")
}

func (a AppModule) Route() sdk.Route {
	panic("implement me")
}

func (a AppModule) QuerierRoute() string {
	panic("implement me")
}

func (a AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	panic("implement me")
}

func (a AppModule) RegisterServices(configurator module.Configurator) {
	panic("implement me")
}

func (a AppModule) ConsensusVersion() uint64 {
	panic("implement me")
}

func (a AppModule) BeginBlock(context sdk.Context, block abci.RequestBeginBlock) {
	panic("implement me")
}

func (a AppModule) EndBlock(context sdk.Context, block abci.RequestEndBlock) []abci.ValidatorUpdate {
	panic("implement me")
}

// AppModuleBasic is the IBC Transfer AppModuleBasic
type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	panic("implement me")
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	panic("implement me")
}

func (a AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	panic("implement me")
}

func (a AppModuleBasic) DefaultGenesis(codec codec.JSONCodec) json.RawMessage {
	panic("implement me")
}

func (a AppModuleBasic) ValidateGenesis(codec codec.JSONCodec, config client.TxEncodingConfig, message json.RawMessage) error {
	panic("implement me")
}

func (a AppModuleBasic) RegisterRESTRoutes(context client.Context, router *mux.Router) {
	panic("implement me")
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {
	panic("implement me")
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	panic("implement me")
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	panic("implement me")
}

// OnRecvPacket implements the IBCModule interface. A successful acknowledgement
// is returned if the packet data is succesfully decoded and the receive application
// logic returns without error.
func (am AppModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
){

	ack := ""
	var data types.NonFungibleTokenPacketData
	// unmarshal data
	// todo

	// only attempt the application logic if the packet data  was successfully decoded
	if ack.success{
		err := am.keeper.OnRecvPacket(ctx, packet, data)
		if err != nil {
			ack = ""
		}
	}

}

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/modules/apps/nft_transfer/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"strings"
)

const (
	Prefix = "tibc/nft"
)


func (k Keeper) SendNftTransfer(
	ctx sdk.Context,
	class, id, uri string,
	sender sdk.AccAddress,
	receiver string,
	awayFromOrigin bool,
	destChain, relayChain string,
) error {

	// get the next sequence
	// todo

	if strings.HasPrefix(class, Prefix){

	}

	if awayFromOrigin{
		// lock nft
		if err := k.nftKeeper.TransferOwner(ctx, class, id, "", uri, "",
			sender, k.GetNftTransferModuleAddr(types.ModuleName)); err != nil{
			return err
		}
	} else {
		// burn nft
		if err := k.nftKeeper.BurnNFT(ctx, class, id,
			k.GetNftTransferModuleAddr(types.ModuleName)); err != nil{
			return err
		}
	}

	// packetdata
	// todo
	_ = types.NewNonFungibleTokenPacketData(class, id, uri,
		sender.String(), receiver, awayFromOrigin)


	// constructs packet
	// todo

	// send packet
	// todo

	return nil
}


func (k Keeper)OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, data types.NonFungibleTokenPacketData) error{
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return err
	}

	// decode the sender address
	sender, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return err
	}

	// decode the receiver address
	receiver, err := sdk.AccAddressFromBech32(data.Receiver)
	if err != nil {
		return err
	}


	//
	if data.AwayFromOrigin{
		if strings.HasPrefix(data.Class, Prefix){
			// has prefix  tibc/nft/a/b/class
			classSplit := strings.Split(data.Class, "/")
			classSplit = append(classSplit[:len(classSplit) - 2], packet.sourceChain)
			newClass := strings.Join(classSplit, "/")
			if err := k.nftKeeper.MintNFT(ctx, newClass, data.Id, "", data.Uri, "", sender); err != nil{
				return err
			}
		} else {
			// not has prefix  tibc/nft/a/class
			newClass := Prefix + "/" + packet.sourceChain + data.Class
			if err := k.nftKeeper.MintNFT(ctx, newClass, data.Id, "", data.Uri, "", sender); err != nil{
				return err
			}
			// lock todo
			// send packet  need judge relay chain empty todo
		}
	} else {
		if strings.HasPrefix(data.Class, Prefix){
			classSplit := strings.Split(data.Class, "/")
			destChain := classSplit[len(classSplit) - 2]
			if destChain != packet.destChain{
				// return err  must equal
			}
			if len(classSplit) == 4{
				newClass := classSplit[len(classSplit) - 1]
			} else {
				newClass :=
			}

		} else {
			//  return err must has prefix if awayfromchain todo
		}
	}
	return nil
}
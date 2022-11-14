package types

import (
	rules "github.com/Mesorem/ttt_blockchain/x/tttblockchain/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidCross     = sdkerrors.Register(ModuleName, 1103, "cross address is invalid")
	ErrInvalidCircle    = sdkerrors.Register(ModuleName, 1101, "circle address is invalid")
	ErrGameNotParseable = sdkerrors.Register(ModuleName, 1102, "game cant be parsed to struct")
)

func (storedGame StoredGame) GetCrossAddress() (cross sdk.AccAddress, err error) {
	cross, errCcross := sdk.AccAddressFromBech32(storedGame.Cross)
	return cross, sdkerrors.Wrapf(errCcross, ErrInvalidCross.Error(), storedGame.Cross)
}

func (storedGame StoredGame) GetCircleAddress() (cross sdk.AccAddress, err error) {
	circle, errCircle := sdk.AccAddressFromBech32(storedGame.Circle)
	return circle, sdkerrors.Wrapf(errCircle, ErrInvalidCircle.Error(), storedGame.Circle)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, errBoard error) {
	game, errBoard = rules.Parse(storedGame.Board)
	if errBoard != nil {
		return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error(), storedGame.Board)
	}
	return
}

func (storedGame StoredGame) Validate() (err error) {
	_, err = storedGame.GetCrossAddress()
	if err != nil {
		return
	}

	_, err = storedGame.GetCircleAddress()
	if err != nil {
		return
	}

	_, err = storedGame.ParseGame()
	return
}

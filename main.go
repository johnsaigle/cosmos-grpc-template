package main

// From: https://docs.cosmos.network/main/run-node/interact-node#programmatically-via-go

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// FIXME: Replace with corresponding types from your target module
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func queryState() error {
	// FIXME: Add validator or recipient address from your local validator
	myAddress, err := sdk.AccAddressFromBech32("cosmos1...") // the my_validator or recipient address.
	if err != nil {
		return err
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		// FIXME: replace with codec, if necessary
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	// FIXME: Modify *types and *Request with your actual targets
	// This creates a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: myAddress.String(), Denom: "stake"},
	)
	if err != nil {
		return err
	}

	// FIXME: Replace with the output you want
	fmt.Println(bankRes.GetBalance()) // Prints the account balance

	return nil
}

func main() {
	if err := queryState(); err != nil {
		panic(err)
	}
}

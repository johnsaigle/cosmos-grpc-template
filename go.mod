module cosmos-grpc-template/grpc

go 1.18

// use informal system fork of tendermint
replace github.com/tendermint/tendermint => github.com/informalsystems/tendermint v0.34.26

// use cosmos fork of unmaintained protobuf
replace github.com/gogo/protobuf v1.3.2 => github.com/cosmos/gogoproto v1.4.5

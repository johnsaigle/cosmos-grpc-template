module cosmos-grpc-template/grpc

go 1.18

// use informal system fork of tendermint
replace github.com/tendermint/tendermint => github.com/informalsystems/tendermint v0.34.26

// use regen-network fork of unmaintained protobuf
replace github.com/gogo/protobuf v1.3.2 => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

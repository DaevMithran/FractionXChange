#!/usr/bin/env bash

mockgen_cmd="mockgen"
$mockgen_cmd -source=x/nft/types/expected_keepers.go -package testutil -destination x/nft/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/guard/types/expected_keepers.go -package testutil -destination x/guard/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/token/types/expected_keepers.go -package testutil -destination x/token/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/did/types/expected_keepers.go -package testutil -destination x/did/testutil/expected_keepers_mocks.go
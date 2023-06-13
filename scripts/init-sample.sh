#!/bin/bash

CHAINID="mantrachain"
HOMEDIR="$HOME/.mantrachain"

KEYS[0]="validator"
KEYS[1]="recipient"
KEYS[2]="admin"
ACCOUNT_PRIVILEGES_GUARD_NFT_COLLECTION_ID="account_privileges_guard_nft_collection"

KEYRING="test"
LOGLEVEL="info"

GAS_ADJ=2
GAS_PRICE=0.0001uaum

source "$PWD"/scripts/common.sh

set -e # exit on first error

VALIDATOR_WALLET=$("$PWD"/build/mantrachaind keys show ${KEYS[0]} -a --keyring-backend $KEYRING --home "$HOMEDIR")
RECIPIENT_WALLET=$("$PWD"/build/mantrachaind keys show ${KEYS[1]} -a --keyring-backend $KEYRING --home "$HOMEDIR")
ADMIN_WALLET=$("$PWD"/build/mantrachaind keys show ${KEYS[2]} -a --keyring-backend $KEYRING --home "$HOMEDIR")

cecho "CYAN" "Send aum from ${KEYS[0]} to ${KEYS[1]}"
"$PWD"/build/mantrachaind tx bank send $VALIDATOR_WALLET $RECIPIENT_WALLET 100000000000000uaum --chain-id $CHAINID --keyring-backend $KEYRING --gas auto --gas-adjustment $GAS_ADJ --gas-prices $GAS_PRICE --home "$HOMEDIR" -y

sleep 7

cecho "CYAN" "Send uaum from ${KEYS[0]} to ${KEYS[2]}"
"$PWD"/build/mantrachaind tx bank send $VALIDATOR_WALLET $ADMIN_WALLET 100000000000000uaum --chain-id $CHAINID --keyring-backend $KEYRING --gas auto --gas-adjustment $GAS_ADJ --gas-prices $GAS_PRICE --home "$HOMEDIR" -y

sleep 7

ACCOUNT_PRIVILEGES_GUARD_COLL_JSON=$(echo '{"id":"{id}","soul_bonded_nfts": true,"restricted_nfts":true,"name":"AccountPrivilegesCollection","description":"AccountPrivilegesCollection","category":"utility"}' | sed -e "s/{id}/$ACCOUNT_PRIVILEGES_GUARD_NFT_COLLECTION_ID/g")

cecho "CYAN" "Create account privileges guard nft collection"
"$PWD"/build/mantrachaind tx token create-nft-collection "$(echo $ACCOUNT_PRIVILEGES_GUARD_COLL_JSON)" --chain-id $CHAINID --from ${KEYS[2]} --keyring-backend $KEYRING --gas auto --gas-adjustment $GAS_ADJ --gas-prices $GAS_PRICE --yes --home "$HOMEDIR"

sleep 7

ACCOUNT_PRIVILEGES_GUARD_NFT_VALIDATOR_JSON=$(echo '{"id":"{id}","title":"AccountPrivileges","description":"AccountPrivileges"}' | sed -e "s/{id}/$VALIDATOR_WALLET/g")

ACCOUNT_PRIVILEGES_GUARD_NFT_ADMIN_JSON=$(echo '{"id":"{id}","title":"AccountPrivileges","description":"AccountPrivileges"}' | sed -e "s/{id}/$ADMIN_WALLET/g")

cecho "CYAN" "Mint account privileges guard nfts"
"$PWD"/build/mantrachaind tx token mint-nft "$(echo $ACCOUNT_PRIVILEGES_GUARD_NFT_VALIDATOR_JSON)" --collection-creator $ADMIN_WALLET --collection-id $ACCOUNT_PRIVILEGES_GUARD_NFT_COLLECTION_ID --chain-id $CHAINID --from ${KEYS[2]} --receiver $VALIDATOR_WALLET --keyring-backend $KEYRING --gas auto --gas-adjustment $GAS_ADJ --gas-prices $GAS_PRICE --home "$HOMEDIR" --yes

sleep 7

"$PWD"/build/mantrachaind tx token mint-nft "$(echo $ACCOUNT_PRIVILEGES_GUARD_NFT_ADMIN_JSON)" --collection-creator $ADMIN_WALLET --collection-id $ACCOUNT_PRIVILEGES_GUARD_NFT_COLLECTION_ID --chain-id $CHAINID --from ${KEYS[2]} --keyring-backend $KEYRING --gas auto --gas-adjustment $GAS_ADJ --gas-prices $GAS_PRICE --home "$HOMEDIR" --yes

sleep 7

cecho "CYAN" "Update guard transfer coins"
"$PWD"/build/mantrachaind tx guard update-guard-transfer-coins true --chain-id $CHAINID --from ${KEYS[2]} --keyring-backend $KEYRING --gas auto --gas-adjustment $GAS_ADJ --gas-prices $GAS_PRICE --home "$HOMEDIR" --yes
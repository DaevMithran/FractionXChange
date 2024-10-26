#!/bin/bash

# Define colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[1;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Error handling function
handle_error() {
  echo -e "${RED}Error: $1${NC}" >&2
  exit 1
}

# Function to display colorful headers
print_header() {
  echo -e "\n${CYAN}$1${NC}"
}

# Function to populate NFTs
populate_nfts() {
  print_header "Populating NFTs..."
  nfts=(
    "Art|Mona Lisa|A painting by Leonardo da Vinci|https://upload.wikimedia.org/wikipedia/commons/6/6a/Mona_Lisa.jpg|1"
    "Real_Estate|Dubai Museum Of Future|A futuristic museum with constant funding and revenue|https://media.istockphoto.com/id/1395117280/photo/the-museum-of-the-future-dubai-uae.webp?s=2048x2048&w=is&k=20&c=y4fm92ak37sLNTug0rw0XST-dvKyrIVImx-Q_kiWTgA=|2"
  )

  minted_nfts=()
  
  for nft in "${nfts[@]}"; do
    IFS='|' read -r category name description image id <<< "$nft"
    ./build/mantrachaind tx fractionnft mint "$name" "$id" "$description" "$image" "$category" --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto --yes
    if [ $? -ne 0 ]; then
      handle_error "Failed to mint NFT: $name"
    fi

    minted_nfts+=("$category|$name|$description|$image|$id")
    echo -e "${GREEN}Minted NFT:${NC} $name with ID: $id in Category: $category"
  done
}

# Function to query and display NFT and token owners
query_owners() {
  local category=$1
  local nft_id=$2

  print_header "Querying Owners of NFT and Tokens for Category: $category, NFT ID: $nft_id"

  # Query NFT owner
  nft_result=$(./build/mantrachaind query nft owner "$category" "$nft_id" -o json)
  if [ $? -ne 0 ]; then
    handle_error "Failed to retrieve NFT owner for $category - $nft_id"
  fi
  nft_owner=$(echo "$nft_result" | jq -r '.owner')
  echo -e "${YELLOW}NFT Owner:${NC} ${GREEN}$nft_owner${NC}"

  # Query Token owners
  token_result=$(./build/mantrachaind query bank denom-owners "fractionNFT-$category-$nft_id" -o json)
  if [ $? -ne 0 ]; then
    handle_error "Failed to retrieve token owners for $category - $nft_id"
  fi

  # Display token ownership details
  echo -e "${YELLOW}Token Owners:${NC}"
  denom_owners_count=$(echo "$token_result" | jq '.denom_owners | length')
  
  if [ "$denom_owners_count" -eq 0 ]; then
    echo -e "${RED}No token owners found for fractionNFT-${category}-${nft_id}.${NC}"
  else
    echo "$token_result" | jq -r '.denom_owners[] | "Address: \(.address), Balance: \(.balance.amount) \(.balance.denom)"' | while read line; do
      echo -e "${GREEN}$line${NC}"
    done
  fi
}

# Function to tokenize NFT
tokenize_nft() {
  local category=$1
  local nft_id=$2
  read -p "Enter the token supply: " token_supply
  read -p "Enter the timeout height: " timeout_height
  print_header "Tokenizing NFT with Category: $category, ID: $nft_id"
  
  tx_result=$(./build/mantrachaind tx fractionnft tokenize "$category" "$nft_id" "$token_supply" "$timeout_height" --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto --yes)
  if [ $? -ne 0 ]; then
    handle_error "Failed to tokenize NFT: $nft_id in Category: $category"
  fi
  sleep 5
}

# Function to transfer fractions of NFT
transfer_fractions() {
  local category=$1
  local nft_id=$2
  read -p "Fractions to Transfer: " amount
  print_header "Transferring $amount fractions of NFT with ID: $nft_id in Category: $category"
  
  tx_result=$(./build/mantrachaind tx bank send mantra1hj5fveer5cjtn4wd6wstzugjfdxzl0xpd8ck0e mantra10e354cme55fph398fhfjt8ujcme9qz9j8eumpd "${amount}fractionNFT-${category}-${nft_id}" --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto --yes)
  if [ $? -ne 0 ]; then
    handle_error "Failed to transfer NFT fractions for $nft_id in Category: $category"
  fi
  sleep 3
}

# Main sequence
populate_nfts

# Set the NFT details for the demo (assuming first NFT for simplicity)
IFS='|' read -r demo_category demo_name demo_description demo_image demo_id <<< "${minted_nfts[0]}"

query_owners "$demo_category" "$demo_id"
tokenize_nft "$demo_category" "$demo_id"
query_owners "$demo_category" "$demo_id"
transfer_fractions "$demo_category" "$demo_id"
query_owners "$demo_category" "$demo_id"

# Wait for the transaction to complete
print_header "Waiting for timeout"

# Allow users to decide to wait longer
read -p "Click enter to continue?"

query_owners "$demo_category" "$demo_id"

# Final output for clarity
print_header "All operations completed successfully!"

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

# Function to display colorful section headers
print_section() {
  echo -e "\n${CYAN}########################################################################################"
  echo -e "                         $1"
  echo -e "#######################################################################################${NC}\n"
}

print_title() {
  echo -e "\n${YELLOW}########################################################################################"
  echo -e "${YELLOW}#                                                                                      #"
  echo -e "#                              ${GREEN}*** FRACTIONAL NFT DEMO ***${YELLOW}                             #"
  echo -e "#                                                                                      #"
  echo -e "########################################################################################${NC}\n"
}

# Function to populate NFTs
populate_nfts() {
  print_section "Populate NFTs"
  
  nfts=(
    "Renewables|Wind Turbine 001|A wind turbine turns wind energy into electricity|https://solar.io|1"
    "Real_Estate|Dubai Museum Of Future|A futuristic museum with constant funding and revenue|https://dubai_museum|2"
  )

  minted_nfts=()
  
  for nft in "${nfts[@]}"; do
    IFS='|' read -r category name description image id <<< "$nft"
    ./build/mantrachaind tx fractionnft mint "$name" "$id" "$description" "$image" "$category" --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto --yes
    if [ $? -ne 0 ]; then
      handle_error "Failed to mint NFT: $name"
    fi
    sleep 2
    minted_nfts+=("$category|$name|$description|$image|$id")
    echo -e "${GREEN}Minted NFT:${NC} $name with ID: $id in Category: $category"
  done
}

# Function to query and display NFT and token owners with conditional status
query_owners() {
  local category=$1
  local nft_id=$2

  echo -e "\n${YELLOW}---------------------------------${NC}"
  
  # Query Token owners
  token_result=$(./build/mantrachaind query bank denom-owners "fractionNFT-$category-$nft_id" -o json)
  if [ $? -ne 0 ]; then
    handle_error "Failed to retrieve token owners for $category - $nft_id"
  fi

  denom_owners_count=$(echo "$token_result" | jq '.denom_owners | length')

  # Display NFT owner status based on token owners' existence
  if [ "$denom_owners_count" -eq 0 ]; then
    nft_result=$(./build/mantrachaind query nft owner "$category" "$nft_id" -o json)
    if [ $? -ne 0 ]; then
      handle_error "Failed to retrieve NFT owner for $category - $nft_id"
    fi
    nft_owner=$(echo "$nft_result" | jq -r '.owner')
    echo -e "${YELLOW}NFT Status:${NC} ${GREEN}Active - Owned by $nft_owner${NC}"
  else
    echo -e "${YELLOW}NFT Status:${NC} ${RED}Locked - NFT has been tokenized${NC}"
  fi

  # Display token ownership details
  if [ "$denom_owners_count" -eq 0 ]; then
    echo -e "${YELLOW}Token Owners:${NC} ${RED}Token Denom fractionNFT-${category}-${nft_id} does not exist${NC}"
  else
    echo -e "${YELLOW}Token Owners:${NC}"
    echo "$token_result" | jq -r '.denom_owners[] | "Address: \(.address), Balance: \(.balance.amount) \(.balance.denom)"' | while read line; do
    echo -e "${GREEN}$line${NC}"
    done
  fi

  echo -e "${YELLOW}---------------------------------${NC}"

  # Ask user if they want to continue with a default "yes"
  read -p "Do you want to continue to the next step? (yes/no) [default: yes]: " continue_choice
  continue_choice="${continue_choice:-yes}"
  if [[ "$continue_choice" != "yes" ]]; then
    echo -e "${RED}Exiting the script...${NC}"
    exit 0
  fi
}

# Function to tokenize NFT
tokenize_nft() {
  local category=$1
  local nft_id=$2

  print_section "Tokenize NFT with Category: $category, ID: $nft_id"

  read -p "Enter the token supply: " token_supply
  read -p "Enter the timeout height: " timeout_height
  
  tx_result=$(./build/mantrachaind tx fractionnft tokenize "$category" "$nft_id" "$token_supply" "$timeout_height" --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto --yes)
  if [ $? -ne 0 ]; then
    handle_error "Failed to tokenize NFT: $nft_id in Category: $category"
  fi

  sleep 3
}

# Function to transfer fractions of NFT
transfer_fractions() {
  local category=$1
  local nft_id=$2
  print_section "Let's Transfer some NFT fractions with ID: $nft_id in Category: $category"
  
  read -p "Fractions to Transfer: " amount

  tx_result=$(./build/mantrachaind tx bank send mantra1hj5fveer5cjtn4wd6wstzugjfdxzl0xpd8ck0e mantra10e354cme55fph398fhfjt8ujcme9qz9j8eumpd "${amount}fractionNFT-${category}-${nft_id}" --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto --yes)
  if [ $? -ne 0 ]; then
    handle_error "Failed to transfer NFT fractions for $nft_id in Category: $category"
  fi

  sleep 3
}

# Call the 3D title at the start of the script
print_title

sleep 3

# Main sequence
populate_nfts

# Set the NFT details for the demo (assuming first NFT for simplicity)
IFS='|' read -r demo_category demo_name demo_description demo_image demo_id <<< "${minted_nfts[0]}"

query_owners "$demo_category" "$demo_id"

tokenize_nft "$demo_category" "$demo_id"
query_owners "$demo_category" "$demo_id"

transfer_fractions "$demo_category" "$demo_id"
query_owners "$demo_category" "$demo_id"

# Wait for transaction inclusion at the end of the demo
print_section "Waiting for Timeout"
while true; do
  query_owners "$demo_category" "$demo_id"
  read -p "Do you want to query owners again? (yes/no) [default: yes]: " repeat_query
  repeat_query="${repeat_query:-yes}"
  if [[ "$repeat_query" != "yes" ]]; then
    break
  fi
done

# Final output for clarity
print_section "All Operations Completed Successfully!"

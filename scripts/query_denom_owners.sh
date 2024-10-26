#!/bin/bash

# Request for inputs
read -p "Enter the category: " category
read -p "Enter the NFT ID: " nft_id

# Query the Token owners
token_result=$(./build/mantrachaind query bank denom-owners "fractionNFT-$category-$nft_id" -o json)
if [ $? -ne 0 ]; then
  handle_error "Failed to retrieve token owners"
fi

# Display the result in green
echo -e "\e[32m$(echo "$token_result" | jq .)\e[0m"

# Query the NFT owner
nft_result=$(./build/mantrachaind query nft owner "$category" "$nft_id" -o json)
if [ $? -ne 0 ]; then
  handle_error "Failed to retrieve nft owner"
fi

# Display the result in green
echo -e "\e[32m$(echo "$nft_result" | jq .)\e[0m"
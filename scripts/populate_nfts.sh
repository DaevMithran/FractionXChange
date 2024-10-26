#!/bin/bash

# Hardcoded NFT data
nfts=(
  "Art|Mona Lisa|A painting by Leonardo da Vinci|https://upload.wikimedia.org/wikipedia/commons/6/6a/Mona_Lisa.jpg|1"
  "Real_Estate|Dubai Museum Of Future|A futuristic museum with constant funding and revenue|https://media.istockphoto.com/id/1395117280/photo/the-museum-of-the-future-dubai-uae.webp?s=2048x2048&w=is&k=20&c=y4fm92ak37sLNTug0rw0XST-dvKyrIVImx-Q_kiWTgA=|2"
)

minted_nfts=()

for nft in "${nfts[@]}"; do
    IFS='|' read -r category name description image id <<< "$nft"
    
    # Check if variables are correctly set
    if [ -z "$category" ] || [ -z "$name" ] || [ -z "$description" ] || [ -z "$image" ] || [ -z "$id" ]; then
        echo "Error parsing NFT details: $nft"
        continue
    fi

    # Execute the command
    ./build/mantrachaind tx fractionnft mint "$name" "$id" "$description" "$image" "$category" --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto --yes

    # Sleep period
    sleep 2

    # Add NFT details to the array
    minted_nfts+=("$category|$name|$description|$image|$id")

    echo "Minted NFT: $name with ID: $id in Category: $category"
done

# Display the minted NFTs in a table
echo -e "\nMinted NFTs:"
echo -e "Category\t| Name\t| Description\t| Image URL\t| ID"
echo -e "-------------------------------------------------------------"
for nft in "${minted_nfts[@]}"; do
    IFS='|' read -r category name description image id <<< "$nft"
    echo -e "$category\t| $name\t| $description\t| $image\t| $id"
done

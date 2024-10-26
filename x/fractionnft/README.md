# FractionNFT Module

## Overview

FractionNFT is a cutting-edge module that allows the fractionalization of NFTs (Non-Fungible Tokens) into smaller, tradable tokens. This module not only democratizes NFT ownership but also introduces a rental mechanism, enabling users to rent fractions of NFTs for a set period defined by the creator.

Using the FractionNFT module is much cheaper than issuing and managing multiple NFTs with expiry dates. This streamlined process not only saves on transaction fees but also simplifies the management of fractional ownership and rentals. Efficient and cost-effective!

## Features

Fractional Ownership: Split an NFT into multiple tokens, each representing a fraction of the original asset.

Rental Mechanism: Rent out fractions of an NFT for a defined period.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/MANTRA-Chain/mantrachain.git
   cd mantrachain
   ```

2. Build the project:
   ```bash
   make install
   ```

### Demo

1. Spawn the network:

   ```bash
   make sh-testnet
   ```

2. Run the demo:

   ```bash
   make demo
   ```

## Commands

### Tokenizing an NFT

To fractionalize an NFT into tokens, use:

```sh
./build/mantrachaind tx fractionnft tokenize <class_id> <nft_id> <token_supply> <timeout_height> --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto
```

- class_id: Category of the NFT.
- nft_id: ID of the NFT to be tokenized.
- token_supply: Total supply of tokens to be created.
- timeout_height: Timeout height for token expiry.

### Renting an NFT Fraction

To rent a fraction of an NFT:

```sh
./build/mantrachaind tx fractionnft rent <class_id> <nft_id> <rental_period> --from=acc0 --keyring-backend test --home=~/.mantrasinglenodetest --fees=500000uom --gas=auto
```

- class_id: Category of the NFT.
- nft_id: ID of the NFT to be rented.
- rental_period: Duration of the rental period.

### Querying Balances

To check the balance of tokens:

```sh
./build/mantrachaind query bank balance <address> <denom>
```

- denom: The denomination of the NFT tokens <b> fractionnft-<class_id>-<nft_id><b>.

## Workflow

1. TokenizeNFT

   - Inputs: Takes the NFT classId, nftId, tokenSupply, and timeoutHeight as inputs.

   - Process

     - Initial Check: Checks if the NFT is already tokenized.

     - Locking: If not tokenized, the NFT is locked in the module address.

     - Minting Tokens: The requested token supply is minted with the denomination format (modulename-classId-nftId) and sent to the owner’s address.

2. RemintNFT

   - Trigger:

     - EndBlocker: Triggered by the EndBlocker function based on the timeoutHeight.
     - Manual Call: Can also be called by the NFT owner if they possess the entire token supply in their wallet.

   - Process:

   - Burning Tokens: The module burns the entire token supply corresponding to the NFT denomination.

   - Releasing NFT: The NFT is then released back to the owner.

## Edge Cases

- IBC Transfer Blacklist: Transactions like IBC transfers are blacklisted to prevent tokens from being locked permanently.

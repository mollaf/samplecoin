#!/bin/bash

sampled init molla --chain-id namechain

samplecli config chain-id namechain
samplecli config output json
samplecli config indent true
samplecli config trust-node true

samplecli config keyring-backend test
samplecli keys add jack
samplecli keys add alice

sampled add-genesis-account $(samplecli keys show jack -a) 1000nametoken,100000000stake
sampled add-genesis-account $(samplecli keys show alice -a) 1000nametoken,100000000stake

sampled gentx --name jack --keyring-backend test

sampled collect-gentxs
sampled validate-genesis

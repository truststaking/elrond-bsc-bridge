# Bridge Roles

## Admin

Deploys and has ownership over all contracts of the bridge.
Has access to:

- add & remove relayers.
- change the quorum size
- add and remove tokens from whitelist.
- change the batch size.

## Relayer

Relayers are defined in the Bridge contract as whitelisted addresses that can sign and execute batches of transactions.

It's not necessary for the admin to also be a relayer.

## Depositor

Depositors are users who need to have their tokens bridged. They interact with the bridge by `deposit`ing into the `ERC20Safe` contract.
A depositor should be a user on both networks (Ethereum and Elrond).

# Contracts

## ERC20Safe

This is where all the ERC20 deposits are sent to.
The `deposit` function is the only toucphoint that the depositor has with the bridge.
All the other funtions are either called by the Admin or by the Bridge contract.

## Bridge

This is the contract that the relayers use to get and execute batches of transactions to and from Elrond network.

# Deploy

Prerequisites:

- Determine the network where you want it deployed
- Admin wallet is the first account in the HD wallet configured in `hardhat.config.js` 
- Determine and get the public addresses of all the relayers
- Determine the quorum size (should be 2/3 from the number of relayers)

## Deploying the contracts

Before being able to run the deploy, the correct `network` and admin wallet must be setup in `hardhat.config.js`

Deploy command:

`npx hardhat deploy --relayer-addresses '["relayer1Address", "relayer2Address","relayer3Address"]' --quorum 3 --network <networkName>`  (e.g. network = `rinkeby` as it is setup in `hardhat.config.js`)

Relayers can be changed over time, but the initial list is the one passed in the constructor of the bridge and saves gas during setup.

The script will output the addresses of the deployed contracts and will also create a `setup.config.json` file containing these scripts.
e.g.

```
{
  "erc20Safe":"0xbC24fB7a7a646d31C50835a3bD4C0BfE32a7AB6a",
  "bridge":"0xdA8B7a4AF87091952f5a80Ba3a048d9fCaBbD2F9"
}
```

## Whitelisting tokens

Before a depositor can go in and call the `deposit` method with an ERC20 token, that ERC20 contract must be whitelisted.
For this, there is a hardhat task: `add-to-whitelist`
Run it with: `npx hardhat add-to-whitelist --address <tokenAddress>`

# For testing in testnets 

## Creating and whitelisting test tokens
In order to have some dummy ERC20 tokens, you can run the command:
`npx hardhat deploy-test-tokens --network rinkeby`
This will create three dummy ERC20 tokens: dUSDC, dDAI, dEGLD and will whitelist them in the contracts deployed with the deploy task.

## Minting test tokens
You can mint test tokens in the dummy contracts by running the command:
`npx hardhat mint-test-tokens --recipient-address <recipientAddress>`
This will mint and send 200 tokens in each of the dummy contracts created before.

# Emergency shutdown

Remove tokens from whitelist (so depositors cannot deposit anymore)
Increase quorum to be higher than the number of relayers (so relayers cannot execute transfers anymore)

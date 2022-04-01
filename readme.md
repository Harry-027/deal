# deal
**deal** is a blockchain built using **Cosmos SDK** and **Tendermint**, specifically to handle deals between online stores and vendors.
It handles deals, contracts and fund related activity to decentralize the business model of online stores which prioritize
the fair distribution of payments and non-violation of agreements.

**Starport** has been used to here to create the blockchain.

### Install

* Install golang and starport.
* Clone the repository.
* Execute the command `starport chain build` to get the node's binary.
* Node's binary `deald` can be used to execute transactions and queries.

### Configure
blockchain in development mode can be configured with `config.yml`.

## Running locally
```
starport chain serve
```
`serve` command installs dependencies, builds, initializes, and starts blockchain in development mode.

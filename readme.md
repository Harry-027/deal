# Deal
**deal** is an application specific blockchain built using **Cosmos SDK** and **Tendermint**, specifically to handle deals between online stores and vendors.
It handles deals, contracts and fund related activity to decentralize the business model of online stores which prioritize
the fair distribution of payments and non-violation of agreements.

### Install

* Install golang and starport.
* Clone the repository.
* Execute the command `starport chain build` to get the node's binary.
* Node's binary `deald` can be used to execute transactions and queries.

### Configure
Blockchain in development mode can be configured via `config.yml`.

## Running locally
```
starport chain serve
```
`serve` command installs dependencies, builds, initializes, and starts blockchain in development mode.

#### Note - 
Please note that the deal blockchain is not part of any production or test network. It has been developed in reference with blog post series - [Building an application specific blockchain using Cosmos SDK](https://medium.com/@harish0y2j/building-an-application-specific-blockchain-using-cosmos-sdk-part-1-1f8388902fc8)

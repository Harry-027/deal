<!--
order: 2
-->

# State

The `x/deal` module stores the `newDealList` and `newContractList` in state.
It also stores the `DealCounter` and `ContractCounter` to track the number of various deals and contracts under them.

```protobuf
message GenesisState {
  Params params = 1 [(gogoproto.nullable) = false];
  DealCounter dealCounter = 2;
  repeated NewDeal newDealList = 3 [(gogoproto.nullable) = false];
  repeated ContractCounter contractCounter = 4 [(gogoproto.nullable) = true];
  repeated NewContract newContractList = 5 [(gogoproto.nullable) = false];
}
```

All `Deals` are retrieved and stored via a prefix `KVStore` using prefix key `NewDeal/value/`.
Whereas `Contracts` are retrieved and stored via a prefix `KVStore` using prefix key `NewContract/value/{dealId}`
where dealId refers to the id of the deal which contract belongs to.
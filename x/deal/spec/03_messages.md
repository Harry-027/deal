<!--
order: 3
-->

# Messages

## MsgCreateDeal

Deal is created by an account acting as deal owner via a `MsgCreateDeal` transaction:

```protobuf
message MsgCreateDeal {
  string creator = 1; // deal owner address
  string vendor = 2; // deal vendor address
  uint64 commission = 3; // vendor commission in percent
}
```

**State modifications:**

* Given the valid vendor address and commission rate, create a `deal` between msg creator and given vendor address.
* Initialize the `contractCounter` for a newly created deal.
* Increment the `dealCounter`.

## MsgCreateContract

Contract is created by an account acting as deal owner via a `MsgCreateContract` transaction:

```protobuf
message MsgCreateContract {
  string creator = 1; // deal owner address
  string dealId = 2; // id of the deal under which contract has to be created
  string consumer = 3; // end user address
  string desc = 4; // order details
  string ownerETA = 5; // estimated time of order delivery in mins
  string expiry = 6; // contract expiry time in mins
  uint64 fees = 7; // order payment , number of coins (denom is token)
}
```

**State modifications:**

* Given the valid owner and deal id, create a `contract`.
* Increment the `contractCounter`.

## MsgCommitContract

Contract is committed by an account acting as a vendor for a given dealId via a `MsgCommitContract` transaction:

```protobuf
message MsgCommitContract {
  string creator = 1; // deal vendor
  string dealId = 2; // refers to the deal to which contract belongs to 
  string contractId = 3; // refers to the contract to be committed
  string vendorETA = 4; // estimated shipping time 
}
```

**State modifications:**

* Given valid shipping time change the contract status to `COMMITTED`.

## MsgApproveContract

Contract is approved by an account acting as a `consumer` for a given order:

```protobuf
message MsgApproveContract {
  string creator = 1; // consumer address
  string dealId = 2; // refers to the deal to which contract belongs to 
  string contractId = 3; // refers to the contract to be approved
}
```

**State modifications:**

* Given valid `dealId`, `contractId` and correct consumer address, transfer funds from consumer account to module escrow account.
* Change the contract status to `APPROVED`.

## MsgShipOrder

Order is shipped by an account acting as a `vendor` for a given order:

```protobuf
message MsgShipOrder {
  string creator = 1; // vendor account address
  string dealId = 2; // refers to the deal to which contract belongs to 
  string contractId = 3; // refers to the contract for which order has been shipped
}
```

**State modifications:**

* Given valid `dealId`, `contractId` and correct vendor address, calculate the shipping delay if any and change the contract status to `IN-DELIVERY`.

## MsgOrderDelivered

Contract is marked as completed by an account acting as a `consumer` for a given order:

```protobuf
message MsgOrderDelivered {
  string creator = 1; // consumer address
  string dealId = 2; // refers to the deal to which contract belongs to 
  string contractId = 3; // refers to the contract to be approved
}
```

**State modifications:**

* Given valid `dealId`, `contractId` and correct consumer address, calculate the delivery delay if any.
* Calculate the payment for vendor based on commission rate and subtract delay shipping charges if any.
* Calculate the payment for owner by subtracting vendor commission and delay delivery charges if any.
* Refund back the subtracted delay charges to the consumer.
* Release calculated payments for owner and vendor to their respective addresses.
* change the contract status to `COMPLETED`.

## MsgCancelOrder

Order can be cancelled by the consumer:

```protobuf
message MsgCancelOrder {
  string creator = 1; // consumer address
  string dealId = 2; // refers to the deal to which contract belongs to 
  string contractId = 3; // refers to the contract to be cancelled
}
```

**State modifications:**

* Mark the contract status as `CANCELLED` and refund back order payment to the consumer in case delay is more than 20 minutes with respect to delivery time.

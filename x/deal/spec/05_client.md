<!--
order: 5
-->

# Client

## CLI

A user can query and interact with the `deal` module using the CLI.

### Query

The `query` commands allow users to query `deal` state.

```bash
deald query deal --help
```

#### list-new-deal

The `list-new-deal` command allows users to query all the created deals.

```bash
deald query deal list-new-deal [flags]
```

Example:

```bash
deald query deal list-new-deal
```

Example Output:

```bash
newDeal:
- commission: "50"
  dealId: "1"
  owner: cosmos1ncl7r9d6q3mqasweschyeaz9u6wplqlaf8ejxz
  vendor: cosmos10xsg99rp8ysjne563u5t7kytprjvs9dfv8us2a
pagination:
  next_key: null
  total: "0"
```

#### list-new-contract

The `list-new-contract` command allows users to query all contracts for a given dealId.

```bash
deald q deal list-new-contract [dealId] [flags]
```

Example:

```bash
deald q deal list-new-contract 1
```

Example Output:

```bash
newContract:
- consumer: cosmos1xlatxzq5dq2ltpgq6yqwrw09uj3zlf3syae8lh
  contractId: "1"
  dealId: "1"
  deliveryDelay: 0
  desc: 5pizza
  expiry: 2022-04-01 17:28:00.902205 +0000 UTC
  fees: "500"
  ownerETA: 50
  shippingDelay: 0
  startTime: 2022-04-01 17:23:00.902205 +0000 UTC
  status: INITIATED
  vendorETA: 0
```

#### show-new-deal

The `show-new-deal` command allows users to query deal details for a given dealId.

```bash
deald query deal show-new-deal [dealId] [flags]
```

Example:

```bash
deald query deal show-new-deal 1
```

Example Output:

```bash
newDeal:
  commission: "50"
  dealId: "1"
  owner: cosmos1ncl7r9d6q3mqasweschyeaz9u6wplqlaf8ejxz
  vendor: cosmos10xsg99rp8ysjne563u5t7kytprjvs9dfv8us2a
```

#### show-new-contract

The `show-new-contract` command allows users to query contract details for given dealId and contractId.

```bash
deald query deal show-new-contract [dealId] [contractId] [flags]
```

Example:

```bash
deald query deal show-new-contract 1 2
```

Example Output:

```bash
newContract:
  consumer: cosmos1xlatxzq5dq2ltpgq6yqwrw09uj3zlf3syae8lh
  contractId: "2"
  dealId: "1"
  deliveryDelay: 0
  desc: 2Burger
  expiry: 2022-04-01 17:28:18.580502 +0000 UTC
  fees: "800"
  ownerETA: 50
  shippingDelay: 0
  startTime: 2022-04-01 17:23:18.580502 +0000 UTC
  status: INITIATED
  vendorETA: 0
```
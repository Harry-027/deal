<!--
order: 4
-->

# Events

The deal module emits the following events:

## Handlers

### MsgCreateDeal

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| message | module        | deal            |
| message | IdValue       | {dealId}        |
| message | owner         | {owner}         |
| message | vendor        | {vendor}        |


### MsgCreateContract

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| message | module        | deal            |
| message | action        | INITIATED       |
| message | IdValue       | {contractId}    |
| message | StartTime     | {StartTime}     |


### MsgCommitContract

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| message | module        | deal            |
| message | action        | COMMITTED       |
| message | IdValue       | {contractId}    |
| message | VendorETA     | {vendorETA}     |


### MsgApproveContract

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| message | module        | deal            |
| message | action        | APPROVED        |
| message | IdValue       | {contractId}    |


### MsgShipOrder

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| message | module        | deal            |
| message | action        | IN-DELIVERY     |
| message | IdValue       | {contractId}    |


### MsgOrderDelivered

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| message | module        | deal            |
| message | action        | COMPLETED       |
| message | IdValue       | {contractId}    |
| message | Consumer      | {consumer}      |
| message | Owner         | {owner}         |
| message | Vendor        | {vendor}        |


### MsgCancelOrder

| Type    | Attribute Key | Attribute Value |
|---------|---------------|-----------------|
| message | module        | deal            |
| message | action        | CANCELLED       |
| message | IdValue       | {contractId}    |
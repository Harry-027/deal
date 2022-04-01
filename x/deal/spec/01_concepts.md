<!--
order: 1
-->

# Concepts

## UseCase

Usually big online ecommerce stores work together with vendors to complete order delivery for a consumer in a given deadline.
An online store owner takes care of handling the online store and order delivery whereas vendor takes care of 
preparing & shipping the product (handing it over to store owner for delivery to end user). Objective here is to handle the deals 
& order's fund over blockchain so that fair amount of commission is distributed among parties as per the initial agreement. 

## Deal

Here, a deal is basically an agreement between two business parties around the commission rate on order completion.
In our use-case basically a store owner creates a deal with vendor for a certain fixed commission rate (`1% <= r% <= 100%`)
which will be applicable after each order delivery. This means after successful delivery under a given deal for a given deadline,
a fixed amount of order payment will be transferred to vendor account. 

## Contract

Here, each order from an end user is represented by a contract. Based on online cart and user details, an online store owner initiates
a contract on blockchain under a deal, which includes certain details such as end user wallet address, order details 
under contract description,start time, time to be taken by owner for delivery and contract expiry time. In order to execute the contract successfully, 
vendor needs to commit the contract before its expiry. Note that, vendor needs to input shipping time for committing the contract.
Contract can be committed only if shipping time is less than half of delivery time.

## Contract Approval

Once the contract has been committed by vendor, end user is asked for payment so that the order can be processed. The payment
from end user is basically held in an escrow account which is basically the deal module account. The contract is marked as approved
after successful payment from end user which also triggers a custom event to inform vendor about contract approval.

## Contract Shipping

Once the contract has been approved, vendor can start processing the order and make it ready for consumer.
While handing over product to store owner, vendor can initiate a tx on blockchain to change the contract status to `IN-DELIVERY`.
This marks the vendor role here as completed. With this tx, shipping delay also gets recorded if any with respect to committed shipping time.

## Order Delivered

Once the order has reached to end user door step, he needs to initiate a tx for marking the order as completed. This tx will 
calculate if there is any delivery delay and will release funds from escrow account to respective parties based on calculated delay charges (shipping/delivery)
if any and commission rate. Note that in case of shipping or delivery delay, consumer will get back an amount of refund (delay charges)
proportional to delay.Delay charges will be deducted from respective parties commission (from vendor commission in case of shipping delay or
from store owner commission in case of delivery delay).

## Order Cancelled

Consumer can mark the order as cancelled if delay with respect to delivery time is more than or equal to 20 minutes. Note that
if the order cancellation is successful, user gets back the complete refund (funds from escrow account are transferred back to user account)

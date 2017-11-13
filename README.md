# healchain

## Problem: 
Basic health information of a person such as Blood group, heart/diabetic problem info, order donor status are quickly needed when a person is fainted / met in accident. Delay in finding these info may cost a life. In emergency need of initial funds is crucial. 

## Solution:
Using hyperledger blockchain, We store a person's both public health information 
& private medical history. We allow hospitals/labs to issue blood donor rebates as ERC20 based HEAL Tokens. And they will accept HEAL tokens as mode of payment.

## HyperLedger
### QR Code based : Public Health Information

```
Stored in: HyperLedger
UI : Toshi SOS + Fabric
```

* Blood Group      : B+
* Diabetic Patient : Yes / No
* Heart Patient    : Yes / No
* Allergic To      : [             ]
* Organ Donor      : Yes / No


### Private Health Information
```
Stored in: HyperLedger
UI : Fabric
```

* Scanned JPG/PDF files
* Generalized report values

Private Keys for these private health information/ records will be available only to respective patient & report adding hospital/lab.

## Blood Donor Rebates as ERC20 Tokens

### Government transfers HEAL Tokens to Hospitals every year/month in advanced
Government has particular scale of rebates to blood donors. Govt organization will send HEAL tokens to all hospitals every month in advanced. 

### Hospital Pays Donor Rebates in HEAL Tokens
Upon receiving blood donations, As per government norms 1 unit blood's rebate value will be calculated in HEAL Tokens and paid to donor's ETH address

### Donor/Patient Pays in HEAL Token
During health checkups, emergency payments Patient can pay in HEAL Token

## Toshi 
### SOS
```
* Displays QR code to access public health information
* Sends emergency alert to configured friends & family with current location
* Sends emergency alert to near by ambulance service / hospital
```
### MyHEALChain
```
Displays Healchain info / wallet
```
### Blood Request
```
Receives blood request from user & broadcasts to all other users in a location/circle
```
### Pay Hospital/Lab
```
When Toshi supports ERC20 tokens, User can send HEAL Tokens to friends, family, hospitals & labs.
```


### Tech Stack

Payments/notifications: Toshi
Frontend: Vue
BlockChain: Chaincode (Hyperledger Fabric), solidity (Ethereum)
Token: ERC20

Misc: 
Docker for app deployment
Node API for Toshi & BlockChain interactions
Go.



# cosmos-ibc-for-dummies
This is an on-hand guide to cosmos ibc for noobs

This guide focuses only on modules/apps section of IBC to show how to build a simple IBC apps

# Table of Contents
- [Basic of IBC](https://docs.cosmos.network/master/ibc/overview.html)
- [Interchain Overview](#interchain-overview)
- [Code Guide](#code-guide)
    - [Client](#client)
    - [Controller](#controller)
    - [Host](#host)
    - [Types](#types)

## Interchain Overview
The link below encompasses a great overview on interchain although the title doesn't sound like one. So patiently read it.

[Interchain Overview](https://github.com/cosmos/ibc/tree/master/spec/app/ics-027-interchain-accounts)

## Code Guide
For this part, you have to look at the code that I have prepared here in folder "modules/apps/template_module". It is a small subset of [full ibc code](https://github.com/cosmos/ibc-go).

There are four folder corresponding to four parts: 
1. client: how client interact with our IBC app?
2. controller: all stuffs related to controller chain (as specified in [Controller Chain](https://github.com/cosmos/ibc/tree/master/spec/app/ics-027-interchain-accounts#definitions))
3. host: all stuffs related to host chain (as specified in [Host Chain](https://github.com/cosmos/ibc/tree/master/spec/app/ics-027-interchain-accounts#definitions))
4. types: data structure and structure manipulation method

### Client

### Controller

### Host

### Types
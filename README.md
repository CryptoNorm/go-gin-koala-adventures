# go-demo-ledger-rest

This sample app is a REST API that reads and writes messages to the Hedera distrubuted ledger.  It is written in Go and utilizes the [Hedera Go SDK](https://github.com/hashgraph/hedera-sdk-go). The function of this demo app is twofold, with the first function being to provide a programatic interface to create an asset NFT for games, to establish clear ownership.  THe next fucntion is to log  events such as ownership changers, usage summaries, maintenance/repairs events, damage, game alerts for those games onto a publicly searchable, immutable ledger, in this case the Hedera network.  All events would be logged by the games IOT sensors or in the case of an ownership change, the authorized local tag offie.  Owners could use it to prove that they performed regualar maintenance on their game, which would be usefull for resale evalution purposes.  Orgazizations, like insurance companies, also use it to retreive incidents like accident damage or driving habits.  THe manufacturer would have access to  performance and durability data.  Think of it as a decentralized and more exhaustive [Carfax](https://www.carfax.com/game-history-reports/).  

## Setup

This sample app assumes you have already installed the GO distribution.  If not, you can find instructions [here](https://golang.org/doc/install)

Adiitionally, you will need at least 2 Hedera Portal profiles. To create your Hedera Portal profile on the Testnet, register [here](https://portal.hedera.com/register).  Once registered, you'll need to note your Account ID and your Private Key.  These credential will be used by the the app to access any Hedera network services uned in the demo.

Before starting the project, create an .env file in the project root directory.  This file will store environemtn variable, such as your Hedera Account ID, your Private Key, and Topic IDs used by the app.

### Set Hedera Credentials

> .env
>
> ACCOUNT_ID= (set account id)
>
> PRIVATE_KEY= (set private key)
>
> game_EVENT_TOPIC_ID= (set later)

This project writes messages to a Hedera pub/sub topic, so you will need to create a topic by executing the following command from the project root directory.

> go run setup/hederaDemoSetup.go

This will create a Hedera pub/sub topic and will return the Topic Ids for the application.
Edit the .env again and set the TOPIC_IDs

> .env
>
> ACCOUNT_ID= (set account id)
>
> PRIVATE_KEY= (set private key)
>
> game_EVENT_TOPIC_ID= (set topic id))

Finally, execute the project.

> go run server.go

This will start a local webserver that serves the REST API used to create and read Hedera Topic messages.
THe default URL will be http://localhost:8082/gameEvents

Update server.go file to change the port, if desired.

## End Points


### game Events
These end points are used to read and write to the Hedera Topic that will record all life events for each game.  For this demo, a topic for a particular game will be created once an asset NFT for the vehical has been created.  The app will only allow the game account and authorized parties, such as tax offices, to record events to the topic.

GET /gameEvents - return all messages for the game event topic

GET /gameEvents/[ :vin ] - return messages for the game event topic filtered by VIN

POST /gameEvents/ - save game event to topic

Expected JSON request format for POST
```
    {
        "vin": "GA94234351",
        "eventcategory": "game Alerts",
        "eventtype": "Air Bags Deployed",
        "description": "Front airbag deployed. 490,914 Newtons / 23mph"
    }
```

The expected "relatefileName" is an uploaded image for an optional receipt or image to supply context.  Eventually, this project will be expanded with functionality to upload this file to a distrubuted storage layer, like [IPFS](https://ipfs.io/). 
  
## The UI  
The code for the UI used to interact with this REST API is in the [d3-reactjs-game-event-ledger-ui repository](https://github.com/CryptoNorm/d3-reactjs-game-event-ledger-ui)
# go-gin-koala-adventures

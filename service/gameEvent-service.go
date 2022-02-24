package service

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

type gameEventService interface {
	Save(model.gameEvent) model.gameEvent
	FindAll() []model.gameEvent
	FindByVin(string) []model.gameEvent
}

type gameEventService struct {
	gameEvents []model.gameEvent
}

func NewEvent() gameEventService {
	return &gameEventService{
		gameEvents: []model.gameEvent{},
	}
}

func (service *gameEventService) Save(gameEvent model.gameEvent) model.gameEvent {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("GAME_EVENT_TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	now := time.Now()
	gameEvent.CreatedAt = now

	ma, err := json.Marshal(gameEvent)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ma))

	//Create the transaction
	transaction := hedera.NewTopicMessageSubmitTransaction().
		SetTopicID(myTopicId).
		SetMessage([]byte(string(ma)))

	//Sign with the client operator private key and submit the transaction to a Hedera network
	txResponse, err := transaction.Execute(client)
	if err != nil {
		panic(err)
	}

	//Request the receipt of the transaction
	transactionReceipt, err := txResponse.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	//Get the transaction consensus status
	transactionStatus := transactionReceipt.Status

	fmt.Printf("The transaction consensus status is %v\n", transactionStatus)
	//v2.0.0

	return gameEvent
}

func (service *gameEventService) FindByVin(searchVin string) []model.gameEvent {
	var client = GetHederaClient()

	myTopicId, err := hedera.TopicIDFromString(os.Getenv("GAME_EVENT_TOPIC_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("The topic ID = %v\n", myTopicId)

	var results []model.gameEvent

	sub, err := hedera.NewTopicMessageQuery().
		SetTopicID(myTopicId).
		SetStartTime(time.Unix(0, 0)).
		Subscribe(client, func(message hedera.TopicMessage) {
			var ma model.gameEvent
			err := json.Unmarshal(message.Contents, &ma)
			if err != nil {
				println(err.Error(), ": error Unmarshalling")
			}
			fmt.Println(ma.Player, "-", ma.GameLevel, "-", ma.Score, "-", ma.CreatedAt)
			if (ma.Player == searchPlayer) || (searchPlayer == "") {
				results = append(results, ma)
			}
		})

	if err != nil {
		println(err.Error(), ": error subscribing to the topic")
		return results
	}

	time.Sleep(3 * time.Second)
	sub.Unsubscribe()

	if err != nil {
		panic(err)
	}

	return results
}

func (service *gameEventService) FindAll() []model.gameEvent {
	return service.FindByVin("")
}

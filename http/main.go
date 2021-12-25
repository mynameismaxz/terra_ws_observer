package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"terra/terra_observer/model"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)

	u := url.URL{Scheme: "wss", Host: "observer.terra.dev"}
	fmt.Printf("Connected to %s...\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	sub := &model.Subscribe{
		Subscribe: "new_txs",
		ChainId:   "columbus-5",
	}

	payload, err := json.Marshal(sub)
	if err != nil {
		panic(err)
	}
	msg := string(payload)
	err = c.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		observerResponse := model.ObserverResponse{}
		for {
			err := c.ReadJSON(&observerResponse)
			if err != nil {
				panic(err)
			}
			fmt.Println("Read data on block: " + observerResponse.Data.Block.Header.Height)
			for i, _ := range observerResponse.Data.Txs {
				for j, _ := range observerResponse.Data.Txs[i].Body.Messages {
					fmt.Println(observerResponse.Data.Txs[i].RawLog)
					fmt.Println("MsgType: " + observerResponse.Data.Txs[i].Events[j].Type)

					if observerResponse.Data.Txs[i].Body.Messages[j].Contract == "terra19qx5xe6q9ll4w0890ux7lv2p4mf3csd4qvt3ex" {
						fmt.Println("Found execute on terraswap router!")
						fmt.Println("Offer assets: " + observerResponse.Data.Txs[i].Body.Messages[j].ExecuteMsg.Swap.OfferAsset.Info.NativeToken.Denom)
						// if observerResponse.Data.Txs[i].Body.Messages[j].Type == "/terra.wasm.v1beta1.MsgExecuteContract" {
						// 	// execute contract only!
						// 	fmt.Println("Type: " + observerResponse.Data.Txs[i].Body.Messages[j].Type)
						// 	fmt.Println("From: " + observerResponse.Data.Txs[i].Body.Messages[j].Sender)
						// 	fmt.Println("Amount: " + observerResponse.Data.Txs[i].Body.Messages[j].ExecuteMsg.Swap.OfferAsset.Amount)
						// 	fmt.Println("Denom: " + observerResponse.Data.Txs[i].Body.Messages[j].ExecuteMsg.Swap.OfferAsset.Info.NativeToken.Denom)
						// }
					}
				}
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			fmt.Println("Interrupt mode.")
			return
		}
	}
}

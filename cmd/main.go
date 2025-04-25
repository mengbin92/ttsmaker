package main

import (
	"context"
	"fmt"

	"github.com/mengbin92/ttsmaker"
)

func main() {
	client := ttsmaker.NewClient("ttsmaker_demo_token")

	listResp, err := client.GetVoiceList(context.Background())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(listResp)

	orderRequest := &ttsmaker.OrderRequest{
		APIKey:                 "ttsmaker_demo_token",
		Text:                   "hello, world!",
		VoiceID:                778,
		AudioFormat:            "mp3",
		AudioSpeed:             1.0,
		AudioVolume:            0,
		TextParagraphPauseTime: 0,
	}

	orderResp, err := client.CreateOrder(context.Background(), orderRequest)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(orderResp)

	statusResp, err := client.GetTokenStatus(context.Background())
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(statusResp)
}

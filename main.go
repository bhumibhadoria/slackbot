package main

import (
	"fmt"
	"os"
	"github.com/shomali11/slacker"
	"context"
	"log"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-7207885706675-7209560660803-FBCyY5xzcLSBygvX1OXjpm9h")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A0765KCE5AN-7209659692290-b605444f5d3b4697d82a04e7c2b1fdb90bf25957f99d277a2cf400769a96b068")


	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
		response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}
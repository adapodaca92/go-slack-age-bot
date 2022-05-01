package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()	
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3461278038869-3464191399699-GNeoKra2503TuQvGl3sObgF5")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03DN59GC02-3464182172530-9cc2c13a165c606b8d559e7c59b18ffdd9868620df3b8c40d6e3e111daff5669")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My YOB is <year>", &slacker.CommandDefinition {
		Description: "YOB calculator",
		Example: "My YOB is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			YOB, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022 - YOB
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
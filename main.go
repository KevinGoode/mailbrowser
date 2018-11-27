package main

import (
	"fmt"

	"./utils"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(utils.NewArgs)
	container.Provide(utils.NewArgumentReader)
	container.Provide(utils.NewAuthenticatedGmailClient)
	container.Provide(utils.NewMailboxRequestor)
	container.Provide(utils.NewCommandExecutor)
	return container
}

// Main executable to perform operations on a user mailbox
func main() {
	//This code uses dig dependency injector for more details see
	//https://godoc.org/go.uber.org/dig
	//https://blog.drewolson.org/dependency-injection-in-go

	container := buildContainer()

	err := container.Invoke(func(executor utils.CommandExecutorAPI) {
		executor.Run()
	})
	if err != nil {
		fmt.Println("General error executing command")
	}
}

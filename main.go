package main

import (
	"context"

	"github.com/kaankoken/helper/cmd"
)

func main() {
	app := cmd.MainApp()

	app.Run()

	defer app.Stop(context.Background())
}

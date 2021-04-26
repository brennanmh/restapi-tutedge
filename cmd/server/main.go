package main

import "fmt"

// App - the strcut which contain things like pointers
// to database connections.
type App struct{}

// Run - sets up our applications
func (app *App) Run() error {
	fmt.Println("Setting up our App")
	return nil
}

func main() {

	fmt.Println("Go REST API Course")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error running rest service")
		fmt.Println(err)
	}
}

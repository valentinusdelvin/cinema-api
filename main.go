package main

import (
	"Cinema_API/config"
	"Cinema_API/internal/repository"
	"Cinema_API/internal/rest"
	"Cinema_API/internal/usecase"
	"Cinema_API/pkg/bcrypt"
	"Cinema_API/pkg/jwt"
	"Cinema_API/pkg/middleware"
	"fmt"
)

func init() {
	config.LoadEnvVariables()
}

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	bcrypt := bcrypt.BcryptInit()
	jwt := jwt.JWTInit()
	db := config.ConnectToDB()
	config.Migrate(db)

	repository := repository.RepositoryInit(db)
	usecase := usecase.UsecaseInit(usecase.InitializersParam{
		Repository: *repository,
		JWT:        jwt,
		Bcrypt:     bcrypt,
	})
	middleware := middleware.Init(*usecase)

	rest := rest.RestInit(usecase, middleware)
	rest.FinalEndpoint()
	rest.Run()

	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}
}

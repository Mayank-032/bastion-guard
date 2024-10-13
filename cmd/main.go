package cmd

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Mayank-032/bastion-guard/infrastructure"
	"github.com/Mayank-032/bastion-guard/internal/repository"
	"github.com/Mayank-032/bastion-guard/internal/usecase"
)

// usecases
var user usecase.User

// Initialise Infrastructure required run application
func initializeInfra() {
	err := infrastructure.InitConfig()
	if err != nil {
		fmt.Println("unable to initialise config")
		os.Exit(1)
		return
	}
	fmt.Println("successfully loaded configs")

	err = infrastructure.InitDB("", "", "")
	if err != nil {
		fmt.Println("unable to initialise DB")
		os.Exit(1)
		return
	}
	fmt.Println("successfully loaded data store")

	var readRepository = repository.NewReadUserRepository(infrastructure.DB)
	var upsertRepository = repository.NewUpsertUserRepository(infrastructure.DB)
	var deleteRepository = repository.NewDeleteUserRepository(infrastructure.DB)

	user = usecase.NewLoginUsecase(readRepository, upsertRepository, deleteRepository)
}

func Execute() {
	initializeInfra()

	// initialise reader instance, which will be used to read the input in terminal
	var reader = bufio.NewReader(os.Stdin)

	// delimiter means the end character till which the input will be read and processed as a single entity
	var delim = '\n'

	for {
		inputCommandStr, err := reader.ReadString(byte(delim))
		if err != nil {
			fmt.Println("invalid input format")
			os.Exit(1)
			return
		}
		inputCommandStr = strings.Trim(inputCommandStr, " ")

		var inputCommandArr = strings.Fields(inputCommandStr)

		var fs = flag.NewFlagSet("command", flag.ExitOnError)
		var execCommand = fs.String("exec", "bastion-guard", "Entry point for processing requests coming to bastion guard")
		var operationCommand = fs.String("operation", "", "Operation to be performed by the application")
		var authCredName = fs.String("name", "", "Username of the user")
		var authCredPwd = fs.String("pwd", "", "Password of the user")
		var authCredNewPwd = fs.String("npwd", "", "New Password of the user")

		if err := fs.Parse(inputCommandArr); err != nil {
			fmt.Printf("Error parsing flags: %v\n", err.Error())
			os.Exit(1)
			return
		}

		// if no executable command is provided or if provided but is not the default command, return it as error
		if execCommand == nil || *execCommand != "bastion-guard" {
			err = errors.New("invalid executable command")
		} else if operationCommand == nil || !contains(operations_arr, *operationCommand) {
			err = errors.New("invalid operation command")
		} else if authCredName == nil {
			err = errors.New("invalid auth cred. please provide username")
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		var ctx = context.Background()

		isUserCreated, err := user.IsCreated(ctx, *authCredName, *authCredPwd)
		if err != nil {
			if err.Error() == "invalid user" {
				fmt.Println("invalid password for existing user. please provide valid auth credentials")
			} else {
				fmt.Println("oops, some error occurred. please try again or contact support")
			}
			return
		}

		switch *operationCommand {
		case LOGIN:
			if !isUserCreated {
				err = user.Create(context.Background(), *authCredName, *authCredPwd)
				if err != nil {
					fmt.Println("oops, some error occurred. please try again or contact support")
					return
				}
			}
			fmt.Println("success...")
			return
		case DELETE:
			if !isUserCreated {
				fmt.Println("cannot perform this operation. no such user exists")
				return
			}

			err = user.MarkInactive(ctx, *authCredName)
			if err != nil {
				fmt.Println("oops, some error occurred. please try again or contact support")
				return
			}
			fmt.Println("success...")
			return
		case UPDATE_PASSWORD:
			authCredNewPwd = fs.String("npwd", "", "New Password of the user")

			err := user.UpdatePassword(ctx, *authCredNewPwd, *authCredPwd, *authCredName)
			if err != nil {
				fmt.Println("oops, some error occurred. please try again or contact support")
				return
			}
			fmt.Println("success...")
			return
		}
	}
}

func contains(stringArr []string, str string) bool {
	for _, currEle := range stringArr {
		if currEle == str {
			return true
		}
	}
	return false
}

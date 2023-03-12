package bye_test

import (
	"errors"
	"fmt"
	"time"

	"go.szostok.io/go-bye"
)

func ExampleParentService() {
	shutdown := bye.NewParentService(bye.WithTimeout(30 * time.Second))

	shutdown.Register(bye.Func(func() {
		fmt.Println("Closing non error function call")
	}))

	shutdown.Register(bye.ErrFunc(func() error {
		fmt.Println("Closing error function call")
		return errors.New("I don't want to quit!")
	}))

	shutdown.Register(&exampleService{})

	fmt.Println("Shutting down the application...")
	err := shutdown.Shutdown()
	fmt.Println(err)

	// output:
	// Shutting down the application...
	// Closing example Shutdownable Service
	// Closing error function call
	// Closing non error function call
	// 1 error occurred:
	// 	* I don't want to quit!
	//
}

var _ bye.ShutdownableService = (*exampleService)(nil)

type exampleService struct{}

func (e exampleService) Shutdown() error {
	fmt.Println("Closing example Shutdownable Service")
	return nil
}

package cmd

import (
	"fmt"

	_ "github.com/aws/aws-sdk-go/aws/awserr"
)

func Run(v string) {
	fmt.Println("running import-go-modules in ", v)
}

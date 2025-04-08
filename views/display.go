package views

import (
	"fmt"
	"ghozz/models"
)

type OutputHandler interface {
	PrintResult(result models.Result)
}

func PrintResult(r models.Result) {
	fmt.Println(r.String())
}

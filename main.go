package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	cost  int
	match string
)

func run(cmd *cobra.Command, args []string) {
	password := args[0]
	if match != "" {
		err := Compare(match, password)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("ok")
		}
		return
	}

	hash, err := Encrypt(password, cost)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(hash)
}

func newCommand() *cobra.Command {
	var c = &cobra.Command{
		Use:   "bcrypt",
		Short: "Encrypt and verify using bcrypt",
		Args:  cobra.MinimumNArgs(1),
		Run:   run,
	}
	c.PersistentFlags().IntVarP(&cost, "cost", "c", 10, "If the cost given is less than 4, the cost will be set to 10")
	c.PersistentFlags().StringVarP(&match, "match", "m", "", "check the password is match hash")
	return c
}

func main() {
	c := newCommand()
	err := c.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}

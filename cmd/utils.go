package cmd

import (
	"fmt"
	"log"

	"github.com/crossphoton/drive-encrypt/src"
	"github.com/spf13/cobra"
)

func getPassword(cmd *cobra.Command) (password string) {
	password, _ = cmd.Flags().GetString("password")
	if password == "" {
		log.Fatal("No password specified")
	}
	return
}

func printFiles(files []src.File) {
	fmt.Printf("%v\t%v\t\t\t\t%v\t%v\n", "Index", "Name", "Path", "Size")
	fmt.Println("-------------------------------------------------------------------------------------------")
	for ind, file := range files {
		fmt.Printf("%v\t%v\t\t\t\t%v\t%v\n", ind+1, file.Name, file.Path, file.Size)
	}
}

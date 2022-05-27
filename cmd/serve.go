/*
Copyright © 2022 Aditya Agrawal adiag1200@gmail.com

*/
package cmd

import (
	"log"
	"net/http"

	srcHttp "github.com/crossphoton/drive-encrypt/src/http"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP file server",
	Long:  `Start the HTTP file server`,
	Run: func(cmd *cobra.Command, args []string) {
		if !HTTP_SERVER_READY {
			log.Fatal("server not ready yet")
		}
		mux := srcHttp.GetRouter()
		log.Printf("starting http server at %v", SERVER_ADDRESS)
		http.ListenAndServe(SERVER_ADDRESS, mux)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

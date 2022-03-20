package console

import (
	"fmt"
	"log"
	"mail-service/internal/config"
	"mail-service/internal/db"
	"mail-service/internal/router"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "This command used to start the server",
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	if err := db.PgConnect(); err != nil {
		log.Panic(err.Error())
		os.Exit(1)
	}

	r := router.Router()
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", config.ServerPort()),
		Handler: r,
	}

	log.Print(fmt.Sprintf("Server is listening on port: %d", config.ServerPort()))

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf("Server failed to start: %s", err.Error()))
	}
}

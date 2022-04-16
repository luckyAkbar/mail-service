package console

import (
	"fmt"
	"mail-service/internal/config"
	"mail-service/internal/db"
	"mail-service/internal/delivery/httpsvc"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
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
		logrus.Error(err.Error())
		os.Exit(1)
	}

	httpServer := echo.New()
	httpServer.Use(middleware.Logger())

	group := httpServer.Group("")

	httpsvc.RouteService(group)

	httpServer.Start(fmt.Sprintf(":%s", config.ServerPort()))
}

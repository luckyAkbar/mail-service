package console

import (
	"fmt"
	"mail-service/internal/config"
	"mail-service/internal/db"
	"mail-service/internal/delivery/httpsvc"
	"mail-service/internal/helper"
	"mail-service/internal/repository"
	"mail-service/internal/usecase"
	"mail-service/internal/worker"

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
	db.InitializePostgresConn()
	setupLogger()

	sqlDB, err := db.PostgresDB.DB()
	if err != nil {
		logrus.Fatal("unable to start server:", err.Error())
	}

	defer helper.WrapCloser(sqlDB.Close)

	cryptor := helper.CreateCryptor()

	mailRepository := repository.NewMailRepository(db.PostgresDB)
	mailUsecase := usecase.NewMailUsecase(mailRepository, cryptor)

	sibClient := helper.NewSIBHelper(config.SIBClient())
	mailWorker := worker.NewMailWorker(mailRepository, sibClient)

	go mailWorker.SpawnWorker()

	server := echo.New()
	server.Pre(middleware.AddTrailingSlash())
	server.Use(middleware.Logger())

	RESTGroup := server.Group("rest")
	httpsvc.InitService(RESTGroup, mailUsecase)

	err = server.Start(fmt.Sprintf(":%s", config.ServerPort()))
	if err != nil {
		logrus.Fatal("unable to start server:", err.Error())
	}

	logrus.Info("server is up and running")
}

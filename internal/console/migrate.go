package console

import (
	"mail-service/internal/config"
	"mail-service/internal/db"
	"strconv"

	"github.com/kumparan/go-utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	migrate "github.com/rubenv/sql-migrate"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "initialize database table",
	Long:  "Use this command to intialize your database scheme for the first time",
	Run:   runMigrate,
}

func init() {
	migrateCmd.PersistentFlags().Int("step", 0, "maximum migration step")
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction")
	RootCmd.AddCommand(migrateCmd)
}

func runMigrate(cmd *cobra.Command, args []string) {
	direction := cmd.Flag("direction").Value.String()
	stepStr := cmd.Flag("step").Value.String()

	step, err := strconv.Atoi(stepStr)
	if err != nil {
		logrus.Fatal("invalid step value. ", err.Error())
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./db/migration",
	}

	migrate.SetTable("schema_migrations")
	db.InitializePostgresConn()

	pgdb, err := db.PostgresDB.DB()
	if err != nil {
		logrus.WithField("DatabaseDSN", config.PostgresDSN()).Fatal("failed to run migration")
	}

	var n int
	if direction == "down" {
		n, err = migrate.ExecMax(pgdb, "postgres", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(pgdb, "postgres", migrations, migrate.Up, step)
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"migrations": utils.Dump(migrations),
			"direction":  direction}).
			Fatal("Failed to migrate database: ", err)
	}

	logrus.Infof("Applied %d migrations!\n", n)
}

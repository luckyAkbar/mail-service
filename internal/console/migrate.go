package console

import (
	"log"
	"mail-service/internal/db"
	"mail-service/internal/models"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "initialize database table",
	Long:  "Use this command to intialize your database scheme for the first time",
	Run:   migrate,
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) {
	log.Print("Preparing to migrating...")

	if err := db.PgConnect(); err != nil {
		log.Panic("Failed to connect to Postgres Database")
	}

	log.Println("Connected to the database!")
	log.Println("Migration start...")

	log.Println("Migrating User model")
	db.DB.AutoMigrate(&models.User{})

	log.Println("Migrating APIKey model")
	db.DB.AutoMigrate(&models.APIKey{})

	log.Println("Migrating MailingList model")
	db.DB.AutoMigrate(&models.MailingList{})

	log.Println("Migration finished.")
}

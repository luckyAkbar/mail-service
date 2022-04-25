package console

import (
	"log"
	"mail-service/internal/db"
	"mail-service/internal/models"

	"github.com/spf13/cobra"
)

var seederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "seed the database",
	Long:  "Use this command to seed your database",
	Run:   seed,
}

func init() {
	RootCmd.AddCommand(seederCmd)
}

func seed(cmd *cobra.Command, args []string) {
	log.Println("Begin seeding database...")

	if err := db.PgConnect(); err != nil {
		log.Panic("Failed to connect to the database.")
	}

	log.Println("Seeding users table")

	freeUser := models.User{
		Email:    "free-user-noreply@mail.service.luckyakbar.tech",
		Username: "Free Email Service User @ mail.service.luckyakbar.tech",
	}

	if db.DB.Create(&freeUser).Error != nil {
		log.Panic("Failed to seed the users table")
	}

	log.Print("Users table already populated.")
}

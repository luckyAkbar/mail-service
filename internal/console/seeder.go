package console

import (
	"log"
	"mail-service/internal/config"
	"mail-service/internal/db"
	"mail-service/internal/helper"
	"mail-service/internal/models"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
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
		log.Fatal("Failed to connect to the database.")
	}

	log.Println("Seeding users table")

	email, err := helper.CreateCryptor().Encrypt(config.GetFreeMailUserEmailAddress())
	if err != nil {
		log.Fatal(err)
	}

	phoneNumber, err := helper.CreateCryptor().Encrypt("123456789010")

	pwd, err := helper.HashString(config.GetFreeMailUserPassword())
	if err != nil {
		log.Fatal(err)
	}

	freeUser := models.User{
		Model: gorm.Model{
			ID: uint(config.FREE_USER_ID),
		},
		Email:       email,
		Password:    pwd,
		PhoneNumber: phoneNumber,
		Username:    "Free Email Service User @ mail.service.luckyakbar.tech",
	}

	if err := db.DB.Create(&freeUser).Error; err != nil {
		log.Panic("Failed to seed the users table")
	}

	log.Print("Users table already populated.")
}

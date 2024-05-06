package main

import (
	QueueOptimization "QueueOptimization"
	"QueueOptimization/pkg/handler"
	"QueueOptimization/pkg/repository"
	"QueueOptimization/pkg/service"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("some error with initializiing: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		log.Fatalf("error with data base: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(QueueOptimization.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error with run server: %s", err.Error())
	}

	const desiredHour = 0 // 00:00
	const desiredWeekday = time.Monday
	go func() {
		for {
			// Get the current time
			now := time.Now()

			// Check if the current time matches the desired time
			if now.Hour() == desiredHour && now.Weekday() == desiredWeekday {
				// Execute your database update task here
				fmt.Println("Executing database update task at", now)

				// Replace the following line with your database update logic
				// updateDatabase()
			}

			// Sleep for a short duration before checking again (e.g., every minute)
			time.Sleep(time.Minute)
		}
	}()

}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}
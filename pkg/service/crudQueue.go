package service

import (
	"QueueOptimization/dtos"
	"QueueOptimization/models"
	"fmt"
	"log"
	"strconv"
	"time"
)

func calculateAge(birthDate time.Time) int {
	today := time.Now()
	age := today.Year() - birthDate.Year()
	if today.YearDay() < birthDate.YearDay() {
		age--
	}
	return age
}

func ParseAgeByIIN(iin string) (int, error) {
	if len(iin) < 6 {
		return -1, fmt.Errorf("invalid IIN: too short")
	}
	year, err := strconv.Atoi(iin[0:2])
	if err != nil {
		return -1, err
	}
	month, err := strconv.Atoi(iin[2:4])
	if err != nil {
		return -1, err
	}
	day, err := strconv.Atoi(iin[4:6])
	if err != nil {
		return -1, err
	}
	if year <= 24 {
		year += 2000
	} else {
		year += 1900
	}

	birthDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return calculateAge(birthDate), nil
}

func (s *service) CreateQueue(input dtos.CreateQueueRequest) (int, error) {
	age, err := ParseAgeByIIN(input.IIN)

	if err != nil {
		return -1, err
	}

	log.Print(age)
	predicted_time, err := s.PredictTime(input.CategoryId, age)

	if err != nil {
		return -1, err
	}

	log.Print(predicted_time)
	placeId, err := s.GetMostFreePlaceByCategoryId(input.CategoryId)

	if err != nil {
		return -1, err
	}

	queue := models.NewQueue(predicted_time, age, placeId, input.IIN)

	return s.repos.CreateQueue(queue)
}

func (s *service) GetQueueInfo(iin string) (dtos.GetQueueInfoResponse, error) {
	log.Print("service send request to response")
	log.Print(iin)
	return s.repos.GetQueueInfo(iin)
}

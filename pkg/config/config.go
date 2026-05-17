package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken  string
	CalendarAPIURL string
	CachePath      string
	UpdateHour     int
	Payday         int
	SalaryDay      int
}

func Load() (*Config, error) {
	const envFile = ".env"

	_ = godotenv.Load(envFile)

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		return nil, errors.New("TELEGRAM_TOKEN environment variable not set")
	}

	return &Config{
		TelegramToken: token,
	}, nil
}

//func Load() (*Config, error) {
//	updateHour, _ := strconv.Atoi(getEnv("UPDATE_HOUR", "6"))
//	payDay, _ := strconv.Atoi(getEnv("PAYDAY", "8"))
//	salaryDay, _ := strconv.Atoi(getEnv("SALARY_DAY", "15"))
//
//	return &Config{
//		TelegramToken:  getEnv("TELEGRAM_TOKEN", ""),
//		CalendarAPIURL: getEnv("CALENDAR_API_URL", "https://calendar.kuzyak.in/api/calendar"),
//		CachePath:      getEnv("CACHE_PATH", "data/calendar_cache.json"),
//		UpdateHour:     updateHour,
//		Payday:         payDay,
//		SalaryDay:      salaryDay,
//	}, nil
//}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

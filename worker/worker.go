package worker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"klp/models"
	pg "klp/storage/pg_storage"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Worker struct {
	sport       string
	LineStorage *pg.LineStorage
	interval    time.Duration
	stop        chan struct{}
}

func NewWorker(sport string, lineStorage *pg.LineStorage, interval time.Duration) *Worker {
	return &Worker{
		sport:       sport,
		LineStorage: lineStorage,
		interval:    interval,
		stop:        make(chan struct{}),
	}
}

func (w *Worker) Start() {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			line, err := GetLine(w.sport)
			if err != nil {
				continue
			}

			err = w.LineStorage.AddLine(line)
			if err != nil {
				continue
			}
		case <-w.stop:
			return
		}
	}
}

func (w *Worker) Stop() {
	close(w.stop)
}

func InitEnv() (*models.Env, error) {
	envFile := ".env"

	err := loadEnvFile(envFile)
	if err != nil {
		return nil, err
	}

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetInt("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")

	return &models.Env{
			DBHost:     dbHost,
			DBPort:     dbPort,
			DBUser:     dbUser,
			DBName:     dbName,
			DBPassword: dbPassword,
		},
		nil
}

func loadEnvFile(filePath string) error {
	// Check if the .env file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filePath)
	}

	// Set the Viper environment variable key delimiter to "_"
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName(".env")

	viper.SetConfigType("env")

	viper.AddConfigPath(".")

	// Read the .env file
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading .env file: %s", err)
	}

	// Automatically bind the environment variables to Viper
	viper.AutomaticEnv()

	return nil
}

type LineMap map[string]map[string]string

func GetLine(sport string) (*models.Line, error) {
	url := fmt.Sprintf("localhost:8000/api/v1/lines/%s", sport)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request failed with status: %d", resp.StatusCode)
	}

	var lm LineMap

	err = json.NewDecoder(resp.Body).Decode(&lm)
	if err != nil {
		return nil, err
	}

	cfStr, ok := lm["lines"][sport]
	if !ok {
		return nil, fmt.Errorf("line is not presented")
	}

	cfFloat, err := strconv.ParseFloat(cfStr, 64)
	if err != nil {
		return nil, err
	}

	line := &models.Line{
		Sport: sport,
		Cf:    cfFloat,
	}

	return line, nil
}

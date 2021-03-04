package local

import (
	"os"
)

func Getenv(key string) string {
	//err := godotenv.Load()
	//if err != nil {
	//	logg.Print("Error loading .env file")
	//}
	return os.Getenv(key)
}

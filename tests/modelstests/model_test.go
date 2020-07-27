package modelstests

import (
	"github.com/joho/godotenv"
	"github.com/snow-dev/super_blog/api/controllers"
	"github.com/snow-dev/super_blog/models"
	"log"
	"os"
	"testing"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
}

package MVC

import (
	_ "context"
	"fmt"
	"log"
	_ "mvc/pkg/api"
	"mvc/pkg/config"
	_ "mvc/pkg/models"
	_ "mvc/pkg/utils"
	_ "net/http"
	_ "os"
	_ "os/signal"
	_ "syscall"
	"testing"
	_ "time"
)

func TestEnv(t *testing.T) {
	config.LoadEnvs()
	fmt.Println(config.GetConnectionString())
	fmt.Println(config.ServerPort)
	fmt.Println(config.JWTSecret)
	if config.GetConnectionString() != "root:pass@tcp(localhost:3306)/ChefDB?parseTime=true" {
		log.Panic("Incorrect connection string")
	}
	if config.ServerPort != "8090" {
		log.Panic("Incorrect server port")
	}
	if config.JWTSecret != "qawghui[poy8756e4srdfcghyu7t65e4wsrdxcghyutr5esdxfcgyuitrde67883" {
		log.Panic("Incorrect JWT secret")
	}

}

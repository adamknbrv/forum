package app

import (
	"github.com/adamknbrv/forum/internal/repository"
	"github.com/adamknbrv/forum/internal/transport/http"
	"github.com/adamknbrv/forum/internal/transport/http/handlers"
	"github.com/adamknbrv/forum/internal/usecase"
	database "github.com/adamknbrv/forum/pkg/db"
	"github.com/adamknbrv/forum/pkg/log"
	"github.com/adamknbrv/forum/pkg/ws"
	"go.uber.org/zap"
)

func Run() {
	db, err := database.SQLiteConnection()
	if err != nil {
		panic(err)
	}

	disRep := repository.NewDiscRep(db)
	messRep := repository.NewMesRep(db)

	disCase := usecase.NewDisUseCase(disRep)
	messCase := usecase.NewMessageUseCase(messRep)

	router := http.SetupRouters(handlers.NewHandler(disCase, messCase), ws.NewHub(messCase, disCase))

	if err = router.Run(":8082"); err != nil {
		log.Send.Fatal("Ошибка создание подключения",
			zap.Error(err))
		return
	}
}

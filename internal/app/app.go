package app

import (
	"github.com/Defenestrationq/forum-api/internal/repository"
	"github.com/Defenestrationq/forum-api/internal/transport/http"
	"github.com/Defenestrationq/forum-api/internal/transport/http/handlers"
	"github.com/Defenestrationq/forum-api/internal/usecase"
	database "github.com/Defenestrationq/forum-api/pkg/db"
	"github.com/Defenestrationq/forum-api/pkg/log"
	"github.com/Defenestrationq/forum-api/pkg/ws"
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

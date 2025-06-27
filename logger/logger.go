package logger

import (
	"log"
	"os"

	"github.com/anuragdaksh7/zapmail-backend/config"
	adapter "github.com/axiomhq/axiom-go/adapters/zap"
	"github.com/axiomhq/axiom-go/axiom"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(config config.Config) {

	if config.Environment == "dev" {
		devLogger, _ := zap.NewDevelopment()
		Logger = devLogger
		return
	}

	core, err := adapter.New(
		adapter.SetDataset(config.AxiomDataset),
		adapter.SetClientOptions(
			axiom.SetToken(config.AxiomToken),
		),
	)
	if err != nil {
		log.Fatalf("Could not create axiom core %v", err)
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	Logger = zap.New(zapcore.NewTee(core, consoleCore))
}

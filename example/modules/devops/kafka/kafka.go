package kafka

import (
	"example/app"
	"github.com/spf13/cobra"
)

type kafka struct {
	app.UseCase
}

var _kafka *kafka

var _kafkaCmd = &cobra.Command{
	Use: "kafka",
}

func InitKafka(parentCmd *cobra.Command) {
	_kafka = &kafka{}
	_kafka.InitUseCase(parentCmd)

	initPartition(_kafkaCmd)

	parentCmd.AddCommand(_kafkaCmd)
}

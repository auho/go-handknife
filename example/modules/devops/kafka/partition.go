package kafka

import (
	"example/app"
	"github.com/auho/go-handknife/emergencybox/toolkit/prompt"
	"github.com/auho/go-toolkit/kafka/partitions/reassign/gen"
	"github.com/spf13/cobra"
)

type partition struct {
	app.UseCase

	topic     string
	partition int
	replica   int
}

func (p *partition) reassignGen() {
	p.PfVoid("reassign partition", func() error {
		_topic, err := prompt.NewStringWithValue("topic", p.topic, nil)
		if err != nil {
			return err
		}

		__partition, err := prompt.NewIntWithValue("partition", p.partition, nil)
		if err != nil {
			return err
		}

		_replica, err := prompt.NewIntWithValue("replica", p.replica, nil)
		if err != nil {
			return err
		}

		var s string
		s, err = gen.ReassignPartitionsToJson(
			p.App().GetEnv().BaseKafka.NetWork,
			p.App().GetEnv().BaseKafka.Address, gen.Req{
				Version:   1,
				Topic:     _topic,
				Partition: __partition,
				Replica:   _replica,
			})

		if err != nil {
			return err
		}

		p.PfBody(s)

		return nil
	})
}

var _partition *partition

var _partitionCmd = &cobra.Command{
	Use: "reassign-gen",
	Run: func(cmd *cobra.Command, args []string) {
		_partition.reassignGen()
	},
}

func initPartition(parentCmd *cobra.Command) {
	_partition = &partition{}
	_partition.InitUseCase(parentCmd)

	_partitionCmd.Flags().StringVarP(&_partition.topic, "topic", "t", "", "Topic name")
	_partitionCmd.Flags().IntVarP(&_partition.partition, "partition", "p", 0, "Partition number")
	_partitionCmd.Flags().IntVarP(&_partition.replica, "replica", "r", 0, "Replica number")

	parentCmd.AddCommand(_partitionCmd)
}

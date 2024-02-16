package devops

import (
	esDescribes "github.com/auho/go-toolkit/elasticsearch/describes"
	kafkaDescribes "github.com/auho/go-toolkit/kafka/describes"

	"example/app"
	"github.com/auho/go-handknife/emergencybox/toolkit/prompt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/spf13/cobra"
)

const (
	serviceElasticSearch = "elasticsearch"
	serviceKafka         = "kafka"
)

type describe struct {
	app.UseCase

	service string
}

func (d *describe) desc() {
	d.PfVoid("desc", func() error {
		service, err := prompt.NewSelectWithValue("input service", d.service, []string{serviceElasticSearch, serviceKafka})
		if err != nil {
			return err
		}

		switch service {
		case serviceElasticSearch:
			table, err := esDescribes.DescribeAllIndices(elasticsearch.Config{
				Addresses: app.App.GetEnv().BaseEs.Address,
			})

			if err != nil {
				return err
			}

			d.PfBody(table.Render())

		case serviceKafka:
			_t, err := kafkaDescribes.DescribeBrokers(app.App.GetEnv().BaseKafka.NetWork, app.App.GetEnv().BaseKafka.Address)
			if err != nil {
				return err
			}

			d.PfBody(_t.Render())

			_t, err = kafkaDescribes.DescribeAllTopics(app.App.GetEnv().BaseKafka.NetWork, app.App.GetEnv().BaseKafka.Address)
			if err != nil {
				return err
			}

			d.PfBody(_t.Render())
		}

		return nil
	})
}

var _describe *describe

var _describeCmd = &cobra.Command{
	Use: "describe",
	Run: func(cmd *cobra.Command, args []string) {
		_describe.desc()
	},
}

func initDescribe(parentCmd *cobra.Command) {
	_describe = &describe{}
	_describe.InitUseCase(parentCmd)

	_describeCmd.Flags().StringVar(&_describe.service, "service", "", "devops service")

	parentCmd.AddCommand(_describeCmd)
}

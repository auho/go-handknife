package devops

import (
	elasticsearch "example/modules/devops/elasticsearch/cmd"
	"example/modules/devops/kafka"
	"github.com/spf13/cobra"
)

var _devopsCmd = &cobra.Command{
	Use: "devops",
}

func Initialization(rootCmd *cobra.Command) {
	initDescribe(_devopsCmd)
	kafka.InitKafka(_devopsCmd)
	elasticsearch.InitElasticsearch(_devopsCmd)

	rootCmd.AddCommand(_devopsCmd)
}

package elasticsearch

import (
	"github.com/spf13/cobra"
)

var _elasticsearchCmd = &cobra.Command{
	Use: "elasticsearch",
}

func InitElasticsearch(parentCmd *cobra.Command) {
	initSummary(_elasticsearchCmd)

	parentCmd.AddCommand(_elasticsearchCmd)
}

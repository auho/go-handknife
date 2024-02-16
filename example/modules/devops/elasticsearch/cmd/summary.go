package elasticsearch

import (
	"os"

	"example/app"
	summary2 "github.com/auho/go-toolkit/elasticsearch/indices/summary"
	elasticsearch2 "github.com/elastic/go-elasticsearch/v7"
	"github.com/spf13/cobra"
)

type summary struct {
	app.UseCase
}

func (s *summary) all() {
	s.PfVoid("all indices summary", func() error {
		_s, err := summary2.NewSummary(elasticsearch2.Config{
			Addresses: app.App.GetEnv().BaseEs.Address,
		})
		if err != nil {
			return err
		}

		rets, err := _s.AllIndices()
		if err != nil {
			return err
		}

		for _, ret := range rets {
			_ = os.WriteFile("output/es/"+ret.Index+".mapping.json", []byte(ret.Mapping), 0644)
			_ = os.WriteFile("output/es/"+ret.Index+".settings.json", []byte(ret.Settings), 0644)
		}

		return nil
	})
}

var _summary *summary

var _summaryCmd = &cobra.Command{
	Use: "summary",
	Run: func(cmd *cobra.Command, args []string) {
		_summary.all()
	},
}

func initSummary(parentCmd *cobra.Command) {
	_summary = &summary{}
	_summary.InitUseCase(parentCmd)

	parentCmd.AddCommand(_summaryCmd)
}

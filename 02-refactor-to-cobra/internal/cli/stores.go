package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

// CobraFunc is an alias to improve readability, we could reuse CobraFn instead but created for testing purpose
type CobraFunc func(*cobra.Command, []string)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}


func InitStoresCmd() *cobra.Command {
	storesCmd := &cobra.Command{
		Use:   "stores",
		Short: "Prints data about stores",
		Run:   runStoresFn(),
	}

	storesCmd.Flags().StringP(idFlag, "i", "", "id of the store")

	return storesCmd
}

func runStoresFn() CobraFunc {
	return func(cmd *cobra.Command, args []string) {
		idValue, _ := cmd.Flags().GetString(idFlag)

		if idValue != "" {
			fmt.Println(stores[idValue])
		} else {
			fmt.Println(stores)
		}
	}
}

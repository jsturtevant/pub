package plan

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	cobraExt "github.com/devigned/pub/pkg/cobra"
	"github.com/devigned/pub/pkg/partner"
)

func init() {
	listCmd.Flags().StringVarP(&listPlansArgs.Publisher, "publisher", "p", "", "publisher ID for your Cloud Partner Provider")
	_ = listCmd.MarkFlagRequired("publisher")
	listCmd.Flags().StringVarP(&listPlansArgs.Offer, "offer", "o", "", "String that uniquely identifies the offer.")
	_ = listCmd.MarkFlagRequired("offer")
	rootCmd.AddCommand(listCmd)
}

type (
	// ListPlansArgs are the arguments for `plans list` command
	ListPlansArgs struct {
		Publisher string
		Offer     string
	}
)

var (
	listPlansArgs ListPlansArgs
	listCmd       = &cobra.Command{
		Use:   "list",
		Short: "list all plans for a given offer and publisher",
		Run: cobraExt.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) {
			client, err := getClient()
			if err != nil {
				log.Fatalf("unable to create Cloud Partner Portal client: %v", err)
			}

			offer, err := client.GetOffer(ctx, partner.ShowOfferParams{
				PublisherID: listPlansArgs.Publisher,
				OfferID:     listPlansArgs.Offer,
			})

			if err != nil {
				log.Fatalf("unable to list offers: %v", err)
			}

			printPlans(offer.Definition.Plans)
		}),
	}
)

func printPlans(plans []partner.Plan) {
	bits, err := json.Marshal(plans)
	if err != nil {
		log.Fatalf("failed to print plans: %v", err)
	}
	fmt.Print(string(bits))
}

package card

import (
	"bank/pkg/bank/types"
)

// PaymentSourse 
func PaymentSource(cards []types.Card) (paymentSources []types.PaymentSource) {
	for _, v := range cards {
		if v.Balance < 0 {
			if !v.Active {
				continue
			}
			
		}
	paymentSources = append(paymentSources, types.PaymentSource{
			Type: "card",
			Number: string(v.PAN),
			Balance:v.Balance,
		},)
		}
		return paymentSources
}
package card

import (
	"bank/pkg/bank/types"
)

// PaymentSourse 
func PaymentSources(cards []types.Card) (paymentSource []types.PaymentSources) {
	for _, v := range cards {
		if v.Balance < 0 {
			continue
		}
			if !v.Active {
				continue
		}	

	paymentSource = append(paymentSource, types.PaymentSources{
			Type: "card",
			Number: string(v.PAN),
			Balance:v.Balance,
	},)

	}
		return paymentSource	
}
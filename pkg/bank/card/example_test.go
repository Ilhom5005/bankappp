package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card {
		{
			PAN:  "5058 2562 5451 0000",
			Balance: 1_000_00,
			Active: true,
		},
		{
			PAN:  "5058 2562 5451 1111",
			Balance: -1,
			Active: false,
		},
		{
			PAN:  "5058 2562 5451 2222",
			Balance: 5_000,
			Active: true,
		},
		{
			PAN:  "5058 2562 5451 8888",
			Balance: -2,
			Active: true,
		},
		{
			PAN:  "5058 2562 5451 9999",
			Balance: 0,
			Active: true,
		},
	}

	paymentSource := PaymentSources(cards)

	for _, v := range paymentSource {
		fmt.Println(v.Number)
		
	}

	//Output:
	//5058 2562 5451 0000
	//5058 2562 5451 2222
	//5058 2562 5451 9999
}

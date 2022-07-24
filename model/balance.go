package model

import "errors"

var (
	balanceHistory []Balance
)

func init() {
	balanceHistory = []Balance{}
}

type Balance struct {
	ID      int
	Segment int
	Notes   string `json:"notes"`
	Amount  int    `json:"amount"`
	Email   string
}

type BalanceModel struct{}

func (bm *BalanceModel) Insert(newData Balance) (Balance, error) {
	if len(balanceHistory) == 0 {
		newData.ID = GenerateID(0)
	} else {
		newData.ID = GenerateID(balanceHistory[len(balanceHistory)-1].ID)
	}
	balanceHistory = append(balanceHistory, newData)
	return newData, nil
}

func (bm *BalanceModel) GetBalance(userEmail string) ([]Balance, int, error) {
	var res []Balance
	var total int
	for _, val := range balanceHistory {
		if val.Email == userEmail {
			res = append(res, val)
			total += val.Amount
		}
	}

	if len(res) == 0 {
		return nil, -1, errors.New("no record found")
	}

	return res, total, nil
}

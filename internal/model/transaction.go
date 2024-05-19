package model

type Transaction struct {
	ContractNumber string `json:"contract_number"`
	NIK            string `json:"nik"`
	OTR            int    `json:"otr"`
	AdminFee       int    `json:"admin_fee"`
	Installment    int    `json:"installment"`
	Interest       int    `json:"interest"`
	AssetName      string `json:"asset_name"`
	LoanDate       string `json:"loan_date"`
	DueDate        string `json:"due_date"`
	IsPaid         bool   `json:"is_paid"`
}

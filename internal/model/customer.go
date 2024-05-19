package model

type Customer struct {
	NIK             string `json:"nik"`
	FullName        string `json:"full_name"`
	LegalName       string `json:"legal_name"`
	BirthPlace      string `json:"birth_place"`
	BirthDate       string `json:"birth_date"`
	Salary          int    `json:"salary"`
	PhotoKTP        string `json:"photo_ktp"`
	PhotoSelfie     string `json:"photo_selfie"`
	OneMonthLimit   int    `json:"one_month_limit"`
	TwoMonthLimit   int    `json:"two_month_limit"`
	ThreeMonthLimit int    `json:"three_month_limit"`
	FourMonthLimit  int    `json:"four_month_limit"`
}

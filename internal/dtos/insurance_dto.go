package dtos

type InsuranceDto struct {
	InsuranceType string  `json:"insurance_type"`
	ObjectValue   float64 `json:"object_value"`
	FromTime      string  `json:"time_from"`
	ToTime        string  `json:"time_to"`
	Value         float64 `json:"value"`
}

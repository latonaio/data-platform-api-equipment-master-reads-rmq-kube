package requests

type OwnerBusinessPartner struct {
	Equipment                int     `json:"Equipment"`
	OwnerBusinessPartner     int     `json:"OwnerBusinessPartner"`
	ValidityStartDate        string  `json:"ValidityStartDate"`
	ValidityEndDate          string  `json:"ValidityEndDate"`
	CreationDate             string  `json:"CreationDate"`
	LastChangeDate           string  `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

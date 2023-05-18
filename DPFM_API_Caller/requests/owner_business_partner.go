package requests

type OwnerBusinessPartner struct {
	Equipment                int     `json:"Equipment"`
	OwnerBusinessPartner     int     `json:"OwnerBusinessPartner"`
	ValidityStartDate        string  `json:"ValidityStartDate"`
	ValidityEndDate          string  `json:"ValidityEndDate"`
	CreationDate             *string `json:"CreationDate"`
	BusinessPartnerEquipment *int    `json:"BusinessPartnerEquipment"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

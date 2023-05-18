package requests

type BusinessPartner struct {
	Equipment                  int    `json:"Equipment"`
	EquipmentPartnerObjectNmbr int    `json:"EquipmentPartnerObjectNmbr"`
	BusinessPartner            int    `json:"BusinessPartner"`
	PartnerFunction            string `json:"PartnerFunction"`
	ValidityStartDate          string `json:"ValidityStartDate"`
	ValidityEndDate            string `json:"ValidityEndDate"`
	CreationDate               string `json:"CreationDate"`
	IsMarkedForDeletion        *bool  `json:"IsMarkedForDeletion"`
}

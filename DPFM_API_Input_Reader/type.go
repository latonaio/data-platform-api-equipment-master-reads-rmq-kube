package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	General           General  `json:"EquipmentMaster"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type General struct {
	Equipment						int			`json:"Equipment"`
	ValidityStartDate				*string		`json:"ValidityStartDate"`
	ValidityEndDate					*string		`json:"ValidityEndDate"`
	EquipmentName					*string		`json:"EquipmentName"`
	EquipmentType					*string		`json:"EquipmentType"`
	EquipmentCategory				*string		`json:"EquipmentCategory"`
	TechnicalObjectType				*string		`json:"TechnicalObjectType"`
	GrossWeight						*float32	`json:"GrossWeight"`
	NetWeight						*float32	`json:"NetWeight"`
	WeightUnit						*string		`json:"WeightUnit"`
	SizeOrDimensionText				*string		`json:"SizeOrDimensionText"`
	InventoryNumber					*string		`json:"InventoryNumber"`
	OperationStartDate				*string		`json:"OperationStartDate"`
	OperationStartTime				*string		`json:"OperationStartTime"`
	OperationEndDate				*string		`json:"OperationEndDate"`
	OperationEndTime				*string		`json:"OperationEndTime"`
	EquipmentStandardID				*string		`json:"EquipmentStandardID"`
	EquipmentIndustryStandardName	*string		`json:"EquipmentIndustryStandardName"`
	BarcodeType						*string		`json:"BarcodeType"`
	AcquisitionDate					*string		`json:"AcquisitionDate"`
	Manufacturer					*int		`json:"Manufacturer"`
	ManufacturerCountry				*string		`json:"ManufacturerCountry"`
	ConstructionYear				*int		`json:"ConstructionYear"`
	ConstructionMonth				*int		`json:"ConstructionMonth"`
	ConstructionDate				*string		`json:"ConstructionDate"`
	ManufacturerPartNmbr			*string		`json:"ManufacturerPartNmbr"`
	ManufacturerSerialNumber		*string		`json:"ManufacturerSerialNumber"`
	MaintenancePlantBusinessPartner	*int		`json:"MaintenancePlantBusinessPartner"`
	MaintenancePlant				*string		`json:"MaintenancePlant"`
	AssetLocation					*string		`json:"AssetLocation"`
	AssetRoom						*string		`json:"AssetRoom"`
	PlantSection					*string		`json:"PlantSection"`
	WorkCenter						*int		`json:"WorkCenter"`
	Project							*int		`json:"Project"`
	WBSElement						*int		`json:"WBSElement"`
	MaintenancePlannerGroup			*string		`json:"MaintenancePlannerGroup"`
	CatalogProfile					*string		`json:"CatalogProfile"`
	FunctionalLocation				*string		`json:"FunctionalLocation"`
	SuperordinateEquipment			*int		`json:"SuperordinateEquipment"`
	EquipInstallationPositionNmbr	*string		`json:"EquipInstallationPositionNmbr"`
	BillOfMaterial					*int		`json:"BillOfMaterial"`
	BillOfMaterialItem				*int		`json:"BillOfMaterialItem"`
	EquipmentIsAvailable			*bool		`json:"EquipmentIsAvailable"`
	EquipmentIsInstalled			*bool		`json:"EquipmentIsInstalled"`
	EquipHasSubOrdinateEquipment	*bool		`json:"EquipHasSubOrdinateEquipment"`
	MasterFixedAsset				*string		`json:"MasterFixedAsset"`
	FixedAsset						*string		`json:"FixedAsset"`
	CreationDate					*string		`json:"CreationDate"`
	LastChangeDate					*string		`json:"LastChangeDate"`
	IsMarkedForDeletion				*bool		`json:"IsMarkedForDeletion"`
	OwnerBusinessPartner            []OwnerBusinessPartner `json:"OwnerBusinessPartner"`
	BusinessPartner                 []BusinessPartner      `json:"BusinessPartner"`
	Address                         []Address              `json:"Address"`
	GeneralDoc                      []GeneralDoc           `json:"GeneralDoc"`
}

type Address struct {
	Equipment   int     `json:"Equipment"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}

type OwnerBusinessPartner struct {
	Equipment                int     `json:"Equipment"`
	OwnerBusinessPartner     int     `json:"OwnerBusinessPartner"`
	ValidityStartDate        string  `json:"ValidityStartDate"`
	ValidityEndDate          string  `json:"ValidityEndDate"`
	CreationDate             *string `json:"CreationDate"`
	LastChangeDate           *string `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type BusinessPartner struct {
	Equipment                  int     `json:"Equipment"`
	EquipmentPartnerObjectNmbr int     `json:"EquipmentPartnerObjectNmbr"`
	BusinessPartner            *int    `json:"BusinessPartner"`
	PartnerFunction            *string `json:"PartnerFunction"`
	ValidityStartDate          *string `json:"ValidityStartDate"`
	ValidityEndDate            *string `json:"ValidityEndDate"`
	CreationDate               *string `json:"CreationDate"`
	LastChangeDate             *string `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}

type GeneralDoc struct {
	Equipment                int     `json:"Equipment"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

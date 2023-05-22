package dpfm_api_output_formatter

import (
	"data-platform-api-equipment-master-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	general := make([]General, 0)

	i := 0
	for rows.Next() {
		pm := &requests.General{}
		i++

		err := rows.Scan(
			&pm.Equipment,
			&pm.BusinessPartner,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.EquipmentName,
			&pm.EquipmentType,
			&pm.EquipmentCategory,
			&pm.TechnicalObjectType,
			&pm.GrossWeight,
			&pm.NetWeight,
			&pm.WeightUnit,
			&pm.SizeOrDimensionText,
			&pm.InventoryNumber,
			&pm.OperationStartDate,
			&pm.OperationStartTime,
			&pm.OperationEndDate,
			&pm.OperationEndTime,
			&pm.EquipmentStandardID,
			&pm.EquipmentIndustryStandardName,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
			&pm.BarcodeType,
			&pm.AcquisitionDate,
			&pm.Manufacturer,
			&pm.ManufacturedCountry,
			&pm.ConstructionYear,
			&pm.ConstructionMonth,
			&pm.ManufacturerPartNmbr,
			&pm.ManufacturerSerialNumber,
			&pm.MaintenancePlantBusinessPartner,
			&pm.MaintenancePlant,
			&pm.AssetLocation,
			&pm.AssetRoom,
			&pm.PlantSection,
			&pm.WorkCenter,
			&pm.Project,
			&pm.MaintenancePlannerGroup,
			&pm.CatalogProfile,
			&pm.FunctionalLocation,
			&pm.SuperordinateEquipment,
			&pm.EquipInstallationPositionNmbr,
			&pm.EquipmentIsAvailable,
			&pm.EquipmentIsInstalled,
			&pm.EquipIsAllocToSuperiorEquip,
			&pm.EquipHasSubOrdinateEquipment,
			&pm.MasterFixedAsset,
			&pm.FixedAsset,
			&pm.CreationDate,
			&pm.LastChangeDateTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		general = append(general, General{
			Equipment:                       data.Equipment,
			BusinessPartner:                 data.BusinessPartner,
			ValidityStartDate:               data.ValidityStartDate,
			ValidityEndDate:                 data.ValidityEndDate,
			EquipmentName:                   data.EquipmentName,
			EquipmentType:                   data.EquipmentType,
			EquipmentCategory:               data.EquipmentCategory,
			TechnicalObjectType:             data.TechnicalObjectType,
			GrossWeight:                     data.GrossWeight,
			NetWeight:                       data.NetWeight,
			WeightUnit:                      data.WeightUnit,
			SizeOrDimensionText:             data.SizeOrDimensionText,
			InventoryNumber:                 data.InventoryNumber,
			OperationStartDate:              data.OperationStartDate,
			OperationStartTime:              data.OperationStartTime,
			OperationEndDate:                data.OperationEndDate,
			OperationEndTime:                data.OperationEndTime,
			EquipmentStandardID:             data.EquipmentStandardID,
			EquipmentIndustryStandardName:   data.EquipmentIndustryStandardName,
			CountryOfOrigin:                 data.CountryOfOrigin,
			CountryOfOriginLanguage:         data.CountryOfOriginLanguage,
			BarcodeType:                     data.BarcodeType,
			AcquisitionDate:                 data.AcquisitionDate,
			Manufacturer:                    data.Manufacturer,
			ManufacturedCountry:             data.ManufacturedCountry,
			ConstructionYear:                data.ConstructionYear,
			ConstructionMonth:               data.ConstructionMonth,
			ManufacturerPartNmbr:            data.ManufacturerPartNmbr,
			ManufacturerSerialNumber:        data.ManufacturerSerialNumber,
			MaintenancePlantBusinessPartner: data.MaintenancePlantBusinessPartner,
			MaintenancePlant:                data.MaintenancePlant,
			AssetLocation:                   data.AssetLocation,
			AssetRoom:                       data.AssetRoom,
			PlantSection:                    data.PlantSection,
			WorkCenter:                      data.WorkCenter,
			Project:                         data.Project,
			MaintenancePlannerGroup:         data.MaintenancePlannerGroup,
			CatalogProfile:                  data.CatalogProfile,
			FunctionalLocation:              data.FunctionalLocation,
			SuperordinateEquipment:          data.SuperordinateEquipment,
			EquipInstallationPositionNmbr:   data.EquipInstallationPositionNmbr,
			EquipmentIsAvailable:            data.EquipmentIsAvailable,
			EquipmentIsInstalled:            data.EquipmentIsInstalled,
			EquipIsAllocToSuperiorEquip:     data.EquipIsAllocToSuperiorEquip,
			EquipHasSubOrdinateEquipment:    data.EquipHasSubOrdinateEquipment,
			MasterFixedAsset:                data.MasterFixedAsset,
			FixedAsset:                      data.FixedAsset,
			CreationDate:                    data.CreationDate,
			LastChangeDateTime:              data.LastChangeDateTime,
			IsMarkedForDeletion:             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &general, nil
}

func ConvertToAddress(rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	address := make([]Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.Equipment,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &address, err
		}

		data := pm
		address = append(address, Address{
			Equipment:   data.Equipment,
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &address, nil
	}

	return &address, nil
}

func ConvertToBusinessPartner(rows *sql.Rows) (*[]BusinessPartner, error) {
	defer rows.Close()
	businessPartner := make([]BusinessPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.BusinessPartner{}

		err := rows.Scan(
			&pm.Equipment,
			&pm.EquipmentPartnerObjectNmbr,
			&pm.BusinessPartner,
			&pm.PartnerFunction,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &businessPartner, err
		}

		data := pm
		businessPartner = append(businessPartner, BusinessPartner{
			Equipment:                  data.Equipment,
			EquipmentPartnerObjectNmbr: data.EquipmentPartnerObjectNmbr,
			BusinessPartner:            data.BusinessPartner,
			PartnerFunction:            data.PartnerFunction,
			ValidityStartDate:          data.ValidityStartDate,
			ValidityEndDate:            data.ValidityEndDate,
			CreationDate:               data.CreationDate,
			IsMarkedForDeletion:        data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &businessPartner, nil
	}

	return &businessPartner, nil
}

func ConvertToOwnerBusinessPartner(rows *sql.Rows) (*[]OwnerBusinessPartner, error) {
	defer rows.Close()
	ownerBusinessPartner := make([]OwnerBusinessPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OwnerBusinessPartner{}

		err := rows.Scan(
			&pm.Equipment,
			&pm.OwnerBusinessPartner,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.BusinessPartnerEquipment,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &ownerBusinessPartner, err
		}

		data := pm
		ownerBusinessPartner = append(ownerBusinessPartner, OwnerBusinessPartner{
			Equipment:                data.Equipment,
			OwnerBusinessPartner:     data.OwnerBusinessPartner,
			ValidityStartDate:        data.ValidityStartDate,
			ValidityEndDate:          data.ValidityEndDate,
			CreationDate:             data.CreationDate,
			BusinessPartnerEquipment: data.BusinessPartnerEquipment,
			IsMarkedForDeletion:      data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &ownerBusinessPartner, nil
	}

	return &ownerBusinessPartner, nil
}

func ConvertToGeneralDoc(rows *sql.Rows) (*[]GeneralDoc, error) {
	defer rows.Close()
	generalDoc := make([]GeneralDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.GeneralDoc{}

		err := rows.Scan(
			&pm.Equipment,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &generalDoc, err
		}

		data := pm
		generalDoc = append(generalDoc, GeneralDoc{
			Equipment:                data.Equipment,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &generalDoc, nil
	}

	return &generalDoc, nil
}

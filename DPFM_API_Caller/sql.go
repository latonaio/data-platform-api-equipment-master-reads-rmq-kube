package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-equipment-master-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-equipment-master-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *[]dpfm_api_output_formatter.General
	var ownerBusinessPartner *[]dpfm_api_output_formatter.OwnerBusinessPartner
	var businessPartner *[]dpfm_api_output_formatter.BusinessPartner
	var address *[]dpfm_api_output_formatter.Address
	var docs *[]dpfm_api_output_formatter.GeneralDoc
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				general = c.General(mtx, input, output, errs, log)
			}()
		case "Generals":
			func() {
				general = c.Generals(mtx, input, output, errs, log)
			}()
		case "OwnerBusinessPartner":
			func() {
				ownerBusinessPartner = c.OwnerBusinessPartner(mtx, input, output, errs, log)
			}()
		case "BusinessPartner":
			func() {
				businessPartner = c.BusinessPartner(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "GeneralDocs":
			func() {
				docs = c.GeneralDocs(mtx, input, output, errs, log)
			}()

		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:              general,
		OwnerBusinessPartner: ownerBusinessPartner,
		BusinessPartner:      businessPartner,
		Address:              address,
		GeneralDoc:           docs,
	}

	return data
}

func (c *DPFMAPICaller) General(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	equipment := input.General.Equipment

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_master_general_data
		WHERE Equipment = ?;`, equipment,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Generals(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.General {
	where := "WHERE 1 = 1"
	if input.General.BusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND BusinessPartner = %v", where, *input.General.BusinessPartner)
	}
	if input.General.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.General.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_master_general_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) OwnerBusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.OwnerBusinessPartner {
	var args []interface{}
	equipment := input.General.Equipment
	ownerBusinessPartner := input.General.OwnerBusinessPartner

	cnt := 0
	for _, v := range ownerBusinessPartner {
		args = append(args, equipment, v.OwnerBusinessPartner, v.ValidityStartDate, v.ValidityEndDate)
		cnt++
	}

	repeat := strings.Repeat("(?,?,?,?),", cnt-1) + "(?,?,?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_master_owner_business_partner_data
		WHERE (Equipment, OwnerBusinessPartner, ValidityStartDate, ValidityEndDate) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToOwnerBusinessPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) BusinessPartner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BusinessPartner {
	var args []interface{}
	equipment := input.General.Equipment
	businessPartner := input.General.EMBusinessPartner

	cnt := 0
	for _, v := range businessPartner {
		args = append(args, equipment, v.EquipmentPartnerObjectNmbr)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_master_business_partner_data
		WHERE (Equipment, EquipmentPartnerObjectNmbr) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToBusinessPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) GeneralDocs(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.GeneralDoc {
	where := "WHERE ( Equipment, DocIssuerBusinessPartner ) IN "
	in := ""
	for _, v := range input.General.GeneralDoc {
		in = fmt.Sprintf("%s ( %d, %d ), ", in, v.Equipment, *v.DocIssuerBusinessPartner)
	}

	where = fmt.Sprintf("%s ( %s )", where, in[:len(in)-2])

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_master_general_doc_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneralDoc(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	equipment := input.General.Equipment
	address := input.General.Address

	cnt := 0
	for _, v := range address {
		args = append(args, equipment, v.AddressID)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_equipment_master_address_data
		WHERE (Equipment, AddressID) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

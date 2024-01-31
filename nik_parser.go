package nikparser

import (
	"github.com/fanchann/nik-parser/functions"
	"github.com/fanchann/nik-parser/models"
	"github.com/fanchann/nik-parser/utils"
)

func ParseNIK(NIK string) (*models.Response, error) {
	nik := new(functions.NIK)
	nik.NIKNumber = NIK

	if err := nik.ValidateNIK(); err != nil {
		return nil, err
	}

	response := models.Response{
		NIK:        NIK,
		Province:   nik.GetProvince(),
		City:       nik.GetCity(),
		District:   nik.GetDistrict(),
		Uniqcode:   nik.GetUniqcode(),
		Gender:     nik.GetGender(),
		BornDay:    nik.GetBornDay(),
		BornMonth:  utils.IntConvertToMonth(nik.GetBornMonth()),
		BornYear:   nik.GetBornYear(),
		PostalCode: nik.GetPostalCode(),
	}

	return &response, nil
}

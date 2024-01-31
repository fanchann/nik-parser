package nikparser

import (
	"fmt"

	"github.com/fanchann/nik-parser/functions"
	"github.com/fanchann/nik-parser/models"
	"github.com/fanchann/nik-parser/utils"
)

type NIKOpts struct {
	District   string
	PostalCode string
	Ttl        int
}

func NIKGen(nikopts ...*NIKOpts) []string {
	var niks []string

	defaultTtl := 10

	if len(nikopts) == 0 {
		for i := 0; i < defaultTtl; i++ {
			randomDistrict := utils.GetRandomNumberFromModel(models.District)
			randomBirth := utils.GenRandomBirth()
			randomStr := utils.GenRandomString()

			nikData := fmt.Sprintf("%s%s%s", randomDistrict, randomBirth, randomStr)
			niks = append(niks, nikData)
		}
	}

	for _, opt := range nikopts {
		district := utils.GetKeyByValue(models.District, fmt.Sprintf("%s -- %s", opt.District, opt.PostalCode))

		if opt.Ttl == 0 {
			opt.Ttl = 10
		}
		for i := 0; i < opt.Ttl; i++ {
			nikData := fmt.Sprintf("%s%s%s", district, utils.GenRandomBirth(), utils.GenRandomString())
			niks = append(niks, nikData)
		}
	}

	var validNIK []string

	// check nik is valid
	checkNIK := new(functions.NIK)
	for _, nik := range niks {
		checkNIK.NIKNumber = nik
		if err := checkNIK.ValidateNIK(); err == nil {
			validNIK = append(validNIK, nik)
		}
	}

	return validNIK
}

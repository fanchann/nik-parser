package functions

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/fanchann/nik-parser/models"
)

type NIK struct {
	NIKNumber string
	Province  string
	City      string
	District  string
	BornAt    string
	Gender    string
	UniqCode  string
}

func (n *NIK) GetGender() string {
	// check charactheristic gender by num in NIK
	if len(n.NIKNumber[6:8]) > 40 {
		return "FEMALE"
	}
	return "MALE"
}

func (n *NIK) ValidateNIK() error {
	if len(n.NIKNumber) != 16 {
		return errors.New("nik not valid!")
	}
	return nil
}

func (n *NIK) GetProvince() string {
	return models.Province[n.NIKNumber[0:2]]
}

func (n *NIK) GetCity() string {
	return models.City[n.NIKNumber[0:4]]
}

func (n *NIK) GetDistrict() string {
	district := models.District[n.NIKNumber[0:6]]
	disSplit := strings.Split(district, "--")
	return strings.Trim(disSplit[0], " ")
}

func (n *NIK) GetPostalCode() string {
	postCode := models.District[n.NIKNumber[0:6]]
	postCodeSplit := strings.Split(postCode, "--")
	return strings.Trim(postCodeSplit[1], " ")
}

func (n *NIK) GetBornDay() int {
	// n.BornAt = n.NIKNumber[6:12]
	bornDayInt, _ := strconv.Atoi(n.NIKNumber[6:8])

	if bornDayInt > 40 {
		bornDayInt = bornDayInt - 40
	}
	return bornDayInt
}

func (n *NIK) GetBornMonth() int {
	n.BornAt = n.NIKNumber[6:12]
	monthInt, _ := strconv.Atoi(n.BornAt[2:4])
	return monthInt
}

func (n *NIK) GetBornYear() int {
	n.BornAt = n.NIKNumber[6:12]

	yearNow, _ := strconv.Atoi(time.Now().Format("06"))

	nikYear, _ := strconv.Atoi(n.BornAt[4:6])

	if nikYear < yearNow {
		nikYear += 2000
	} else {
		nikYear += 1900
	}
	return nikYear
}

func (n *NIK) GetUniqcode() string {
	return n.NIKNumber[12:16]
}

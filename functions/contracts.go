package functions

type INIK interface {
	ValidateNIK() error
	GetProvince() string
	GetCity() string
	GetDistrict() string
	GetGender() string
	GetPostalCode() string
	GetBornDay() int
	GetBornMonth() int
	GetBornYear() int
	GetUniqcode() string
}

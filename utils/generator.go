package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomNumberFromModel(mapData map[string]string) string {
	rand.Seed(time.Now().UnixNano())
	keys := make([]string, 0, len(mapData))
	for key, _ := range mapData {
		keys = append(keys, key)
	}

	rand.Shuffle(len(keys), func(i, j int) {
	})

	return keys[0]
}

func GenRandomBirth() string {
	rand.Seed(time.Now().UnixNano())
	dateInt := rand.Intn(29+40) + 1 // 31
	monthInt := rand.Intn(12) + 1   // month
	yearInt := rand.Intn(17) + 88   //  88,89,90,91,92,93,94,95,96,97,98,99,00,01,02,03,04

	result := fmt.Sprintf("%02d%02d%02d", dateInt, monthInt, yearInt)

	if len(result) != 6 {
		return GenRandomBirth()
	}
	return result
}

func GenRandomString() string {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(10000)
	result := fmt.Sprintf("%04d", randomNumber)

	return result
}

func GetKeyByValue(m map[string]string, value string) string {
	{
		for key, val := range m {
			if val == value {
				return key
			}
		}
		return ""
	}
}

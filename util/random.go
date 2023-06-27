package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	alphabet = "qwertyuiopasdfghjklzxcvbnm"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	l := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomClient() string {
	return randomString(6)
}

func RandomCountry() string {
	countries := []string{"RUSSIAN FEDERATION",
		"KAZAKHSTAN",
		"GERMANY",
		"CHINA",
		"SOUTH KOREA",
	}
	n := len(countries)
	return countries[rand.Intn(n)]
}

func RandomCity() string {
	towns := []string{
		"MOSCOW",
		"ROSTOV",
		"KAZAN",
		"SAINT-PETERSBURG",
	}
	n := len(towns)
	return towns[rand.Intn(n)]
}

func RandomPhoneNumber() string {
	var number strings.Builder
	tire := "-"
	operators := []string{"+7-926", "+7-925", "+7-915", "+7-903", "+7-999", "+7-909"}
	randOperator := operators[rand.Intn(len(operators))]

	number.WriteString(randOperator)
	number.WriteString(tire)
	number.WriteString(randomIntStringForNumberTriple())
	number.WriteString(tire)
	number.WriteString(randomIntStringForNumberDouble())
	number.WriteString(randomIntStringForNumberDouble())

	return number.String()
}

func randomIntStringForNumberDouble() string {
	number := rand.Intn(100-10) + 10
	return strconv.Itoa(number)
}
func randomIntStringForNumberTriple() string {
	number := rand.Intn(1000-100) + 100
	return strconv.Itoa(number)
}

func RandomAddress() string {
	return randomString(10)
}

func RandomCar() string {
	model_name := []string{
		"BMW",
		"Mercedes-Benz",
		"LADA",
		"Hyundai",
		"Toyota",
		"Changan",
		"Audi",
		"Kia",
	}
	n := len(model_name)
	return model_name[rand.Intn(n)]
}

func RandomEquipment() string {
	eq := []string{
		"base",
		"Allroad",
		"luxury",
	}
	return eq[rand.Intn(len(eq))]
}

func RandomPrice() int64 {
	n := rand.Intn(100000-10000) + 10000
	return int64(n)
}

func RandomColor() string {
	color := []string{"black", "red", "sky", "brown", "grey", "white"}
	return color[rand.Intn(len(color))]
}

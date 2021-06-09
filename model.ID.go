package main

//Person godoc
var posistionValues = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var cities = map[string]string{
	"412": "بروجرد",
	"413": "بروجرد",
	"406": "خرم آباد",
	"407": "خرم آباد",
	"421": "دورود",
	"498": "بابلسر",
	"483": "چالوس",
	"227": "رامسر",
	"627": "کلاردشت",
}

const lengthOfNationalCode = 10
const modValue = 11

type Id struct {
	ID      string `json:"id"`
	IsValid bool   `json:"isvalid"`
	City    string `json:"city"`
}

func idValidation(id string) (*Id, error) {
	instance := Id{id, false, "تعریف نشده"}
	if instance.validation() {
		instance.IsValid = instance.validation()
		instance.City = instance.cityName()
	}
	return &instance, nil
	//return nil, errors.New("Person not found")
}

func (n *Id) cityName() string {
	if val, ok := cities[n.ID[0:3]]; ok {
		return val
	} else {
		return "تعریف نشده"
	}
}

func (n *Id) validation() bool {
	var (
		sum   int
		res   int
		index int
	)

	for index = 0; index < lengthOfNationalCode-1; index++ {
		sum = sum + posistionValues[index]*int(n.ID[index]-'0')
	}

	res = sum % modValue
	controlDigit := int(n.ID[lengthOfNationalCode-1] - '0')

	if res < 2 {
		if res == controlDigit {
			n.IsValid = true
			return true
		}
	} else {
		if res == modValue-controlDigit {
			n.IsValid = true
			return true
		}
	}
	n.IsValid = false
	return false
}

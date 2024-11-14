package pkg

import "time"

// should use "2006-01-02"
// like this type
func GetAge(birthday string) (int, error) {
	birthDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age = age - 1
	}
	return age, nil
}

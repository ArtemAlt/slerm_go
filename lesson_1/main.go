package main

import "fmt"

type HealthCheck struct {
	serialId int
	status   string
}

const PassStatus = "pass"
const FailStatus = "fail"

func GenerateCheck() []HealthCheck {
	result := make([]HealthCheck, 0, 5)
	for i := 0; i < 5; i++ {
		var status = FailStatus
		if i%2 == 0 {
			status = PassStatus
		}
		result = append(result, HealthCheck{i, status})
	}
	return result
}
func main() {
	fmt.Println("Тут будет выведен идентификатор")
	currentCheck := GenerateCheck()
	for key, value := range currentCheck {
		if value.status == PassStatus {
			fmt.Println(key)
		}

	}
}

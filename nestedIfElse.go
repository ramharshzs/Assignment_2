package main

func salary(age int, department string) string {
	if age > 20 {
		if department == "IT" {
			return "45000"
		} else if department == "CSEAIML" {
			return "60000"
		} else {
			return "30000"
		}
	} else {
		return "0"
	}
}

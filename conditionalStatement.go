package main

func weather(temperature int) string {
	if temperature < 20 {
		return "cold"
	} else if temperature > 30 {
		return "hot"
	} else {
		return "normal"
	}
}

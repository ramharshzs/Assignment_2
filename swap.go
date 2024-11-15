package main

func swapNo(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

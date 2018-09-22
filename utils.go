package main

func contains(ps []int, p int) bool {
	for _, a := range ps {
		if a == p {
			return true
		}
	}
	return false
}

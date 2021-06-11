package main

func Sum(n []int) int {
	var result int
	for _, i := range n {
		result += i
	}
	return result
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, n := range numbers {
		sums = append(sums, Sum(n))
	}
	return sums
}

func SumAllTrails(numbers ...[]int) []int {
	var results []int
	for _, number := range numbers {
		if len(number) > 0 {
			results = append(results, Sum(number[1:]))
		} else {
			results = append(results, 0)
		}
	}
	return results
}

func main() {

}

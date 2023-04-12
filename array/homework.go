package array_lesson

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var (
	ErrorFindMaxUnderBoundary = errors.New("boundary is not correct")
)

func FindSmallest(nums []int) int {
	min := 0

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func FindAverageTime(nums []int) float64 {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	result := float64(sum) / float64(len(nums))

	return math.Round(result*100) / 100
}

func FindMostProfitableClient(customers [][]int) int {
	maxIncome, indexMaxIncome := 0, 0

	for id, customer := range customers {
		sum := 0
		for _, money := range customer {
			sum += money
		}

		if maxIncome < sum {
			indexMaxIncome = id
			maxIncome = sum
		}
	}

	return indexMaxIncome
}

func FindMaxUnderBoundary(nums []int, boundary int) (int, error) {
	max := math.MinInt

	for _, num := range nums {
		if num < boundary && num > max {
			max = num
		}
	}

	if max == math.MinInt {
		return 0, ErrorFindMaxUnderBoundary
	}

	return max, nil
}

func FindLettersLearnedToday(letters string) string {
	var learnedLetters string

	bag := make(map[rune]int, 0)

	placeIndex := 0
	for _, letter := range letters {
		if _, ok := bag[letter]; ok {
			continue
		}

		bag[letter] = placeIndex
		placeIndex++
	}

	for i := 0; i < placeIndex; i++ {
		for letter, index := range bag {
			if index == i {
				learnedLetters += string(letter)
				break
			}
		}
	}

	return learnedLetters
}

func FindAvoidJailDueToTaxFraud(customers [][]int) int {
	doubleCustomer := make(map[int]bool, 0)

	for _, customer := range customers {
		for _, money := range customer {
			if ok, _ := doubleCustomer[money]; ok {
				return money
			}
			doubleCustomer[money] = true
		}
	}

	return -1
}

func CalcUniqNumber(apps []int) map[int]int {
	result := make(map[int]int, 0)

	for _, app := range apps {
		if _, ok := result[app]; ok {
			result[app]++
		} else {
			result[app] = 1
		}
	}

	return result
}

func CryptoCurrencyAnalysis(income string) map[string]float64 {
	result := make(map[string]float64)
	countIncome := make(map[string]int)

	incomes := strings.Fields(income)

	for _, singleIncome := range incomes {
		transaction := strings.Split(singleIncome, ":")
		transFloat, _ := strconv.ParseFloat(transaction[1], 32)

		if _, ok := result[transaction[0]]; ok {
			result[transaction[0]] += transFloat
			countIncome[transaction[0]]++
		} else {
			result[transaction[0]] = transFloat
			countIncome[transaction[0]] = 1
		}
	}

	for name, money := range result {
		result[name] = money / float64(countIncome[name])
	}

	return result
}

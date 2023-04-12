package array_lesson

import (
	"fmt"
	"log"
	"testing"
)

func TestFindSmallest(T *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{-1000, -100, -10, -1},
			-1000,
		},
		{
			[]int{-1000},
			-1000,
		},
		{
			[]int{-1000, -100, -10, -1, -1},
			-1000,
		},
		{
			[]int{-1000, -100, -10, -1, -1, 0},
			-1000,
		},
	}

	for _, test := range testCases {
		if test.expected != FindSmallest(test.nums) {
			log.Fatal("TestFindSmallest failed")
		}
	}
}

func TestFindAverageTime(T *testing.T) {
	testCases := []struct {
		nums     []int
		expected float64
	}{
		{
			[]int{9999},
			9999,
		},
		{
			[]int{9, 4, 1, 8, 7, 9, 4, 1, 8, 7, 8, 7, 18, 3, 13, 6, 5},
			6.94,
		},
	}

	for _, test := range testCases {
		if test.expected != FindAverageTime(test.nums) {
			log.Fatal("TestFindAverageTime failed")
		}
	}
}

func TestFindMostProfitableClient(T *testing.T) {
	testCases := []struct {
		nums     [][]int
		expected int
	}{
		{
			[][]int{
				{11, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{12, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{13, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{14, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{15, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{16, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{17, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{18, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{19, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				{10, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			9,
		},
		{
			[][]int{
				{1, 9999, -10},
				{1},
			},
			0,
		},
		{
			[][]int{
				{95, 67, 13, 55, 44, 11, 10},
				{7, 190, 4, 44, 11, 1, 99},
				{0, 5, -1, 500, 14, 90, 1},
			},
			2,
		},
	}

	for _, test := range testCases {
		if test.expected != FindMostProfitableClient(test.nums) {
			fmt.Println(FindMostProfitableClient(test.nums), test.expected)
			log.Fatal("TestFindMostProfitableClient failed")
		}
	}
}

func TestFindMaxUnderBoundary(T *testing.T) {
	testCases := []struct {
		nums           []int
		boundary       int
		expected       int
		exptectedError error
	}{
		{
			[]int{10, 20, 30, 40},
			-1,
			0,
			ErrorFindMaxUnderBoundary,
		},
		{
			[]int{10, 20, 30, 40},
			0,
			0,
			ErrorFindMaxUnderBoundary,
		},
		{
			[]int{10, 20, 30, 40},
			100,
			40,
			nil,
		},
	}

	for _, test := range testCases {
		result, err := FindMaxUnderBoundary(test.nums, test.boundary)
		if test.expected != result || err != test.exptectedError {
			log.Fatal("FindMaxUnderBoundary failed")
		}
	}
}

func TestFindLettersLearnedToday(T *testing.T) {
	testCases := []struct {
		letters  string
		expected string
	}{
		{
			"АААФФФФФФФЖЫЫЫЫБЫРВАААААЛГГГХЫХЫБЛИА",
			"АФЖЫБРВЛГХИ",
		},
		{
			"ОК",
			"ОК",
		},
		{
			"СКИЛЛБОКСТОПЧИК",
			"СКИЛБОТПЧ",
		},
		{
			"ААААААААА",
			"А",
		},
		{
			"",
			"",
		},
	}

	for _, test := range testCases {
		if test.expected != FindLettersLearnedToday(test.letters) {
			log.Fatal("FindLettersLearnedToday failed", test.expected)
		}
	}
}

func TestFindAvoidJailDueToTaxFraud(T *testing.T) {
	testCases := []struct {
		customers [][]int
		expected  int
	}{
		{
			[][]int{
				{12391203, 3828382, 334934939},
				{45345345, 5341312, 55345345},
				{334934939, 1234122, 657657},
			},
			334934939,
		},
		{
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
			},
			1,
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			-1,
		},
	}

	for _, test := range testCases {
		if test.expected != FindAvoidJailDueToTaxFraud(test.customers) {
			log.Fatal("TestFindAvoidJailDueToTaxFraud failed", test.expected)
		}
	}
}

func TestCalcUniqNumber(t *testing.T) {
	testCases := []struct {
		apps     []int
		expected map[int]int
	}{
		{
			[]int{
				79000000000,
				79000000000,
				79000000001,
				79000000002,
				79000000002,
				79000000003,
				79000000003,
				79000000003,
				79000000003,
				79000000004,
			},
			map[int]int{
				79000000000: 2,
				79000000001: 1,
				79000000002: 2,
				79000000003: 4,
				79000000004: 1,
			},
		},
	}

	for _, test := range testCases {

		result := CalcUniqNumber(test.apps)

		for number, count := range result {
			value, ok := test.expected[number]

			if ok && count == value {
				continue
			}

			log.Fatal("TestCalcUniqNumber failed", test.expected)
		}
	}
}

func TestCryptoCurrencyAnalysis(t *testing.T) {
	testCases := []struct {
		income   string
		expected map[string]float64
	}{
		{
			`
			BTC:42
			BTC:600
			BTC:900
			DOGE:123456
			DOGE:69420
			ETH:220
			ETH:666
			XMR:14
			XMR:88
			`,
			map[string]float64{
				"BTC":  514,
				"DOGE": 96438,
				"ETH":  443,
				"XMR":  51,
			},
		},
		{
			`
               BTC:600
                BTC:600
                BTC:600
                BTC:600
                BTC:600
                BTC:600
                BTC:600
			`,
			map[string]float64{
				"BTC": 600,
			},
		},
		{
			`
                BTC:1
                DOGE:2
			`,
			map[string]float64{
				"BTC":  1,
				"DOGE": 2,
			},
		},
		{
			`
                DOGE:69420
			`,
			map[string]float64{
				"DOGE": 69420,
			},
		},
	}

	for _, test := range testCases {
		result := CryptoCurrencyAnalysis(test.income)

		for number, count := range result {
			value, ok := test.expected[number]

			if ok && count == value {
				continue
			}

			log.Fatal("TestCryptoCurrencyAnalysis failed", test.expected)
		}
	}

}

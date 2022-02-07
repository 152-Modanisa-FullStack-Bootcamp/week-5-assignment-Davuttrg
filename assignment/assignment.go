package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	result := x + y
	return result, math.MaxUint32-y < x
}

func CeilNumber(f float64) float64 {

	if math.Mod(f, 0.25) == float64(0) {
		return f
	}
	return float64(int64(f/0.25)+1) * 0.25
}

func AlphabetSoup(s string) string {
	result := strings.Split(s, "")
	sort.Strings(result)
	return strings.Join(result, "")
}

func StringMask(s string, n uint) string {
	if len(s) <= int(n) {
		n = 0
	}
	if len(s) == 0 {
		return "*"
	}
	result := s[:int(n)]
	for i := 0; i < len(s)-int(n); i++ {
		result += "*"
	}
	return result
}

func WordSplit(arr [2]string) string {
	words := strings.Split(arr[1], ",")
	result := []string{}

forLoop:
	for _, v := range words {

		if i := strings.Index(arr[0], v); i != -1 {

			if i == 0 {
				result = append([]string{v}, result...)
			} else {
				result = append(result, v)
			}
			if len(result) == 2 {
				break forLoop
			}
		}
	}
	if len(result) == 2 {
		return strings.Join(result, ",")
	} else {
		return "not possible"
	}

}

func VariadicSet(i []interface{}) []interface{} {

	res := []interface{}{}
	temp := make(map[interface{}]bool)
	for _, v := range i {
		if !temp[v] {
			res = append(res, v)
		}
		temp[v] = true

	}
	return res
}

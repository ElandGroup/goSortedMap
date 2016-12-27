package main

import "sort"
import "fmt"

func sortedKeys(m *map[string]string) []string {
	sortedKeys := make([]string, len(*m))
	i := 0
	for key := range *m {
		sortedKeys[i] = key
		i++
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

func sortedMap(m *map[string]string) string {
	if len(*m) == 0 {
		return ""
	}
	sk := sortedKeys(m)
	var sortedData string
	for _, k := range sk {
		sortedData += "&" + k + "=" + (*m)[k]
	}
	return sortedData[1:]
}

func main() {
	s := map[string]string{"abb": "111", "bbb": "222", "ggg": "233", "eee": "212"}
	for i, res := range sortedKeys(&s) {
		fmt.Println(i, res, s[res])
	}
	fmt.Println(sortedMap(&s))

}

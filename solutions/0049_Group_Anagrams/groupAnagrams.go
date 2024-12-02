package GroupAnagrams

//package main

// sort and use map
func GroupAnagrams(strs []string) [][]string {
	record := make(map[string][]string)
	for _, value := range strs {
		sortedstr := sortString(value)
		record[sortedstr] = append(record[sortedstr], value)
	}
	tmpstrs := [][]string{}
	for _, value := range record {
		tmpstrs = append(tmpstrs, value)
	}
	return tmpstrs
}

func sortString(s string) string {
	if len(s) <= 1 {
		return s
	}
	chars := []rune(s) //字符串处理成字符切片
	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] > chars[j] {
				chars[i], chars[j] = chars[j], chars[i]
			}
		}
	}
	return string(chars)
}

// use PrimeNumber，by overload with int64
func GroupAnagrams1(strs []string) [][]string {
	record := make(map[int64][]string)
	Prime_dict := map[rune]int{
		'a': 2,
		'b': 3,
		'c': 5,
		'd': 7,
		'e': 11,
		'f': 13,
		'g': 17,
		'h': 19,
		'i': 23,
		'j': 29,
		'k': 31,
		'l': 37,
		'm': 41,
		'n': 43,
		'o': 47,
		'p': 53,
		'q': 59,
		'r': 61,
		's': 67,
		't': 71,
		'u': 73,
		'v': 79,
		'w': 83,
		'x': 89,
		'y': 97,
		'z': 101,
	}
	for _, value := range strs {
		tmp_key := int64(1)
		for _, charv := range value {
			tmp_key *= int64(Prime_dict[charv])
			//fmt.Printf("charv:%c, tmp_key:%d\n", charv, tmp_key)
		}
		record[tmp_key] = append(record[tmp_key], value)
	}
	tmpstrs := [][]string{}
	for _, value := range record {
		tmpstrs = append(tmpstrs, value)
	}
	return tmpstrs
}

/*
func main() {
	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs := []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"}
	fmt.Println(GroupAnagrams(strs))
}
*/

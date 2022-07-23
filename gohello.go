package main

import "fmt"

func zeroFilledSubarray(nums []int) int64 {
	var length []int
	tagL := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			tagL += 1
		} else if nums[i] != 0 {
			if tagL > 0 {
				length = append(length, tagL)
			}
			tagL = 0
		}
	}
	if tagL > 0 {
		length = append(length, tagL)
	}
	sum := 0
	for i := 0; i < len(length); i++ {
		sum += (length[i] * (length[i] + 1)) / 2
	}
	return int64(sum)
}

func bestHand(ranks []int, suits []byte) string {
	mapStr := map[byte]int{}
	Fl := 0
	Th := 0
	Pa := 0
	Hi := 0
	tag := 1
	mapStr[suits[0]] = 1
	for i := 1; i < len(suits); i++ {
		if _, ok := mapStr[suits[i]]; ok {
			tag = 0
		} else {
			tag += 1
		}
	}
	if tag == 0 {
		Fl = 1
	}
	if len(mapStr) == 5 {

	}
	mapRanks := map[int]int{}
	for i := 0; i < len(ranks); i++ {
		mapRanks[ranks[i]]++
	}
	for _, val := range mapRanks {
		if val == 2 {
			Pa = 1
		}
		if val >= 3 {
			Th = 1
		}
	}
	fmt.Println(mapRanks)
	if Fl == 1 {
		return "Flush"
	} else if Th == 1 {
		return "Three of a Kind"
	} else if Pa == 1 {
		return "Pair"
	} else if Hi == 1 {
		return "High Card"
	}
	return ""
}

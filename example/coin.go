/**
 * Author: richen
 * Date: 2020-08-10 14:39:27
 * LastEditTime: 2020-08-10 18:02:20
 * Description:
 * Copyright (c) - <richenlin(at)gmail.com>
 */
package main

import "fmt"

var (
	coins = 120
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func calcCoin(username string, left int) int {
	var sum = 0
	for _, char := range username {
		if left < 1 {
			break
		}
		switch char {
		case 'a', 'A':
			if left >= 1 {
				left -= 1
				sum += 1
			} else {
				sum += left
				left = 0
			}
		case 'e', 'E':
			if left >= 1 {
				left -= 1
				sum += 1
			} else {
				sum += left
				left = 0
			}
		case 'i', 'I':
			if left >= 2 {
				left -= 2
				sum += 2
			} else {
				sum += left
				left = 0
			}
		case 'o', 'O':
			if left >= 2 {
				left -= 2
				sum += 2
			} else {
				sum += left
				left = 0
			}
		case 'u', 'U':
			if left >= 5 {
				left -= 5
				sum += 5
			} else {
				sum += left
				left = 0
			}
		}
	}
	return sum
}

func disPatchCoin() int {
	var left int = coins
	for _, username := range users {
		allCoin := calcCoin(username, left)
		left = left - allCoin
		value, ok := distribution[username]
		if !ok {
			distribution[username] = allCoin
		} else {
			distribution[username] = value + allCoin
		}
	}
	return left
}

func main() {
	left := disPatchCoin()

	for username, coin := range distribution {
		fmt.Printf("user: %s have %d conis\n left conis: %d\n", username, coin, left)
	}

}

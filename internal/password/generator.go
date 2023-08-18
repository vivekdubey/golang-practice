package password

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

func specialCharactersList() []string {
	return []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")"}
}

func alphabetList() []string {
	var c []string
	for i := 97; i < 97+26; i++ {
		c = append(c, fmt.Sprintf("%c", i))
	}

	for i := 65; i < 65+26; i++ {
		c = append(c, fmt.Sprintf("%c", i))
	}
	return c
}

func numberList() []string {
	var n []string
	for i := 0; i < 10; i++ {
		n = append(n, strconv.Itoa(i))
	}
	return n
}

func RandomPassword(l, mn, mc, ms int) (string, error) {
	if mn+mc+ms > l {
		return "", errors.New("Minimum character + number + special characters >  total length")
	}
	var r = []string{"c", "n", "s"}
	var password string
	var counter = map[string]int{
		"c": mc,
		"n": mn,
		"s": ms,
	}
	c := alphabetList()
	n := numberList()
	s := specialCharactersList()

	for i := 0; i < l; {
		t := r[rand.Intn(len(r))]
		if t == "c" {
			if counter["n"]+counter["s"] >= l-i {
				continue
			}
			password = password + c[rand.Intn(len(c))]
			if counter["c"] > 0 {
				counter["c"] -= 1
			}
		} else if t == "n" {
			if counter["c"]+counter["s"] >= l-i {
				continue
			}
			password = password + n[rand.Intn(len(n))]
			if counter["n"] > 0 {
				counter["n"] -= 1
			}
		} else if t == "s" {
			if counter["c"]+counter["n"] >= l-i {
				continue
			}
			password = password + s[rand.Intn(len(s))]
			if counter["s"] > 0 {
				counter["s"] -= 1
			}
		}
		i++
	}
	return password, nil
}

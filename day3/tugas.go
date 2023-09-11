package main

import (
	"fmt"
	"slices"
)

func main() {
	// soal 1
	fmt.Println(ArrayMerge([]string{"asd","asf","fefw","geqrg", "asd"}, []string{"erba","asd","asf"}))
	
	// soal 2
	fmt.Println(Mapping([]string{"asd", "fdsa", "asd", "fes", "efs"}))

	// soal 3
	fmt.Println(MunculSekali("11442325"))
}

func MunculSekali(angka string) []string {
  arr := make([]string, len(angka))
  hasil := []string{}
  dobel := []string{}

  for i, r := range angka {
    if slices.Contains(arr, string(r)) {
			dobel = append(dobel, string(r)) 
		}
    arr[i] = string(r)
  }

  for _, v := range dobel {		
		hasil = slices.DeleteFunc(arr, func(n string) bool {
			return n == v
		})
	}
	
	return ArrayMerge(hasil, []string{})
}

func Mapping(slice []string) map[string]int {
	hasil := map[string]int{}
	
	for _, v := range slice {
		if _, key := hasil[v]; key {
			hasil[v]++
		} else {
			hasil[v] = 1
		}
	}
	
	return hasil
}

func contains(slice []string, val string) bool {
    for _, v := range slice {
        if v == val {
            return true
        }
    }
    return false
}

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := append(arrayA, arrayB...)
	hasil := []string{}
	
	for _, v := range merged {
		if !contains(hasil, v) {
			hasil = append(hasil, v)
		}
	}
	
	return hasil
}
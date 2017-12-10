package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "sort"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	checksum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
		nums := []int(nil)
		for _, token := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
		found := false
		for dpos, denominator := range nums {
			for _, numerator := range nums[:dpos] {
				if numerator % denominator == 0 {
					fmt.Printf("%d  / %d = %d\n", numerator, denominator, numerator / denominator)
					checksum += numerator / denominator
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found == false {
			fmt.Printf("failed: %d -> %d %s\n", nums[0], nums[len(nums)-1], scanner.Text()) 
		}
	}
	fmt.Printf("Checksum: %d\n", checksum);
}

package main

import "fmt"

/*

How to handle a scenario where we have to terminate the idle sessions which are older than 15 min

Soln - Use LRU cache
*/

/*

Question -1
need to find the length overloaaping part
"STRINGMATCH" and "MATCHBOX"
with O(n)
*/

func findOverlappingLength(str1, str2 string) int {
	len1 := len(str1)
	len2 := len(str2)

	// Initialize a matrix to store the lengths of common suffixes
	lcs := make([][]int, len1+1)
	for i := range lcs {
		lcs[i] = make([]int, len2+1)
	}

	// Variable to store the length of the longest common suffix
	maxLength := 0

	// Iterate through the characters of both strings
	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
				maxLength = max(maxLength, lcs[i][j])
			} else {
				lcs[i][j] = 0
			}
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str1 := "STRINGMATCH"
	str2 := "MATCHBOX"

	overlapLength := findOverlappingLength(str1, str2)
	fmt.Printf("Length of overlapping substring: %d\n", overlapLength)
}



/*
Question 2 - 
Develop an event scheduler system to schedule special events, parties, birthdays, friendship gatherings, and share-meetings

The system should allow users to input their event details such as the date, time, location, and other relevant information. It should also allow users to search for and book available venues. Additionally, the system should generate automated reminders to users a few days before the event date in order to ensure that all necessary preparations are made.
*/

/*
Question 3- 
song list in an array 
develop a song player to play song with shuffle without repeating any song 
*/
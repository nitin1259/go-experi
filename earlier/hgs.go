package main

/*
Given a hotel which has 10 floors [0-9] and each floor has 26 rooms [A-Z]. You are given a sequence of rooms, where + suggests room is booked, - room is freed. You have to find which room is booked maximum number of times.

You may assume that the list describes a correct sequence of bookings in chronological order; that is, only free rooms can be booked and only booked rooms can be freed. All rooms are initially free. Note that this does not mean that all rooms have to be free at the end. In case 2 rooms have been booked the same number of times, return the lexicographically smaller room.

You may assume:

N (length of input) is an integer within the range [1, 600] each element of array A is a string consisting of three characters: "+" or "-"; a digit "0"-"9"; and uppercase English letter "A" - "Z" the sequence is correct. That is every booked room was previously free and every freed room was previously booked.

Example: Input: ["+1A", "+3E", "-1A", "+4F", "+1A", "-3E"] Output: "1A" Explanation: 1A as it has been booked 2 times.

*/

import (
	"fmt"
)

func findMaxBookedRoom(bookings []string) string {
	roomCounts := make(map[string]int)
	maxCount := 0
	var maxRoom string

	for _, booking := range bookings {
		action, room := booking[0], booking[1:]

		if action == '+' {
			roomCounts[room]++
			if roomCounts[room] > maxCount || (roomCounts[room] == maxCount && room < maxRoom) {
				maxCount = roomCounts[room]
				maxRoom = room
			}
		} else if action == '-' {
			roomCounts[room]--
		}
	}

	return maxRoom
}

func main_hotelbooking() {
	bookings := []string{"+1A", "+3E", "-1A", "+4F", "+1A", "-3E"}
	result := findMaxBookedRoom(bookings)
	fmt.Println(result)
}
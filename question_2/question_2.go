package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// []int -> string
func joinIntSliceToString(s []int) string {
	var answer bytes.Buffer
	for _, value := range s {
		answer.WriteString(strconv.Itoa(value))
	}
	return answer.String()
}

func DecodeLeftRightEqual(encoded string) string {
	n := len(encoded)
	encodedSplit := strings.Split(encoded, "")
	mem := make([]int, n+1) // default value = 0

	//from left to right handle with R and =
	for i := 0; i < n; i++ {
		if encodedSplit[i] == "R" {
			mem[i+1] = mem[i] + 1 // ตัวขวามีค่ามากกว่าตัวซ้าย
		} else if encodedSplit[i] == "=" {
			mem[i+1] = mem[i] //ตัวขวามีค่าเท่ากับตัวซ้าย
		} else if encodedSplit[i] != "L" { //ดักตัวอักษรอื่นๆ
			fmt.Println("Invalid Input")
			return ""
		}
	}

	//from right to left handle with L and =
	for j := n - 1; j >= 0; j-- {
		if encodedSplit[j] == "L" && mem[j] <= mem[j+1] { // ตัวซ้ายต้องมากกว่าตัวขวา; ถ้าตัวซ้ายเคยอัพเดทแล้ว(เพราะค่า R) ต้องเทียบว่าฝั่งซ้ายมากกว่าฝั่งขวาหรือยัง
			mem[j] = mem[j+1] + 1
		} else if encodedSplit[j] == "=" {
			mem[j], mem[j+1] = max(mem[j], mem[j+1]), max(mem[j], mem[j+1]) // ตัวซ้ายต้องเท่ากับตัวขวา; ถ้าตัวซ้ายเคยอัพเดทแล้ว(เพราะค่า R) ต้องเทียบว่าฝั่งไหนมีค่ามากกว่ากัน
		}
	}

	return joinIntSliceToString(mem) // แปลงคำตอบเป็น string และส่งกลับไป
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJsonFile(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	var convertData [][]int

	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if err := json.Unmarshal(data, &convertData); err != nil {
		return nil, fmt.Errorf("error convert file: %w", err)
	}

	return convertData, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// คำนวนหาเส้นทางที่มีค่ามากที่สุดในโครงสร้างจากแถวล่างสุดขึ้นบน (Dynamic Programming)
func FindMostValuePath(filePath string) int {
	// อ่านไฟล์ Json
	valuePath, err := ReadJsonFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	// ค้นหาแถวสุดท้าย
	lastRows := len(valuePath) - 1
	if lastRows < 0 {
		fmt.Println("error reading file: path not found")
		return -1
	}

	// ใช้ตัวแปร mem เก็บผลลัพธ์ที่มีค่ามากที่สุด โดยตั้งต้นด้วยข้อมูลจากแถวสุดท้าย
	mem := make([]int, len(valuePath[lastRows]))
	copy(mem, valuePath[lastRows])

	// เริ่มจากแถวรองสุดขึ้นไป คำนวนหาค่าที่มากที่สุดของเส้นทาง
	for i := lastRows - 1; i >= 0; i-- {
		for j, value := range valuePath[i] {
			// อัพเดทค่า mem[j] เป็นค่าผลรวมที่มากที่สุดของทั้งเส้นทาง
			mem[j] = max(value+mem[j], value+mem[j+1])
		}
	}

	// คืนค่าผลลัพธ์ตำแหน่งบนสุด (ผลรวมสุดท้ายของทั้งเส้นทาง)
	return mem[0]
}

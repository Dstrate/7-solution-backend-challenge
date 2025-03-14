package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ใช้สำหรับอ่านไฟล์ที่เก็บข้อมูลชื่อเนื้อ
// ในที่นี้ใช้ไฟล์ CSV สำหรับเก็บชื่อเนื้อแทน DB เพื่อความง่ายในการทดสอบ
func ReadCsvFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var records []string
	for {
		record, err := reader.Read()
		if err != nil {
			break // ออกจากลูปเมื่อ EOF
		}
		records = append(records, record[0]) // เก็บค่าเฉพาะคอลัมน์เดียว
	}

	return records, nil
}

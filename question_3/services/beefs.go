package services

import (
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/Dstrate/7-solution-backend-challenge-3/utils"
)

var (
	beefInstance *BeefService
	once         sync.Once
)

type BeefService struct {
	BeefTerms      []string
	BeefDataSource string
	Beefs          []string
}

// return new beef
func GetNewBeefService() *BeefService {
	once.Do(func() {
		beefInstance = &BeefService{}
		beefInstance.GetBeefTerms()
		beefInstance.SetBeefDataSource("")
	})
	return beefInstance
}

// อ่านไฟล์ beefTerms.csv ใช้เป็นข้อมูลในการคัดกรองเนื้อ
func (b *BeefService) GetBeefTerms() {
	filePath := os.Getenv("BEEF_TERMS_PATH")
	if filePath == "" {
		log.Fatalf("BEEF_TERMS_PATH is not set.")
	}
	beefs, err := utils.ReadCsvFile(filePath)
	if err != nil {
		log.Fatalf("error read csv : %s", err)
	}
	b.BeefTerms = beefs
}

// ตั้งค่า api route สำหรับใช้เป็น datasource
func (b *BeefService) SetBeefDataSource(beefDs string) {
	if beefDs == "" {
		beefDs = os.Getenv("BACON_API")
		if beefDs == "" {
			log.Fatalf("BACON_API is not set.")
		}
	}
	b.BeefDataSource = beefDs
}

// ดึงข้อมูลจาก api และแปลงข้อมูลเก็บไว้
func (b *BeefService) GetBeefs() error {
	// Get Data from Bacon API
	respBody, err := utils.RequestExternalApi(b.BeefDataSource)
	if err != nil {
		return err
	}
	beefData, err := io.ReadAll(respBody)
	if err != nil {
		return err
	}

	// คัดแยกเนื้อออกมาจากข้อความ
	pattern := `(?i)\b(?:` + strings.Join(b.BeefTerms, "|") + `)\b`
	beefSlice := utils.RegexFindAllString(string(beefData), pattern)
	b.Beefs = beefSlice

	return nil
}

// นับจำนวนเนื้อและส่งคืน
func (b *BeefService) CountBeefs() map[string]int64 {
	countBeef := make(map[string]int64)
	for _, beef := range b.Beefs {
		key := strings.ToLower(beef)
		if _, exist := countBeef[key]; !exist {
			countBeef[key] = 0
		}
		countBeef[key]++
	}
	return countBeef
}

func (b *BeefService) GetBeefSummary() (map[string]int64, error) {
	err := b.GetBeefs() // Get Data from Bacon API
	if err != nil {
		return nil, err
	}
	countBeefs := b.CountBeefs() // Count number of beefs
	return countBeefs, nil
}

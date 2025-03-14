package main

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"

	utils "github.com/Dstrate/7-solution-backend-challenge-3/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Dstrate/7-solution-backend-challenge-3/beef/proto"
)

const (
	filePath = "./src/meats.csv"
	beefApi  = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
)

// เก็บคำศัพท์ที่เกี่ยวกับเนื้อ
var beefTerms ProxyBeefData

type ProxyBeefData struct {
	Beefs []string
}

func (b *ProxyBeefData) GetBeefTerms(filePath string) error {
	beefs, err := utils.ReadCsvFile(filePath)
	if err != nil {
		return fmt.Errorf("error read csv : %w", err)
	}
	b.Beefs = beefs
	return nil
}

func RegexBeef(beefData string) []string {
	pattern := `(?i)\b(?:` + strings.Join(beefTerms.Beefs, "|") + `)\b`
	regex := regexp.MustCompile(pattern)
	beefSlice := regex.FindAllString(beefData, -1)
	return beefSlice
}

func CountBeef(beefs []string) map[string]int64 {
	countBeef := make(map[string]int64)
	for _, beef := range beefs {
		key := strings.ToLower(beef)
		if _, exist := countBeef[key]; !exist {
			countBeef[key] = 0
		}
		countBeef[key]++
	}
	return countBeef
}

func BeefSummary() (map[string]int64, error) {
	// ไม่โหลดข้อมูลซ้ำ
	if beefTerms.Beefs == nil {
		err := beefTerms.GetBeefTerms(filePath)
		if err != nil {
			return nil, err
		}
	}

	// Get ข้อมูลเนื้อผ่าน api
	respBody, err := utils.RequestExternalApi(beefApi)
	if err != nil {
		return nil, err
	}
	beefData, err := io.ReadAll(respBody)
	if err != nil {
		return nil, err
	}

	// คัดแยกข้อมูลเนื้อออกมาจาก beefData
	beefSlice := RegexBeef(string(beefData))

	// นับจำนวนเนื้อ
	beefs := CountBeef(beefSlice)
	return beefs, nil
}

func (s *server) BeefSummaryService(ctx context.Context, req *pb.BeefSummaryRequest) (*pb.BeefSummaryResponse, error) {
	beefs, err := BeefSummary()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BeefSummaryResponse{
		Beef: beefs,
	}, nil
}

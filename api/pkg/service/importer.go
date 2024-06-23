package service

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/Sigma-Ratings/sigma-code-challenges/api/config"
	"github.com/Sigma-Ratings/sigma-code-challenges/api/pkg/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate mockery --name DB
type SanctionsImport struct {
	DB *gorm.DB
}

func (s *SanctionsImport) ImportSanctions() {
	logrus.Info("starting bootstrapping data")
	resp, errFetch := fetchSanctions()
	if errFetch != nil {
		logrus.WithError(errFetch).Error("could not fetch sanctions data")
	}
	csvData, errRead := getCSVFromResponse(resp)
	if errRead != nil {
		logrus.WithError(errRead).Error("could not read csv")
	}
	err := s.saveSanctionsData(csvData)
	if err != nil {
		logrus.WithError(err).Error("could not save sanctions data")
	}
	logrus.Info("finished bootstrapping data")
}

func fetchSanctions() (*http.Response, error) {
	resp, err := http.Get(config.App.SanctionFileURL)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getCSVFromResponse(resp *http.Response) ([][]string, error) {
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1
	reader.Comma = ';'
	reader.LazyQuotes = true
	values, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (s *SanctionsImport) saveSanctionsData(csv [][]string) error {
	for index, rec := range csv {
		if index != 0 && rec[17] != "" {
			logicalId, err := strconv.Atoi(rec[1])
			if err != nil {
				return err
			}
			alias := rec[17]
			sr := model.SanctionEntity{
				LogicalID: logicalId,
				Alias:     alias,
			}
			if err = s.DB.Save(&sr).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

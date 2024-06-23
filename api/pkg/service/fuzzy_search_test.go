package service

import (
	"testing"

	"github.com/Sigma-Ratings/sigma-code-challenges/api/config"
	"github.com/Sigma-Ratings/sigma-code-challenges/api/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

func setup() (testDb *gorm.DB) {
	testDb, err := db.GetConnection(config.TestDatabase())
	if err != nil {
		logrus.WithError(err).Fatal("could not connect to database") // Fatal will exit test
	}

	return testDb
}

type aliasesInput struct {
	LogicalId    int
	excludeAlias string
}

func TestGetAlises(t *testing.T) {
	type successTestCases struct {
		description    string
		input          aliasesInput
		expectedReturn []string
	}

	db := setup()
	ss := SanctionsImport{DB: db}

	for _, scenario := range []successTestCases{
		{
			description: "return aliases",
			input: aliasesInput{
				LogicalId:    13,
				excludeAlias: "Abu Ali",
			},
			expectedReturn: []string{"Saddam Hussein Al-Tikriti", "Abou Ali"},
		},
		{
			description: "return aliases",
			input: aliasesInput{
				LogicalId:    46,
				excludeAlias: "Muhammad Hamza Zubaidi",
			},
			expectedReturn: []string{"Mohammed Hamza Zoubaïdi"},
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			result, _ := ss.GetAliases(scenario.input.LogicalId, scenario.input.excludeAlias)
			assert.Equal(t, scenario.expectedReturn, result)
		})
	}

}

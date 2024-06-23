package controller

import (
	"net/http"

	"github.com/Sigma-Ratings/sigma-code-challenges/api/pkg/model"
	"github.com/Sigma-Ratings/sigma-code-challenges/api/pkg/service"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

// since the file is fixed and we don't change it is safe that this number will not change
const totalNumberOfEntities = 15462

func (c *Controller) Status(context echo.Context) (err error) {
	var count int64
	if errCount := c.DB.Model(&model.SanctionEntity{}).Count(&count).Error; errCount != nil {
		return context.JSON(http.StatusInternalServerError, errCount)
	}

	// checks if full list was saved
	if count == totalNumberOfEntities {
		return context.NoContent(http.StatusOK)
	}

	return context.NoContent(http.StatusServiceUnavailable)
}

// Return model for the search endpoint, your search must not contain multiple results with the same logical id
type Sanction struct {
	LogicalId     int      `json:"logicalId"`
	MatchingAlias string   `json:"matchingAlias"`
	OtherAliases  []string `json:"otherAliases"`
	Exact         bool     `json:"exact"`     // true if relevance 1, otherwise false
	Relevance     float64  `json:"relevance"` //between 0 and 1, 1 means the query and the match name is equal
}

func (c *Controller) Search(context echo.Context) (err error) {
	searchName := context.QueryParam("name")

	ss := service.SanctionsImport{DB: c.DB}

	var topAliasMatches, errFuzzyMatch = ss.FuzzyNameMatches(searchName)
	if topAliasMatches == nil && errFuzzyMatch == nil {
		return context.NoContent(http.StatusNoContent)
	}
	if errFuzzyMatch != nil {
		logrus.WithError(errFuzzyMatch).Error("Error querying fuzzy match algorithms.")
		return context.NoContent(http.StatusInternalServerError)
	}

	var mostRelevantMatch = topAliasMatches[0]
	var filteredMatches map[int]int
	filteredMatches = make(map[int]int)

	var results []Sanction
	for idx, match := range topAliasMatches {
		if _, ok := filteredMatches[match.LogicalId]; ok {
			continue
		}
		filteredMatches[match.LogicalId] = idx

		var otherAliases, errGetAliases = ss.GetAliases(match.LogicalId, mostRelevantMatch.Alias)
		if errGetAliases != nil {
			logrus.WithError(errGetAliases).Error("Error getting aliases.")
			return context.NoContent(http.StatusInternalServerError)
		}

		var sanction = Sanction{
			LogicalId:     match.LogicalId,
			MatchingAlias: match.Alias,
			OtherAliases:  otherAliases,
			Relevance:     match.LevDist,
			Exact:         checkExactMatch(match.LevDist),
		}

		results = append(results, sanction)
	}

	return context.JSON(http.StatusOK, results)
}

func checkExactMatch(normalizedLevDist float64) bool {
	if normalizedLevDist == 1 {
		return true
	}

	return false
}

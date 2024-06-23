package service

import (
	"github.com/Sigma-Ratings/sigma-code-challenges/api/pkg/model"
	"gorm.io/gorm"
)

type SanctionsDistancesQuery struct {
	Alias     string
	LogicalId int
	Id        int
	LevDist   float64
}

func (s *SanctionsImport) FuzzyNameMatches(searchName string) ([]SanctionsDistancesQuery, error) {
	var topAliasMatches []SanctionsDistancesQuery
	var maxLevenshteinDist = len([]rune(searchName)) / 3
	if result := s.DB.Raw(
		`SELECT
			alias as Alias,
			logical_id,
			id as Id,
			(
				(
					GREATEST(length(alias), length(@search)) - levenshtein(alias, @search)
				)::float8 / (GREATEST(length(alias), length(@search)))::float8
			)::float8  as lev_dist

		FROM sanction_entities
		WHERE
			levenshtein(dmetaphone(alias), dmetaphone(@search)) < 1
		AND levenshtein(alias, @search) < @maxDist
		GROUP BY logical_id, alias, id
		ORDER BY lev_dist DESC;`,
		map[string]interface{}{"search": searchName, "maxDist": maxLevenshteinDist}).Scan(
		&topAliasMatches); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return topAliasMatches, nil
}

func (s *SanctionsImport) GetAliases(logicalId int, excludeAlias string) ([]string, error) {
	var allAliases []model.SanctionEntity
	result := s.DB.Find(&allAliases, "logical_id = ? AND alias <> ?", logicalId, excludeAlias)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	var otherAliases []string
	for _, alias := range allAliases {
		otherAliases = append(otherAliases, alias.Alias)
	}

	return otherAliases, nil
}

package maxcdn

import (
	"fmt"
	"net/url"
)

// Get does an OAuth signed http.Get for "/reports/popularfiles.json"
func (max *MaxCDN) GetPopularFiles(form url.Values) (mapper *PopularFiles, err error) {
	mapper = new(PopularFiles)
	raw, res, err := max.Do("GET", PopularFilesEndpoint, form)
	mapper.Response = res
	if err != nil {
		return
	}

	err = mapper.Parse(raw)
	return
}

// Get does an OAuth signed http.Get for "/reports/stats.json"
func (max *MaxCDN) GetStatsSummary(form url.Values) (mapper *SummaryStats, err error) {
	mapper = new(SummaryStats)
	raw, res, err := max.Do("GET", StatsEndpoint, form)
	mapper.Response = res
	if err != nil {
		return
	}

	err = mapper.Parse(raw)
	return
}

// GetHourlyStats does an OAuth signed http.Get for "/reports/stats.json/{report_type}".
//
// Valid report types are; 'hourly', 'daily' and 'monthly'
func (max *MaxCDN) GetStatsByType(report string, form url.Values) (mapper *MultiStats, err error) {
	mapper = new(MultiStats)
	endpoint := fmt.Sprintf("%s/%s", StatsEndpoint, report)
	raw, res, err := max.Do("GET", endpoint, form)
	mapper.Response = res
	if err != nil {
		return
	}

	err = mapper.Parse(raw)
	return
}

// GetHourlyStats does an OAuth signed http.Get for "/reports/stats.json/hourly"
func (max *MaxCDN) GetHourlyStats(form url.Values) (mapper *MultiStats, err error) {
	return max.GetStatsByType("hourly", form)
}

// GetDailyStats does an OAuth signed http.Get for "/reports/stats.json/hourly"
func (max *MaxCDN) GetDailyStats(form url.Values) (mapper *MultiStats, err error) {
	return max.GetStatsByType("daily", form)
}

// GetMonthlyStats does an OAuth signed http.Get for "/reports/stats.json/hourly"
func (max *MaxCDN) GetMonthlyStats(form url.Values) (mapper *MultiStats, err error) {
	return max.GetStatsByType("daily", form)
}

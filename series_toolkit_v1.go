package fred_go_toolkit

// Series is a single instance of a FRED series.
type Series struct {
	ID                      string `json:"id" xml:"id,attr"`
	Start                   string `json:"realtime_start" xml:"realtime_start,attr"`
	End                     string `json:"realtime_end" xml:"realtime_end,attr"`
	Title                   string `json:"title" xml:"title,attr"`
	ObsStart                string `json:"observation_start" xml:"observation_start,attr"`
	ObsEnd                  string `json:"observation_end" xml:"observation_end,attr"`
	Frequency               string `json:"annual" xml:"annual,attr"`
	FrequencyShort          string `json:"frequency_short" xml:"frequency_short,attr"`
	Units                   string `json:"units" xml:"units,attr"`
	UnitsShort              string `json:"units_short" xml:"units_short,attr"`
	SeasonalAdjustment      string `json:"seasonal_adjustment" xml:"seasonal_adjustment,attr"`
	SeasonalAdjustmentShort string `json:"seasonal_adjustment_short" xml:"seasonal_adjustment_short,attr"`
	LastUpdated             string `json:"last_updated" xml:"last_updated,attr"`
	Popularity              int    `json:"popularity" xml:"popularity,attr"`
	GroupPopularity         int    `json:"group_popularity" xml:"group_popularity,attr"`
	Notes                   string `json:"notes" xml:"notes,attr"`
}

// Observation is a single instance of a FRED observation.
type Observation struct {
	Start string `json:"realtime_start" xml:"realtime_start,attr"`
	End   string `json:"realtime_end" xml:"realtime_end,attr"`
	Date  string `json:"date" xml:"date,attr"`
	Value string `json:"value" xml:"value,attr"`
}

// GetSeries will get an economic data series.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series.html
func (f *FredClient) GetSeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesParam)

	if err != nil {
		f.logError(seriesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSeriesCategories will get the categories for an economic data series.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_categories.html
func (f *FredClient) GetSeriesCategories(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesCategoriesParam)

	if err != nil {
		f.logError(seriesCategoriesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetReleaseObservations will get the observations or data values for an economic data series.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_observations.html
func (f *FredClient) GetSeriesObservations(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesObservationsParam)

	if err != nil {
		f.logError(seriesObservationsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSeriesRelease will get the release for an economic data series.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_release.html
func (f *FredClient) GetSeriesRelease(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesReleaseParam)

	if err != nil {
		f.logError(seriesReleaseParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSeriesSearch will get economic data series that match keywords.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_search.html
func (f *FredClient) GetSeriesSearch(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesSearchParam)

	if err != nil {
		f.logError(seriesSearchParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSeriesSearchTags will get the FRED tags for a series search.
// Optionally, filter results by tag name, tag group, or tag search.
// See the related request GetSeriesSearchRelatedTags.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_search_tags.html
func (f *FredClient) GetSeriesSearchTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesSearchTagsParam)

	if err != nil {
		f.logError(seriesSearchTagsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSeriesSearchRelatedTags will get teh related tags for a series search.
// Optionally, filter results by tag group or tag search.
//
// FRED tags are attributes assigned to series.
// For this request, related FRED tags are the tags assigned to series that match all tags in the tag_names parameter, no tags in the exclude_tag_names parameter, and the search words set by the series_search_text parameter.
// See the related request GetSeriesSearchTags.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_search_related_tags.html
func (f *FredClient) GetSeriesSearchRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesSearchRelatedTagsParam)

	if err != nil {
		f.logError(seriesSearchRelatedTagsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSeriesTags will get the tags for an economic data series.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_tags.html
func (f *FredClient) GetSeriesTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesTagsParam)

	if err != nil {
		f.logError(seriesTagsParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSeriesUpdates will get economic data series sorted by when observations were updated on the FRED® server (attribute last_updated).
// Results are limited to series updated within the last two weeks.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_updates.html
func (f *FredClient) GetSeriesUpdates(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesUpdatesParam)

	if err != nil {
		f.logError(seriesUpdatesParam, err)
		return nil, err
	}

	return fc, nil

}

// GetSeriesVintageDates will get the dates in history when a series' data values were revised or new data values were released. Vintage dates are the release dates for a series excluding release dates when the data for the series did not change.
//
// Schema for the request and response objects and source for the documentation can be found at the following link: https://research.stlouisfed.org/docs/api/fred/series_vintagedates.html
func (f *FredClient) GetSeriesVintageDates(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, seriesVintagedatesParam)

	if err != nil {
		f.logError(seriesVintagedatesParam, err)
		return nil, err
	}

	return fc, nil

}

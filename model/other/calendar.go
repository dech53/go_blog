package other

type Calendar struct {
	Date         string `json:"date"`
	LunarDate    string `json:"lunar_date"`
	Ganzhi       string `json:"ganzhi"`
	Zodiac       string `json:"zodiac"`
	DayOfYear    string `json:"day_of_year"`
	SolarTerm    string `json:"solar_term"`
	Auspicious   string `json:"auspicious"`
	Inauspicious string `json:"inauspicious"`
}

package other

type IPResponse struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	InfoCode  string `json:"infocode"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
}

type Cast struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	DayWeather   string `json:"dayweather"`
	NightWeather string `json:"nightweather"`
	DayTemp      string `json:"daytemp"`
	NightTemp    string `json:"nighttemp"`
	DayWind      string `json:"daywind"`
	NightWind    string `json:"nightwind"`
	DayPower     string `json:"daypower"`
	NightPower   string `json:"nightpower"`
}

type Live struct {
	Province         string `json:"province"`
	City             string `json:"city"`
	Adcode           string `json:"adcode"`
	Weather          string `json:"weather"`
	Temperature      string `json:"temperature"`
	WindDirection    string `json:"winddirection"`
	WindPower        string `json:"windpower"`
	Humidity         string `json:"humidity"`
	ReportTime       string `json:"reporttime"`
	TemperatureFloat string `json:"temperature_float"`
	HumidityFloat    string `json:"humidity_float"`
}

type Forecast struct {
	City       string `json:"city"`
	Adcode     string `json:"adcode"`
	Province   string `json:"province"`
	ReportTime string `json:"reporttime"`
	Casts      []Cast `json:"casts"`
}

type WeatherResponse struct {
	Status   string   `json:"status"`
	Count    string   `json:"count"`
	Info     string   `json:"info"`
	InfoCode string   `json:"infocode"`
	Lives    []Live   `json:"lives"`
	Forecast Forecast `json:"forecast"`
}

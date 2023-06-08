package bitgethelpers

type bitgetServerTimeStampResponse struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int    `json:"requestTime"`
	Data        string `json:"data"`
}

type AccountResponse struct {
	Code        string        `json:"code"`
	Message     string        `json:"msg"`
	RequestTime int64         `json:"requestTime"`
	Data        []AccountData `json:"data"`
}
type bybitServerTimeStampResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     map[string]interface{} `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type AccountData struct {
	MarginCoin         string `json:"marginCoin"`
	AvailableBalance   string `json:"available"`
	TotalMarginBalance string `json:"equity"`
	MarginValueUSDT    string `json:"usdtEquity"`
	MarginValueBTC     string `json:"btcEquity"`
	FloatingPnl        string `json:"unrealizedPL"`
}

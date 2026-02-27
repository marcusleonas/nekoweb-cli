package nekoweb

type Limits struct {
	General    LimitInfo
	BigUploads LimitInfo
	Zip        LimitInfo
}

type LimitInfo struct {
	Limit     int   `json:"limit"`
	Remaining int   `json:"remaining"`
	Reset     int64 `json:"reset"`
}

func (c *Client) GetLimits() (Limits, error) {
	req, err := c.newRequest("GET", "/files/limits", nil)
	if err != nil {
		return Limits{}, err
	}

	var result Limits
	if err := c.do(req, &result); err != nil {
		return Limits{}, err
	}

	return result, nil
}

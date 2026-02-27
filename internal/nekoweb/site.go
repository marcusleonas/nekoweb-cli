package nekoweb

import "fmt"

type SiteInfo struct {
	Domain    string `json:"domain"`
	Title     string `json:"title"`
	Updates   int    `json:"updates"`
	Followers int    `json:"followers"`
	Views     int    `json:"views"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (c *Client) GetSiteInfo(domain string) (*SiteInfo, error) {
	reqPath := fmt.Sprintf("/site/info/%v", domain)

	req, err := c.newRequest("GET", reqPath, nil)
	if err != nil {
		return nil, err
	}

	var result SiteInfo
	if err := c.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAllSitesInfo() (*[]SiteInfo, error) {
	req, err := c.newRequest("GET", "/site/info_all", nil)
	if err != nil {
		return nil, err
	}

	var result []SiteInfo
	if err := c.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

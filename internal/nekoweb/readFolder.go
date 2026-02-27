package nekoweb

import "fmt"

type FolderItem struct {
	Name string `json:"name"`
	Dir  bool   `json:"dir"`
}

func (c *Client) ReadFolder(path string) ([]FolderItem, error) {
	reqPath := fmt.Sprintf("/files/readfolder?pathname=%v", path)

	req, err := c.newRequest("GET", reqPath, nil)
	if err != nil {
		return nil, err
	}

	var result []FolderItem
	if err := c.do(req, &result); err != nil {
		return nil, err
	}

	return result, nil
}

package harbor

import "github.com/go-resty/resty/v2"

//HarborClient a harbor client
type HarborClient struct {
	Client *resty.Client
}

//NewHarborClient new a harbor client
func NewHarborClient(host, userName, passWord string) *HarborClient {
	client := resty.New()
	client.HostURL = host
	client.SetBasicAuth(userName, passWord)
	return &HarborClient{Client: client}
}

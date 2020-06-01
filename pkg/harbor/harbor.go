package harbor

import "github.com/go-resty/resty/v2"

//HarborClient a harbor clientv2
type HarborClient struct {
	Client *resty.Client
}

//NewHarborClient new a harbor clientv2
func NewHarborClient(host, userName, passWord string) *HarborClient {
	client := resty.New()
	client.HostURL = host
	client.SetBasicAuth(userName, passWord)
	return &HarborClient{Client: client}
}

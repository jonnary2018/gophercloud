package profiles

import "github.com/gophercloud/gophercloud"

var apiVersion = "v1"
var apiName = "profiles"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func createURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}

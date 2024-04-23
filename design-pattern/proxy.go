package main

import (
	"errors"
	"fmt"
)

type INetworkClient interface {
	call(url string) (string, error)
}
type NetworkClient struct{}
type ProxyNetworkClient struct {
	networkClient INetworkClient
	blackUrls     []string
}

func (*NetworkClient) call(url string) (string, error) {
	response := "response from " + url
	return response, nil
}
func (proxy *ProxyNetworkClient) addBlackUrl(url string) {
	proxy.blackUrls = append(proxy.blackUrls, url)
}
func (proxy *ProxyNetworkClient) call(dst string) (string, error) {
	for _, url := range proxy.blackUrls {
		if dst == url {
			return "", errors.New(url + " is in blacklist")
		}
	}
	response := "response from " + dst
	return response, nil
}
func main() {
	client := NetworkClient{}
	call, err := client.call("google.com")
	fmt.Print("call from client: ")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(call)
	}

	proxy := ProxyNetworkClient{}
	proxy.addBlackUrl("google.com")
	call, err = proxy.call("google.com")
	fmt.Print("call from proxy: ")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(call)
	}
}

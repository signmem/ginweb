package http

import "fmt"

type UserInfo struct {
	Name		string		`json:"name"`
	Domain		[]string	`json:"domain"`
	Token		string		`json:"token"`
}

func (this *UserInfo) String() string {
	return fmt.Sprintf("user: %s, domain: %s, token: %s", this.Name, this.Domain, this.Token)
}

package structs

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (info authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", info.username, info.password)
}

func UsingStructMethods() {
	authInfo := authenticationInfo{
		username: "ajani",
		password: "password",
	}

	println(authInfo.getBasicAuth())
}

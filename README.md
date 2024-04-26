# authsignal-management-go

This Go SDK is a wrapper for Authsignal's Management API and is used by Authsignal's Terraform provider.

## Installation

```
go get github.com/authsignal/authsignal-management-go
```

## Example Usage

```
package main

import (
	"encoding/json"
	"fmt"

	"github.com/authsignal/authsignal-management-go"
)

const authsignalSecret = "<management api secret retrieved from the admin portal>"
// The URI below will depend on which region your tenant is located in.
const authsignalURI = "https://api.authsignal.com/v1/management"
const authsignalTenantId = "aaaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"

func main() {
	var actionConfiguration authsignal.ActionConfiguration
	actionConfiguration.ActionCode = "helloworld"

	var client authsignal.Client = authsignal.NewClient(authsignalURI, authsignalTenantId, authsignalSecret)
	var actionConfigurationResponse, err = client.GetActionConfiguration(actionConfiguration.ActionCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", actionConfigurationResponse)
}
```

## Available methods

```
getActionConfigurationHttp()
createActionConfigurationHttp()
updateActionConfigurationHttp()
deleteActionConfigurationHttp()

getRuleHttp()
createRuleHttp()
updateRuleHttp()
deleteRuleHttp()
```

## Documentation

Check out our [official documentation](https://docs.authsignal.com/category/api-reference) to get up and running quickly.

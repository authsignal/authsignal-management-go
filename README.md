# authsignal-management-go

This Go SDK is a wrapper for Authsignal's Management API and is used by Authsignal's Terraform provider.

## Installation

```
go get github.com/authsignal/authsignal-management-go/v2
```

## Example Usage

```
package main

import (
	"fmt"

	"github.com/authsignal/authsignal-management-go/v2"
)

const authsignalSecret = "<management api secret retrieved from the admin portal>"
// The URI below will depend on which region your tenant is located in.
const authsignalURI = "https://api.authsignal.com/v1/management"
const authsignalTenantId = "aaaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"

func main() {
	actionCode := "helloworld"
	var actionConfiguration authsignal.ActionConfiguration
	actionConfiguration.ActionCode = authsignal.SetValue(actionCode)

	var client authsignal.Client = authsignal.NewClient(authsignalURI, authsignalTenantId, authsignalSecret)
	var actionConfigurationResponse, err = client.GetActionConfiguration(actionCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", actionConfigurationResponse)
}
```

## Available methods

```
SetValue()
SetNull()

GetActionConfigurationHttp()
CreateActionConfigurationHttp()
UpdateActionConfigurationHttp()
DeleteActionConfigurationHttp()

GetRuleHttp()
CreateRuleHttp()
UpdateRuleHttp()
DeleteRuleHttp()
```

## Documentation

Check out our [official documentation](https://docs.authsignal.com/category/api-reference) to get up and running quickly.

// Copyright 2018 Microsoft Corporation
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"fmt"
	"os"

	auth "github.com/Azure/azure-auth-go"
)

func main() {
	authConfig := auth.AzureAuthConfig{
		Cloud:                       "AzurePublicCloud",
		UseManagedIdentityExtension: true,
	}
	servicePrincipalToken, err := auth.GetAzureServicePrincipalToken(&authConfig)
	if err != nil {
		fmt.Printf("Failed to get Azure service principal token: %v\n", err)
		os.Exit(1)
	}

	if err = servicePrincipalToken.EnsureFresh(); err != nil {
		fmt.Printf("Failed to refresh token: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got new token: %s\n", servicePrincipalToken.Token())
}

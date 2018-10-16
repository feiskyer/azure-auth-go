# azure-auth-go

Package azure-auth-go provides a library for authorizing with Azure.

## Usage

```go

import (
    auth "github.com/Azure/azure-auth-go"
)

authConfig = auth.AzureAuthConfig{
    Cloud: "AzurePublicCloud",
    UseManagedIdentityExtension: true,
}
servicePrincipalToken, err := auth.GetAzureServicePrincipalToken(&authConfig)
if err != nil {
    // Handle error
}

// continue with other logics
```

## License

See [LICENSE](LICENSE) file.


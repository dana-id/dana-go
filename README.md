# Go API client for dana

The official DANA Go SDK provides a simple and convenient way to call DANA's REST API in applications written in Go (based on https://dashboard.dana.id/api-docs-v2/)

## ⚠️ Run This First - Save Days of Debugging

Before writing any integration code, **run our automated test suite**. It takes **under 2 minutes** and shows you how the full flow works — **with your own credentials**.

Here is the link: https://github.com/dana-id/uat-script.

### Why This Matters

- 🧪 Validates your setup instantly
- 👀 **See exactly how each scenario flows**
- 🧾 Gives us logs to help you faster
- 🚫 Skipping this = guaranteed delays 


### What It Does

✅ Runs full scenario checks for DANA Sandbox

✅ Installs and executes automatically

✅ Shows real-time results in your terminal

✅ Runs in a safe, simulation-only environment

> Don't fly blind. Run the test first. See the flow. Build with confidence.

  
  .  

  .


## Installation

Install the following dependencies:

```sh
go get github.com/dana-id/dana-go
```

## Environment Variables

Before using the SDK, please make sure to set the following environment variables (In .env):

| Name                   | Description                                                                                   | Example Value                                                                   |
| ---------------------- | ---------------------------------------------------------------------------------------       | ------------------------------------------------------------------------------- |
| `ENV` or `DANA_ENV`    | Defines which environment the SDK will use. Possible values: `SANDBOX` or `PRODUCTION`.       | `SANDBOX`                                                                       |
| `X_PARTNER_ID`         | Unique identifier for partner, provided by DANA, also known as `clientId`.                    | 1970010100000000000000                                                          |
| `PRIVATE_KEY`          | Your private key string.                                                                      | `-----BEGIN PRIVATE KEY-----MIIBVgIBADANBg...LsvTqw==-----END PRIVATE KEY-----` |
| `PRIVATE_KEY_PATH`     | Path to your private key file. If both are set, `PRIVATE_KEY_PATH` is used.                   | /path/to/your_private_key.pem                                                   |
| `DANA_PUBLIC_KEY`      | DANA public key string for parsing webhook.                                                   | `-----BEGIN PUBLIC KEY-----MIIBIjANBgkq...Do/QIDAQAB-----END PUBLIC KEY-----`   |
| `DANA_PUBLIC_KEY_PATH` | Path to DANA public key file for parsing webhook. If both set, `DANA_PUBLIC_KEY_PATH is used. | /path/to/dana_public_key.pem                                                    |
| `ORIGIN`               | Origin domain.                                                                                | https://yourdomain.com                                                          |
| `CLIENT_SECRET`        | Assigned client secret during registration. Must be set for DisbursementApi                   | your_client_secret                                                              |
| `X_DEBUG`              | Debug mode. Set to `true` to enable debug mode (will show reason of failed request in additionalInfo.debugMessage in response).                                               | true                                                                            |

You can see these variables in .env.example, fill it, and change the file name to .env (remove the .example extension)

Then you can choose these following APIs based on the business solution you want to integrate:

## Documentation for API Endpoints

API | Description
------------- | -------------
[**PaymentGatewayApi**](docs/PaymentGatewayAPI.md) | API for doing operations in DANA Payment Gateway (Gapura)
[**WidgetApi**](docs/WidgetAPI.md) | API for enabling the user to make payment from merchant’s platform with redirecting to DANA’s platform
[**DisbursementApi**](docs/DisbursementAPI.md) | API for enabling the user to disburse money to merchant’s platform
[**MerchantManagementApi**](docs/MerchantManagementAPI.md) | API for managing merchant resources and shops


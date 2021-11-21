# OPE API Documentation

[Introduction](#introduction) | [Register](#register) | [Profile Update](#to-update-user-profile) | [Bank Details](#update-bank-details) | [Login](#login) | [Wallet Funding](#fund-user-wallet-flow) | [Investments](#investment) | [Loan Management](#loan-management-api) | [Register API](#register-api) | [Profile Update API](#profile-update-api) | [Update Bank Details API](#update-bank-details-api) | [Login API](#login-api)| [Complete-Login API](#complete-login-api) |[Login Sample Response](#login-sample-response) | [Funding API](#funding-api) | [Payment Response](#payment-response) | [Complete Funding API](#complete-funding-api) | [Investment API](#investment-api) | [Request Loan API](#request-loan-api) | [Fetch Loans API](#fetch-loans-api) | [Repayment Api](#repayment-api)

## Introduction

This is the backend infrastructure build for [Ope App](https://loaner-two.vercel.app/), a stock portfolio web application with loan management features.

To interract with the application core resources, a user is required to register and verified their account. After registering, the user can login to the application using their registered credentials. This will provide the user an hour long authorization token that can be used to access the API core resorces.

### REGISTER

To register a user, the client should make a request as detailed below. AFter complete registration, users are sent verification link via email which can be used to verify user's account by a click.

### Register API

```

- RGISTER API
  URL : https://be-ope.herokuapp.com/register
  METHODE TYPE: Post
  PAYLOAD:
  {
     "firstname": "User first name",
     "lastname" :"last name",
     "email":"email formatted",
     "phone":"phone nuber as string",
     "password":"string",
  }

```

### TO UPDATE USER PROFILE

Client should send request as follows

### Profile Update API

```

URL : https://be-ope.herokuapp.com/update-profile/{userId}
METHODE TYPE: Patch
Authorization: Bearer token (Valid user token that is not PAYMENT TOKEN)
PAYLOAD:
{
    "firstname":"string",
    "lastname":"strin",
    "email": "string",
    "phone": "string",
    "account_no": "string",
    "bank_name": "Zenith Bank"
}

```

### UPDATE BANK DETAILS

Users can decide to add their bank details to ease wallet funds withdrawer. To do this, clent should send request as below

### Update Bank Details API

```

URL: https://be-ope.herokuapp.com/user/bankupdate/{userId}
METHOD TYPE: Patch
Authorization: Bearer token (Valid user token that is not PAYMENT TOKEN)
PAYLOAD:
{
    "bank_name":"Zenith Bank",
    "account_no":"2085xxxxxx"
}

```

### LOGIN

The application login flow include a miniature 2FA that involves an OTP being sent and used to conclude the login flow. Hence, the login process involves the client sending request to, first, Login Endpoint, and then, Complete-Login endpoints.

To initiate a Loin process, the client should request as below

### Login API

```

URL : https://be-ope.herokuapp.com/login
METHODE TYPE: Post
PAYLOAD:
{
    "email":"string",
    "password":"string"
}

SAMPLE RESPONSE:

{
    "status": "success",
    "collection": "Login Token",
    "data": {
        "message": "Check Mail For Login OTP",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYW5rX2FjY291bnQiOnsi ..."
    }
}

```

This will trigger an OTP emailing to the user. the user can then use the otp to complete login as follows

### Complete Login API

```

URL : https://be-ope.herokuapp.com/complete-login
METHODE TYPE: Post
AUTHORIZATION: Bearer token (Token recieved for login)
PAYLOAD:
{
    "otp":numeric otp,
}

```

Upon successfull Login, client should get details of users activities as shown in SAMPLE RESPONSE below.

### Login Sample Response

```

{
    "status": "success",
    "collection": "User",
    "data": {
        "user_id": "5",
        "firstname": "Ogie",
        "lastname": "xxxxx",
        "email": "gulephil44@gmail.com",
        "created_at": "2021-11-15 00:52:42",
        "Holdings": "",
        "bank_account": {
            "account_no": "Zenith Bank",
            "bank_name": "208xxxxxxx"
        },
        "otp": "803634",
        "wallet": {
            "amount": 34046.758,
            "address": "$2a$06$WRKRFEsAP/meZbjMP1lkOuzyu7jtZ66cu8uH0dQZPKP3pwzDwYRvi"
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYW5rX2FjY291bnQiOnsi...",
        "portfolio": [
            {
                "id": "5",
                "symbol": "TM",
                "image_url": "https://st2.depositphotos.com/5943796/11433/v/600/gold.jpg",
                "quantity": "241",
                "unit_price": "179.28",
                "equity_value": "17553.932",
                "percentage_change": "5.84"
            },
            {
                "id": "15",
                "symbol": "TYO",
                "image_url": "https://global.toyota/pages/global_toyota/mobility/toyota-brand/emblem_ogp_001.png",
                "quantity": "31",
                "unit_price": "17.94",
                "equity_value": "879.33",
                "percentage_change": "-3.39"
            },
            {
                "id": "45",
                "symbol": "FB",
                "image_url": "https://cdn3.vectorstock.com/i/1000x1000/02/37/logo-facebook-vector-31060237.jpg",
                "quantity": "2",
                "unit_price": "341.13",
                "equity_value": "682.26",
                "percentage_change": "1.57"
            }
        ]
    }
}

```

NB:

- The portfolio array will be empty if user is yet to buy investment in stock

- The bank attribute will read default state, "none" for both account_no and bank_name if the user is yet to update her bank detals

### FUND USER WALLET FLOW

The process of funding a user's wallet include a series of processes. These processes are categorized nto two stages, viz.

Stage ONE

1 - Call FUND WALLET EndPoint

2 - Backend will respond with PAYMENT-BODY and PAYMENT-TOKEN, see PAYMENT RESPONSE below.

3 - Call FLutterwave gateway (Preferably Flutterwave inline), supply the "payment_body" sent in step-2 as payload, and proceed as will be prompted by Flutterwave

Stage TWO

1- Call COMPLETE-FUNDING EndPoint and send Back PAYMENT_TOKEN sent earlier

#### DETAILS:

### Funding API

```

FUND-WALLET ENDPOINT : "https://be-ope.herokuapp.com/fund-wallet"
METHOD TYPE: POST
Authorization: login token or any active non-payment token
Payload:
{
    "amount":"18000",
    "currency":"NGN",
    "payment_option":"card"
}

```

### PAYMENT RESPONSE

```

{
    "status": "success",
    "collection": "FultterWave",
    "data": {
        "payment_body": {
            "tx_ref": "Ogie-tx-228544",
            "amount": "18000",
            "currency": "NGN",
            "payment_option": "card",
            "redirect_url": "",
            "meta": {
                "consumer_id": 5,
                "consumer_mac": "$2a$06$WRKRFEsAP/meZbjMP1lkOuzyu7jtZ66cu8uH0dQZPKP3pwzDwYRvi"
            },
            "customer": {
                "email": "gulephil44@gmail.com",
                "phonenumber": "",
                "name": "Ogie"
            },
            "customizations": {
                "title": "Fund Wallet",
                "description": "Funding wallet for subsequent trasaction",
                "logo": "www.mylogo.com"
            }
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbW91bnQiOiIxODAwMCIsImiwiY ..."
    }
}


```

### Complete Funding API

```
COMPLETE-FUNDING ENDPOINT : "https://be-ope.herokuapp.com/complete-funding"
METHOD TYPE: PATCH
Authorization: Payment-token (You can only use payment token for this)
Payload:
     {
    "tx_ref": "James-tx-724701",
    "amount": "18000",
    "wallet": "$2a$06$WRKRFEsAP/meZbjMP1lkOuzyu7jtZ66cu8uH0dQZPKP3pwzDwYRvi"
}


```

### NB:

1 - tx-ref is transaction reference. It will be the first attribute of the PAYMENT_BODY sent to you when you call FUND_WALLET endpoint

2 - Amount has to be the same as the one in the PAYMENT BODY

3 - wallet is the user's wallet address. It will also be sent as "customer_mac" under the "meta" attribute of the PAYMENT-BODY

### INVESTMENT

Once a user's wallet has been successfully funded, it can then be used to buy investments in company stocks, Repay loans or withdrawn to user's registered bank account. The amount of stocks that can be bought by user is only limited by the user and the amount available in his/her wallet.

To buy investment, the client should call the endpoint as detailed below.

### Investment Payload

```

URL : https://be-ope.herokuapp.com/buy-stock/{userId}
METHODE TYPE: POST
Authorization: Bearer token (Any valid TOKEN of user that is not PAYMENT TOKEN)
PAYLOAD:
{
    "symbol":"FB",
	"image_url":"https://cdn3.vectorstock.com/i/1000x1000/02/37/logo-facebook-vector-31060237.jpg",
	"quantity" : "2",
	"unit_price" : 341.13,
	"percentage_change": 1.57
}



```

## LOAN MANAGEMENT API

The loan management API is categorized into two, LOANS and loan REPAYMENT API.

For context, the user can take a loan as long as he/she doesn't have any open/active loan prior, and the proposed loan must not exceed 60% of the user's total investment in stocks.

Depeending if loan has been fully repaid, a loan could be open or closed. It is open when the user is yet to conclude repayments on the loan and closed when repayment has been completed on the loan. Repayments are only possible for loans with open status.

Upon taking a loan, the user can then make REPAYMENTS in installments, called REPAYMENTS. T

On Successful loan request, the user wallet is credited with the requested amount which can then be withdrawn to the user registered bank account.

CHECKERS

- In the course of repayment in installements, if a user attempt to pay an amount greater than the loan balance, the system checker will ensure that only the loan balanced is lessed from the user's wallet balance.

- If the user request a loan greater than 60% of her investment in stocks, the checker program will flag the transaction attempt with a 406 error code along with a detailed message.

## - Loan API

To request a loan, the client should make request as detailed below

### Request Loan Payload

```
URL : https://be-ope.herokuapp.com/loan/request
METHODE TYPE: POST
Authorization: Bearer token (Any valid TOKEN of user that is not PAYMENT TOKEN)
PAYLOAD:
{
	"amount" : numerical e.g 32435.00
	"duration" : number of months (Numerical) e.g. 8
}

```

To Fetch user's loans, both open and closed, client should request as below

### Fetch Loans Api

```

URL : https://be-ope.herokuapp.com/loans
METHODE TYPE: GET
Authorization: Bearer token (Any valid TOKEN of user that is not PAYMENT TOKEN)

```

## - Repayment

To repay a loan in installments, the client should send a request as detailed below

### Repayment Api

```

URL : https://be-ope.herokuapp.com/payment/loan/{loanId}
METHODE TYPE: POST
Authorization: Bearer token (Any valid user token that is not PAYMENT TOKEN)
PAYLOAD:
{
    "loan_id": "string"
    "payment": nerical (1200.00)
}

```

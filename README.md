# OPE API Documentation

## Introduction

This is the backend infrastructure build for [Ope App](https://loaner-two.vercel.app/), a stock portfolio web application with loan management features.

To interract with the application core resources, a user is required to register and verified their account. After registering, the user can login to the application using their registered credentials. This will provide the user an hour long authorization token that can be used to access the API core resorces.

### REGISTER

To register a user, the client should make a request as detailed below. AFter complete registration, users are sent verification link via email which can be used to verify user's account by a click.

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

### TO UPDATE ONLY BANK DETAILS

Users can decide to add their bank details to ease wallet funds withdrawer. To do this, clent should send request as below

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
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6I.eyJiYW5rX2FjY291bnQiOnsiYWNjb3VudF9uYW1l ..."
    }
}

```

This will trigger an OTP emailing to the user. the user can then use the otp to complete login as follows

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

SAMPLE RESPONSE

```
{
    "status": "success",
    "collection": "User",
    "data": {
        "user_id": "5",
        "firstname": "Aghahowa",
        "lastname": "Ogie",
        "email": "gulephil44@gmail.com",
        "created_at": "2021-11-15 00:52:42",
        "bank_account": {
            "account_no": "2085xxxxxx",
            "bank_name": "Zenith Bank"
        },
        "otp": "",
        "wallet": {
            "amount": 54000,
            "address": "$2a$06$WRKRFEsAP/meZbjMP1lkOuzyu7jtZ66cu8uH0dQZPKP3pwzDwYRvi"
        },
        "token": "",
        "portfolio": [
            {
                "id": "5",
                "symbol": "AWS",
                "image_url": "https://buyshares.co.uk/wp-content/uploads/2020/07/Screenshot-2020-07-04-at-16.49.55.png",
                "quantity": "120",
                "unit_price": "1200",
                "equity_value": "144000",
                "percentage_change": "-9"
            },
            {
                "id": "15",
                "symbol": "TSLA",
                "image_url": "https://g.foolcdn.com/art/companylogos/square/tsla.png",
                "quantity": "25",
                "unit_price": "800",
                "equity_value": "20000",
                "percentage_change": "-1"
            }
        ]
    }
}

```

NB:

- The portfolio array will be empty if user is yet to buy investment in stock

- The bank attribute will read default state, "none" for both account_no and bank_name if the user is yet to update her bank detals

### FUNDWALLET FLOW

The process of funding a user's wallet include a series of processes. These processes are categorized nto two stages, viz.

Stage ONE

1 - Call FUND WALLET EndPoint

2 - Backend will respond with PAYMENT-BODY and PAYMENT-TOKEN, see PAYMENT RESPONSE below.

3 - Call FLutterwave gateway (Preferably Flutterwave inline), supply the "payment_body" sent in step-2 as payload, and proceed as will be prompted by Flutterwave

Stage TWO

1- Call COMPLETE-FUNDING EndPoint and send Back PAYMENT_TOKEN sent earlier

#### DETAILS:

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

PAYMENT RESPONSE

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
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbW91bnQiOiIxODAwMCIsImN1cnJlbmN5IjoiTkdOIiwiY ..."
    }
}
```

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

### PLEASE NOTE:

1 - tx-ref is transaction reference. It will be the first attribute of the PAYMENT_BODY sent to you when you call FUND_WALLET endpoint

2 - Amount has to be the same as the one in the PAYMENT BODY

3 - wallet is the user's wallet address. It will also be sent as "customer_mac" under the "meta" attribute of the PAYMENT-BODY

### INVESTMENT

Once a user's wallet has been successfully funded, it can then be used to buy investments in company stocks, Repay loans or withdrawn to user's registered bank account. The amount of stocks that can be bought by user is only limited by the user and the amount available in his/her wallet.

To buy investment, the client should call the endpoint as detailed below.

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

Upon taking a loan, the user can then make REPAYMENTS in installments. These installments are called REPAYMENTS and gonverned by the REPAYMENT API

On Successful loan request, the user wallet is credited with the requested amount which can then be withdrawn to the user registered bank account.

CHECKERS

In the course of repayment in installements, if a user attempt to pay an amount greater than the loan balance, the system checker will ensure that only the loan balanced is lessed from the user's wallet balance.

If the user request a loan greater than 60% of her investment in stocks, the checker program will flag the transaction attempt with a 406 error code along with a detailed message.

## - Loan API

To request a loan, the client should make request as detailed below

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

```
URL : https://be-ope.herokuapp.com/loans
METHODE TYPE: GET
Authorization: Bearer token (Any valid TOKEN of user that is not PAYMENT TOKEN)

```

## - Repayment API

To repay a loan in installments, the client should send a request as detailed below

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

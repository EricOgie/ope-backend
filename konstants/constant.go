package konstants

import "time"

const (
	LOCAL_ADD    string = "localhost:8000"
	CONTENT_TYPE string = "Content-Type"
	TYPE_XML     string = "application/xml"
	TYPE_JSON    string = "application/json"
	MSG_500      string = "Unexpected Server Error!"
	MSG_403      string = "User Already Exist"
	MSG_START    string = "Starting Appliction ...."
	LOGGER_TYPE  string = "logger-type"
	TIME         string = "timestamp"
	MSG_404      string = "Resource Not Found"
	QUERY_ERR    string = "Query Error Noticed! Error Msg: "
	KEY_SERVER   string = "SERVER_ADDRESS"
	KEY_PORT     string = "SERVER_PORT"
	KEY_DBUSER   string = "DB_PORT"
	KEY_DB_ADD   string = "DB_ADDR"
	KEY_DB_NAME  string = "DB_NAME"
	AUTH         string = "Authorization"

	MSG_REG string = "Check Email To Verify Account"

	REQ_VALIDITY_ERR string = "422 Validiity Error"

	ERR_REQ_SEND string = "Error while sending req. Err: "
	ERR_DECODE   string = "ERROR while decoding response. e: "

	INVALID_EMAIL  string = "Invalid Email Addrees. Email should follow correct email format"
	INVALID_PWORD  string = "Password must be 6 or more charracters long"
	PHONE_ERR      string = "Phone number must be 11 numbers long"
	BANK_NO_ERR    string = "Invalid Account Number"
	BANK_NAME_ERR  string = "Invalid Bank Name"
	NAME_TOO_SHORT string = "One or two of your names is too short. Name should be 3 or more char long"

	EXP_TIME = int64(time.Hour)

	DB_CON_OK     string = "Database Connection is Successful"
	DB_CON_ERR    string = "Database Connection Error. Err: "
	DB_ERROR      string = "Unexpected Database Error"
	NO_DER_ERR    string = "User no dey. Err = "
	DB_SCAN_ERROR string = "DB Row Scan Error"
	DB_NO_ROW     string = "sql: no rows in result set"
	LOCAL_DB_USER string = "root"
	LOCAL_DB_NAME string = "ope"
	LOCAL_DB_ADD  string = "localhost"

	STRING_INT_ERR    string = "String To Int Conversion Err: "
	INT_TO_STRING_ERR string = "Int To String Conversion Err: "

	ERR  string = "Error"
	USER string = "User"

	BAD_REQ   string = "Bad Request"
	USER_COLL string = "Users Collection"
	LOGIN     string = "Login Token"

	T_FORMAT string = "2006-01-02 15:04:05"

	HASH_ERR       string = "Password Hash Error. Err : "
	DB_INSERT_ERR  string = "DB/User Insert Error. Err: "
	DB_UPDATE_ERR  string = "DB/row update error. Err: "
	DB_ID_ERR      string = "DB/Last Inserted ID retrieval Error. Err: "
	CREDENTIAL_ERR string = "Wrong Email or Password"
	FORBIDDEN      string = "Forbidden"

	VET_ACC_ERR string = "Verify Err: "
	LOGIN_ERR   string = "Login Err: "

	UAUTH_ERR string = "Unauthorized"
	NO_AUTH   string = "Missing auth token"
	EXP_TOKEN string = "Aauthorization has expired"

	Mail_CON_ERR   string = "Mail Server Connection Error. Err: "
	MAIL_DEL_ERR   string = "Mail Sending Err: "
	MAIL_PARSE_ERR string = "ParseFiles Err"
	VERIFY_SUB     string = "Verify Your Ope Account"
	YAHOO          string = "huy@yahoo.com"

	FROM_PREFIX string = "Ope"

	ENV_PROD    string = "Production"
	ERR_OS_READ string = "OS READ ERR = "
	READ_OS     string = "Reading From Production Os "

	ERR_SANITY_CHECK string = "Sanity check Err: "

	VERIFY_URL string = "https://be-ope.herokuapp.com/verified" //
	HOME_URL   string = "https://loaner-two.vercel.app/"
	ROOT_ADD   string = "https://be-ope.herokuapp.com/"
	// ROOT_ADD  string = "http://localhost:8080/"
	LOGIN_URL string = "https://loaner-two.vercel.app/login"

	FLUTTERWAVE_URL string = "https://api.flutterwave.com/v3/payments"

	DT_KEY    string = "props"
	CLAIM_ERR string = "ERROR while trying to extract claims from request context. Err: "

	// MAIL_VET_PATH string = "mailables/verification.html"
	MAIL_VET_PATH string = "mailables/vet.html"
	MAIL_OTP_PATH string = "mailables/twofaemail.html"

	MAIL_PURPOSE_VERIFY string = "verification"
	MAIL_PURPOSE_REQ    string = "Password"
	MAIL_PURPOSE_OTP    string = "OTP"

	MAIL_BTN_VET   string = "Verify My Account"
	MAIL_BTN_PWORD string = "Change Password"

	MAIL_TAIL_ACT_NOTICED    string = "We Notice traffic on your account. Use the OTP below to confirm your operation on Ope app"
	MAIL_TAIL_VERIFY         string = "We would like to verify your account quickly so you can get to the fun part. Click the button below to verify your account"
	MAIL_TAIL_PASSWORD_REQ   string = "Click the button below to change your password"
	MAIL_TAIL_OTP            string = "Use the below one time password to conclude your transaction "
	MAIL_BODY_PASSWORD_RESET string = "You requested for a password change"
	MAIL_BODY_VERIFY         string = "One more step required"
	MAIL_BODY_OTP            string = "You initiated a process"

	CAPTION_WELCOME string = "Welcome To Ope"
	CAPTION_HELLO   string = "Hello"

	SUBJECT_WELCOME         string = "Welcome To Ope"
	SUBJECT_VERIFY_ACC      string = "Verify Your Ope Account"
	SUBJECT_OTP             string = "OTP"
	SUBJECT_PASSWORD_CHANGE string = "Change Password"

	PAY_VALIDATION_ERR_MSG string = "Amount >= 5000 & PaymentOption = Card & Curreny = NG"

	ERR_FLOAT_CONV string = "Couldn't Make float64"
	ERR_FRAUD      string = "Fraudulent transaction"
	ERR_TRANS_REC  string = "DID NOT SAVE TRANSACTION"
	ERR_USER_ID    string = "Invalid user Id"
	ERR_ACC_NO     string = "Invalid Account Number"
	ERR_BANK_NAME  string = "Invalid Bank Name"

	ERR_INSURFICIENCY string = "Insurficient Funds"
	ERR_60PER_CHECK   string = "Loan request can not exceed 60% of total Investment"
	ERR_OPEN_LOAN     string = "All Pending Loans must be closed before another loan can be granted"
	ERR_WALLET_CREDIT string = "Loan went through but user wallet was not credited"
	ERR_LOAN_AMT      string = "Invalid Loan Amount"
	ERR_LOAN_DUR      string = "Invalid Loan Duration"
	ERR_LOAN_ID       string = "Invalid Laon Id"
	ERR_PAY_AMT       string = "Invalid Payment Amount"

	LOAN_OPEN     string = "open"
	LOAN_INACTIVE string = "This loan has been closed"
)

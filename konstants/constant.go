package konstants

import "time"

const (
	LOCAL_ADD    string = "localhost:8000"
	CONTENT_TYPE string = "Content-Type"
	TYPE_XML     string = "application/xml"
	TYPE_JSON    string = "application/json"
	MSG_500      string = "Unexpected Server Error!"
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

	ERR_REQ_SEND string = "Error while sending req. Err: "
	ERR_DECODE   string = "ERROR while decoding response. e: "

	INVALID_EMAIL  string = "Invalid Email Addrees. Email should follow correct email format"
	INVALID_PWORD  string = "Password must be 6 or more charracters long"
	PHONE_ERR      string = "Phone number must be 11 numbers long"
	NAME_TOO_SHORT string = "One or two of your names is too short. Name should be 3 or more char long"

	EXP_TIME = int64(time.Hour)

	DB_CON_OK     string = "Database Connection is Successful"
	DB_CON_ERR    string = "Database Connection Error. Err: "
	DB_ERROR      string = "Unexpected Database Error"
	NO_DER_ERR    string = "User no dey. Err = "
	DB_SCAN_ERROR string = "DB Row Scan Error"
	LOCAL_DB_USER string = "root"
	LOCAL_DB_NAME string = "ope"
	LOCAL_DB_ADD  string = "localhost"

	T_FORMAT string = "2006-01-02 15:04:05"

	HASH_ERR       string = "Password Hash Error. Err : "
	DB_INSERT_ERR  string = "DB/User Insert Error. Err: "
	DB_UPDATE_ERR  string = "DB/row update error. Err: "
	DB_ID_ERR      string = "DB/Last Inserted ID retrieval Error. Err: "
	CREDENTIAL_ERR string = "Wrong Email or Password"

	VET_ACC_ERR string = "Verify Err: "
	LOGIN_ERR   string = "Login Err: "

	UAUTH_ERR string = "Unauthorized"
	NO_AUTH   string = "Missing auth token"
	EXP_TOKEN string = "Aauthorization has expired"

	Mail_CON_ERR   string = "Mail Server Connection Error. Err: "
	MAIL_DEL_ERR   string = "Mail Sending Err: "
	MAIL_PARSE_ERR string = "ParseFiles Err"
	VERIFY_SUB     string = "Verify Your Ope Account"
	YAHOO          string = "ericogia@yahoo.com"

	FROM_PREFIX string = "Ope"

	VERIFY_URL string = "http://localhost:8080/verify" //

	DT_KEY    string = "props"
	CLAIM_ERR string = "ERROR while trying to extract claims from request context. Err: "

	MAIL_VET_PATH string = "mailables/verification.html"
	MAIL_OTP_PATH string = "mailables/otp.html"

	MAIL_ACT_NOTICED string = "We Notice traffic on your account.Enter the OTP below to confirm your operation on Ope app"
)

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

	INVALID_EMAIL  string = "Invalid Email Addrees. Email should follow correct email format"
	INVALID_PWORD  string = "Password must be 6 or more charracters long"
	PHONE_ERR      string = "Phone number must be 11 numbers long"
	NAME_TOO_SHORT string = "One or two of your names is too short. Name should be 3 or more char long"

	EXP_TIME = int64(time.Hour)

	DB_ERROR      string = "Unexpected Database Error"
	DB_SCAN_ERROR string = "DB Row Scan Error"
	LOCAL_DB_USER string = "root"
	LOCAL_DB_NAME string = "ope"
	LOCAL_DB_ADD  string = "localhost"

	T_FORMAT string = "2006-01-02 15:04:05"

	DB_INSERT_ERR string = "DB/User Insert Error. Err: "
	DB_ID_ERR     string = "DB/Last Inser ID retrieval Error. Err: "
)

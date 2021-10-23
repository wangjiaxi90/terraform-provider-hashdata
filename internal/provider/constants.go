package provider

const (
	DEFAULT_ENDPOINT       = "https://console.hashdata.xyz"
	AUTH_PREFIX            = "Bearer "
	TOKEN_PATH             = "/account/oauth/token"
	DEFAULT_HEADER_KEY     = "Authorization"
	MAX_RETRY_TIMES        = 2000
	waitJobTimeOutDefault  = 6000
	waitJobIntervalDefault = 5
	JOB_SUCCESS            = "success"
	JOB_WAIT_PENDING       = "pending"
	JOB_WAIT_RUNNING       = "running"
	JOB_FAILED_FAILURE     = "failure"
	JOB_FAILED_ABANDONED   = "abandoned"
	WAREHOUSE_ID           = "warehouse_id"
)

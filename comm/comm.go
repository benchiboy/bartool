// CommonInfo.go

package comm

const (
	LOG_SYS_PREFIX  = "==="
	LOG_SEND_PREFIX = "===>"
	LOG_RECV_PREFIX = "<==="
	ADMIN_USER      = "admin"
	DEFAULT_PWD     = "admin123!"
	MAX_COUNT       = 100
	LOG_START_FUNC  = "************************************************************************************************"
)

var (
	Is_Quit bool
)

type Login_Node struct {
	User_Name string `json:"username"`
	User_Pwd  string `json:"password"`
}

type ChangePwd_Node struct {
	User_Name string `json:"user_name"`
	Old_Pwd   string `json:"old_pwd"`
	New_Pwd   string `json:"new_pwd"`
}

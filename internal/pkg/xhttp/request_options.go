package xhttp

const groupPath = "group_path"
const GroupPathHeader = "X-Group-Path"
const RequestIDHeader = "X-Request-ID"

type RequestOption struct {
	Header    map[string]string
	GroupPath string
}

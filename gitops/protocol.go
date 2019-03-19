package gitops

/*Protocol used for cloning*/
type Protocol string

/*Cloning protocols*/
const (
	HTTPS           Protocol = "https"
	SSH             Protocol = "ssh"
	GIT             Protocol = "git"
	DefaultProtocol Protocol = GIT
)

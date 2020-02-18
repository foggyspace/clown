package bruteforce

// digest authentication struct
type DigestAuthentication struct {
	UserName  string `json:"digest_username"`
	Password  string `json:"digest_password"`
	Message   string `json:"digest_message"`
	TargetURL string `json:"digest_target_url"`
}

// digest authentication bruteforce
func (d *DigestAuthentication) DigestAuthenticationBruteforce(targetURL string) {}

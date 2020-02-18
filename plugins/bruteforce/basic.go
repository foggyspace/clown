package bruteforce

// basic authorication struct
type BasicBruteforce struct {
	UserName  string `json:"basic_username"`
	Password  string `json:"basic_password"`
	Message   string `json:"basic_message"`
	TargetURL string `json:"basic_target_url"`
}

// basic bruteforce
func (b *BasicBruteforce) BruteforceBasicAuthentication(targetURL string) {}

func BasicBruteforceRun(targetURL string) {
	basic := BasicBruteforce{}
	basic.BruteforceBasicAuthentication(targetURL)
}

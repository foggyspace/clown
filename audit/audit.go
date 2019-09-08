package audit

type Auditor interface {
	Audit(freq, origi_resp interface{})
	Message(msg string)
}

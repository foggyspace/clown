package audit

type Xss struct{}

func (self *Xss) Audit(frep, origin_resp interface{}) {}

func (self *Xss) Message(msg string) {}

package audit

type Csrf struct{}

func (self *Csrf) Audit(frep, orgin_resp interface{}) {}

func (self *Csrf) Message(msg string) {}

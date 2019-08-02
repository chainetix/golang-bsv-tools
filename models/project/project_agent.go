package models

type Agent struct {
	ProjectProto
	Label string `json:"label"`
}

func (self *Agent) GetUID() string {
	return self.UID
}

func (self *Agent) Tree() []string {
	return []string{
		self.Project,
		self.UID,
	}
}

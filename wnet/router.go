package wnet

import "tcpw/wiface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request wiface.IRequest)  {}
func (br *BaseRouter) Handle(request wiface.IRequest)     {}
func (br *BaseRouter) PostHandle(request wiface.IRequest) {}

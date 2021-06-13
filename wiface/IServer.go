package wiface

type IServer interface {
    Start()
    Stop()
    Serve()
}

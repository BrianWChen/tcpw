package utils

import (
    "encoding/json"
    "io/ioutil"
    "tcpw/wiface"
)

type GlobalObj struct {
    /*
    	Server
    */
    TCPServer wiface.IServer
    Host      string
    TCPPort   int
    Name      string

    /*
    	TCPw
    */
    Version        string
    MaxConn        int
    MaxPackageSize uint32
    WorkerPoolSize uint32 //业务工作Worker池的数量
    MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
    data, err := ioutil.ReadFile("conf/tcpw.json")
    if err != nil {
        panic(err)
    }
    //将json数据解析到struct中
    err = json.Unmarshal(data, &GlobalObject)
    if err != nil {
        panic(err)
    }
}

func init() {
    GlobalObject = &GlobalObj{
        Name:           "TCPwServerApp",
        Version:        "V0.10",
        TCPPort:        8999,
        Host:           "0.0.0.0",
        MaxConn:        12000,
        MaxPackageSize: 4096,
        WorkerPoolSize: 10,
        MaxWorkerTaskLen: 1024,
    }

    //NOTE: 从配置文件中加载一些用户配置的参数
    GlobalObject.Reload()
}

package wnet

import (
    "fmt"
    "strconv"

    "tcpw/wiface"
)

// MsgHandle -
type MsgHandle struct {
    Apis           map[uint32]wiface.IRouter //存放每个MsgID 所对应的处理方法的map属性
}

//NewMsgHandle 创建MsgHandle
func NewMsgHandle() *MsgHandle {
    return &MsgHandle{
        Apis:           make(map[uint32]wiface.IRouter),
    }
}

//DoMsgHandler 马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(request wiface.IRequest) {
    handler, ok := mh.Apis[request.GetMsgID()]
    if !ok {
        fmt.Println("api msgID = ", request.GetMsgID(), " is not FOUND!")
        return
    }

    //执行对应处理方法
    handler.PreHandle(request)
    handler.Handle(request)
    handler.PostHandle(request)
}

//AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router wiface.IRouter) {
    //1 判断当前msg绑定的API处理方法是否已经存在
    if _, ok := mh.Apis[msgID]; ok {
        panic("repeated api , msgID = " + strconv.Itoa(int(msgID)))
    }
    //2 添加msg与api的绑定关系
    mh.Apis[msgID] = router
    fmt.Println("Add api msgID = ", msgID)
}

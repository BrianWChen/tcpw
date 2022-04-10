package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("client0 start...")

	time.Sleep(1 * time.Second)

	for i := 0; i < 100; i++ {
		go func() {
			conn, err := net.Dial("tcp", "127.0.0.1:5000")
			if err != nil {
				fmt.Println("client start err, exit!")

				return
			}

			go func() {
				for {
					var siteId = int(rand.Intn(100))
					siteId++
					siteIdStr := strconv.Itoa(siteId)
					_, err0 := conn.Write([]byte("ST:" + siteIdStr + ";4294966825/01/01 10:13:18;A1:32.8;A2:0.5;A3:462.0;A4:4.3;A5:45.1;A6:46.4;A7:138.7;B:00;Q:0001;J1:1;R:00;W:30.0"))
					if err0 != nil {
						fmt.Println("writ conn err ", err)
						return
					}
					//time.Sleep(60 * time.Second)
					time.Sleep(100 * time.Millisecond)
					//
					//_, err1 := conn.Write([]byte("ST:3;4294966825/01/01 10:13:18;A1:15.3;A2:0.0;A3:279.0;A4:3.1;A5:61.6;A6:36.1;A7:536.0;B:110;Q:0111;J1:1;R:12;W:27.0."))
					//if err1 != nil {
					//	fmt.Println("writ conn err ", err)
					//	return
					//}
					//
					//time.Sleep(100 * time.Millisecond)
					//
					//_, err2 := conn.Write([]byte("ST:1;4294966825/01/01 10:13:18;A1:14.4;A2:2.0;A3:279.0;A4:3.1;A5:61.6;A6:36.1;A7:536.0;B:00;Q:01;J1:1;R:12;W:27.0."))
					//if err2 != nil {
					//	fmt.Println("writ conn err ", err)
					//	return
					//}
					//
					//time.Sleep(100 * time.Millisecond)
					//
					//_, err3 := conn.Write([]byte("ST:2;4294966825/01/01 10:13:18;A1:14.4;A2:0.0;A3:279.0;A4:3.1;A5:61.6;A6:36.1;A7:536.0;B:00;Q:01;J1:1;R:12;W:27.0."))
					//if err3 != nil {
					//	fmt.Println("writ conn err ", err)
					//	return
					//}
					//
					//_, err4 := conn.Write([]byte("4294966825/01/01 10:13:18;A1:14.4;A2:0.0;A3:279.0;A4:3.1;A5:61.6;A6:36.1;A7:536.0;B:00;Q:01;J1:1;R:12;W:27.0.ST:9999;4294966825/01/01 10:13:18;A1:14.4;"))
					//if err4 != nil {
					//	fmt.Println("writ conn err ", err)
					//	return
					//}
					//
					//_, err5 := conn.Write([]byte("4294966825/01/01 10:13:18;A1:14.4;A2:0.0;A3:279.0;A4:3.1;A5:61.6;A6:36.1;A7:536.0;B:00;Q:01;J1:1;R:12;W:27.0."))
					//if err5 != nil {
					//	fmt.Println("writ conn err ", err)
					//	return
					//}

					// recvData := make([]byte, 15)
					// _, err4 := io.ReadFull(conn, recvData)
					// if err4 != nil {
					//   fmt.Println("server unpack data err:", err)
					//   return
					// }

					// fmt.Println("==> Recv Msg Data=", string(recvData))

					//time.Sleep(5 * time.Millisecond)
					//bytes := Int2Byte(100)
					//n2 := 100
					//n2Str16 := Int2Str16(n2)
					//fmt.Println(n2Str16)
					////n2Byte := Hex2Byte4(n2Str16)
					////fmt.Println(n2Byte)
					////n2Byte2 := Hex2Byte2(n2Str16)
					////fmt.Println(n2Byte2)
					//
					//data := [50]byte{0x53, 0x50, 0x30, 0x01}
					//// Site ID
					//siteBytes := Hex2Byte4(int32(14))
					//data[4] = siteBytes[0]
					//data[5] = siteBytes[1]
					//data[6] = siteBytes[2]
					//data[7] = siteBytes[3]
					//
					//hwList := strings.Split(strings.Trim("V3.5", "V"), ".")
					//hwStr := hwList[0] + hwList[1]
					//hwByte := Hex2Byte(hwStr)
					//fmt.Println(hwByte)
					//data[8] = Hex2Byte(hwStr)
					//
					//
					//fmt.Println(data)
					//n2Byte := Int2Byte(n2Str16.)
				}
			}()

			for {
				data := make([]byte, 100)
				data, err = io.ReadAll(conn) //ReadFull 会把msg填充满为止
				if err != nil {
					fmt.Println("read head error")
					continue
				}
				fmt.Println(data)
			}
		}()
	}

	for {
		time.Sleep(60 * time.Second)
	}
}

func Int2Str16(n int) string {
	return strconv.FormatInt(int64(n), 16)
}

func Hex2Byte4(x int32) (ret []byte) {
	//var len = unsafe.Sizeof(data)
	//ret = make([]byte, len)
	//var tmp = 0xff
	//var index uint = 0
	//for index = 0; index < uint(len); index++ {
	//    ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	//}
	//return ret
	//res, _ := strconv.ParseInt(str, 10, 32)
	//x := int32(res)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()

	//slen := len(str)
	//bHex := make([]byte, 8 / 2)
	//ii := 0
	//for i := 0; i < len(str); i = i + 2 {
	//	if slen != 1 {
	//		ss := string(str[i]) + string(str[i+1])
	//		bt, _ := strconv.ParseInt(ss, 10, 32)
	//		bHex[ii] = byte(bt)
	//		ii = ii + 1
	//		slen = slen - 2
	//	}
	//}
	//
	//return bHex
}

func Hex2Byte2(x int16) (ret []byte) {
	//var len = unsafe.Sizeof(data)
	//ret = make([]byte, len)
	//var tmp = 0xff
	//var index uint = 0
	//for index = 0; index < uint(len); index++ {
	//    ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	//}
	//return ret
	//res, _ := strconv.ParseInt(str, 10, 32)
	//x := int16(res)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()

	//slen := len(str)
	//bHex := make([]byte, 8 / 2)
	//ii := 0
	//for i := 0; i < len(str); i = i + 2 {
	//	if slen != 1 {
	//		ss := string(str[i]) + string(str[i+1])
	//		bt, _ := strconv.ParseInt(ss, 10, 32)
	//		bHex[ii] = byte(bt)
	//		ii = ii + 1
	//		slen = slen - 2
	//	}
	//}
	//
	//return bHex
}

//func Number2BCD(number string) []byte {
//	var rNumber = number
//	for i := 0; i < 8 - len(number); i++ {
//		rNumber = "f" + rNumber
//	}
//	bcd := Hex2Byte(rNumber)
//	return bcd
//}

func Hex2Byte(str string) byte {
	ss := string(str[0]) + string(str[1])
	bt, _ := strconv.ParseInt(ss, 16, 32)
	return byte(bt)
}

//func Number2BCD(number string) []byte {
//	var rNumber = number
//	for i := 0; i < 8 - len(number); i++ {
//		rNumber = "f" + rNumber
//	}
//	bcd := Hex2Byte(rNumber)
//	return bcd
//}
//
//func Hex2Byte(str string) []byte {
//	slen := len(str)
//	bHex := make([]byte, 1)
//	ii := 0
//	for i := 0; i < len(str); i = i + 2 {
//		if slen != 1 {
//			ss := string(str[i]) + string(str[i+1])
//			bt, _ := strconv.ParseInt(ss, 16, 32)
//			bHex[ii] = byte(bt)
//			ii = ii + 1
//			slen = slen - 2
//		}
//	}
//
//	return bHex
//}

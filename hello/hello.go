package main

import (
	"fmt"
	//"time"
	//"godev/newmath"
	//"github.com/yeochinyi/newmath"
	// "net/http/httputil"
	//"bufio"
	//"net"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	//"io"
	//"encoding/binary"
	//"encoding/hex"
	"github.com/yeochinyi/godev/codec"
	"os"
)

func main() {
	//var _ = foo
	//fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))
	//fmt.Printf("Hello, world.")
	//conn := net.Dial("tcp",)
	//client := httputil.NewClientConn()
	//client := &http.DefaultClient
	//client.Get()
	//client := &http.Client
	//client.
	//conn, _ := net.Dial("tcp", "sinrs03064.fm.rbsgrp.net:8080")
	//fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	//status, _ := bufio.NewReader(conn).ReadString('\n')
	//fmt.Print(status)
	//resp, _ := http.Get("http://sinrs03064.fm.rbsgrp.net:8080/nexus/index.html#welcome")
	//fmt.Print(resp)
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//for c := range body {
	//	fmt.Print(string(c))
	//}
	//fmt.Print(string(body))

	//http.ProxyURL()
	//tr := &http.Transport{
	//	Proxy: func(*Request) (*url.URL, error)
	//}
	//proxyUrl, err := url.Parse("http://fx_build:Franom52@FM-AP-HKG-PROXY.fm.rbsgrp.net:8080")
	//myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	//resp, _ = myClient.Get("http://www.google.com")

	//os.Setenv("HTTP_PROXY", "http://fx_build:Franom52@FM-AP-HKG-PROXY.fm.rbsgrp.net:8080")
	//resp, _ := http.Get("http://www.google.com")
	//body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Print(string(body))

	f, err := os.Open("c:/temp/heap.20140109.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, e := codec.MakeMsg(f, 31)
	if e != nil {
		panic(e)
	}

	m.Get(19)
	idSize := codec.ToInt(m.Get(4))
	fmt.Printf("id Size:%v\n", idSize)
	timeSince := codec.ToInt(m.Get(8))
	fmt.Printf("timeSince:%v\n", codec.MsToTime(timeSince))

	counterMap := make(map[uint64]uint64)

	//stringCounter := 0
	for {
		m, e = codec.MakeMsg(f, 9)
		if e != nil {
			//panic(e)
			break
		}
		msgType := codec.ToInt(m.Get(1))
		//fmt.Printf("msgType:%v\n", msgType)
		_ = msgType
		msSince := codec.ToInt(m.Get(4)) //Time
		//fmt.Printf("msSince:%v\n", msSince)
		_ = msSince
		msgLen := codec.ToInt(m.Get(4))
		_ = msgLen
		//fmt.Printf("msgLen:%v\n", msgLen)
		msgText, e := codec.MakeMsg(f, msgLen)
		if e != nil {
			//panic(e)
			break
		}
		_ = msgText
		//fmt.Printf("msgText:%s\n", msgText.data)

		counterMap[msgType]++

		switch msgType {
		case 0x01:
			//stringCounter++
			//fmt.Printf("String %s", msgText)

		case 0x05:
			//Print(msgText.Get(4))
			//Print(msgText.Get(4))
			//Print(Rev(msgText.Get(4)))
		case 0x0C:
			fallthrough
		case 0x1C:
			fmt.Printf("msgLen:%v\n", msgLen)
			HeapDump(&msgText, idSize)
		default:
			//panic(fmt.Sprintf("msgType %v not found!", msgType))
		}

	}

	//fmt.Print(counterMap)
	//fmt.Printf("Total strings : %v", stringCounter)
	for k, v := range counterMap {
		fmt.Printf(" %#x -> %v  ", k, v)
	}

}

func HeapDump(m *codec.Message, idSize uint64) {

	fmt.Printf("heapDump:%v\n", m.Data())

	var entryTypeMap = map[uint64]uint64{
		2:  8, //obj
		4:  1, //bool
		5:  2, //char
		6:  4, //float
		7:  8, //double
		8:  1, //byte
		9:  2, //short
		10: 4, //int
		11: 8, //long

	}

	for {
		subtag := codec.ToInt(m.Get(1))
		fmt.Printf("subtag %v\n", subtag)
		switch subtag {
		case 0xFF:
			m.Get(idSize)
		case 0x01:
			m.Get(idSize * 2)
		case 0x02:
			m.Get(8 + idSize)
		case 0x03:
			m.Get(8 + idSize)
		case 0x04:
			m.Get(4 + idSize)
		case 0x05:
			m.Get(idSize)
		case 0x06:
			m.Get(4 + idSize)
		case 0x07:
			m.Get(idSize)
		case 0x08:
			m.Get(8 + idSize)
		case 0x20:
			m.Get(8 + idSize*7)
			//size of constant pool and number of records that follow:
			constantPoolSize := int(codec.ToInt(m.Get(2)))
			for i := 0; i < constantPoolSize; {
				m.Get(2)
				entryType := codec.ToInt(m.Get(1))
				entrySize := entryTypeMap[entryType]
				m.Get(entrySize)
			}
			//Number of static fields:
			staticFieldsSize := int(codec.ToInt(m.Get(2)))
			for i := 0; i < staticFieldsSize; {
				m.Get(idSize)
				entryType := codec.ToInt(m.Get(1))
				entrySize := entryTypeMap[entryType]
				m.Get(entrySize)
			}
			//Number of instance fields:
			instanctFieldsSize := int(codec.ToInt(m.Get(2)))
			for i := 0; i < instanctFieldsSize; {
				m.Get(idSize)
				entryType := m.Get(1)
				_ = entryType
			}
		case 0x21:
			m.Get(idSize*2 + 4)
			bytes := codec.ToInt(m.Get(4))
			m.Get(bytes)
		case 0x22:
			m.Get(idSize + 4)
			numElements := codec.ToInt(m.Get(4))
			m.Get(idSize)
			m.Get(numElements * idSize)
		case 0x23:
			m.Get(idSize + 4)
			numElements := codec.ToInt(m.Get(4))
			entryType := codec.ToInt(m.Get(1))
			entrySize := entryTypeMap[entryType]
			m.Get(entrySize * numElements)
		default:
			panic(fmt.Sprintf("No case for %#x", subtag))
		}
	}
}

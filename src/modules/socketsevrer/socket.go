package socketsevrer

import (
	"net"
)


func launch(address any,port any){
	var _address = address
	var _port = port
	
	if (_address == nil){
		_address = "0.0.0.0"
	}
	if (_port == nil){
		_port = 20642
	}
	// 强制转换类型
	tcp6listener,_:=net.Listen("tcp6",address.(string)+":"+string(port.(int)))
	tcp4listener,_:=net.Listen("tcp4",address.(string)+":"+string(port.(int)))
}
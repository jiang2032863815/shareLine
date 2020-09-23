package socks5

import "net"

type Server struct{
	listenAddr string
	authType int
}
func(s *Server)serveProxy(conn net.Conn){
	defer conn.Close()

}
func(s *Server)ListenAndRun(){
	if listener,err:=net.Listen("tcp",s.listenAddr);err==nil{
		defer listener.Close()
		for{
			if conn,err:=listener.Accept();err==nil{
				go s.serveProxy(conn)
			}
		}
	}else{
		panic(err)
	}
}
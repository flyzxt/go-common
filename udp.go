package common

import (
	"fmt"
	"net"
)

func RunUDPServer() {
	addr := "0.0.0.0"
	port := 3478
	udpAddr := net.UDPAddr{IP: net.ParseIP(addr), Port: port}
	udpConn, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		fmt.Printf("ListenUDP error: %v\n", err)
		defer udpConn.Close()
		return
	}
	fmt.Printf("Listening on %s:%d\n", udpAddr.IP, udpAddr.Port)
	for {
		buffer := make([]byte, 2048)
		n, remoteAddr, err := udpConn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("ReadFromUDP error: %v\n", err)
			defer udpConn.Close()
			return
		}
		fmt.Printf("Received %s from %s\n", buffer[:n], remoteAddr)
		// send new message back to the client
		_, err = udpConn.WriteToUDP(buffer[:n], remoteAddr)
		if err != nil {
			fmt.Printf("WriteToUDP error: %v\n", err)
			defer udpConn.Close()
			return
		}
		fmt.Printf("Sent %s to %s\n", buffer[:n], remoteAddr)
	}
}

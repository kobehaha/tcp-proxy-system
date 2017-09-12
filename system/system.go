package system

import "net"

// description
// const channale Num
const ChannelPairNum = 2

// description
// iner channel struct for managing src, dst connection 

type Channel struct {
    SrcConnection net.Conn
    DstConnection net.Conn
}
// descrition
// get src url from connection
func (channel *Channel) SrcUrl() string{
    return channel.SrcConnection.RemoteAddr().String()
}

// descrition
// get dst url from connection
func (channel *Channel) DstUrl() string {
    return channel.DstConnection.RemoteAddr().String()
}

// descrition
// channel  close
func (channel *Channel) Close() {
    channel.SrcConnection.Close()
    channel.DstConnection.Close()
}



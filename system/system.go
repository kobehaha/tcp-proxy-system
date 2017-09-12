package system

import "net"

const ChannelPairNum = 2

type Channel struct {
    SrcConnection net.Conn
    DstConnection net.Conn
}

func (channel *Channel) SrcUrl() string{
    return channel.SrcConnection.RemoteAddr().String()
}

func (channel *Channel) DstUrl() string {
    return channel.DstConnection.RemoteAddr().String()
}

func (channel *Channel) Close() {
    channel.SrcConnection.Close()
    channel.DstConnection.Close()
}



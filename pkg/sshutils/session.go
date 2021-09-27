package sshutils

import "golang.org/x/crypto/ssh"

//创建一个ssh连接,重要是创建一个session,session下面有很多方法,可以执行相关命令,
//而session是由client类型创建的,通过client创建session只有一个方法NewSession(),创建session没有任何参数
//因此首先我们需要创建一个Client
//创建一个Client有两种方法
//func Dial(network, addr string, config *ClientConfig) (*Client, error)
//func NewClient(c Conn, chans <-chan NewChannel, reqs <-chan *Request) *Client
//

type SSHSession struct {
	User      string
	Host      string
	Auth      string
	Session   *ssh.Session
}


func aa()  {

}


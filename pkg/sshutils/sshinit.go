package sshutils

import (
	"golang.org/x/crypto/ssh"
	"time"
)

type SSHClient *ssh.Client
type SSHClientConfig *ssh.ClientConfig

type SSH struct {
	User                    string
	Password                string
	PrivateKeyFileName      string
	PrivateKeyPassword      string
	AuthMethod              string
	Timeout                 *time.Duration
}

//任何一个ssh连接都必须有一个该连接相关的配置,*ssh.ClientConfig,
/*
1.config: config是客户端和服务端必须共同的一些配置,
2.user: 连接所使用的用户名
3.Auth:
4.HostKeyCallback:
5.BannerCallback:
6.HostKeyAlgorithms:
7.Timeout:
*/
func (ss *SSH)GetSSHClientConfig(user string) *ssh.ClientConfig {

	config := ssh.Config{
		KeyExchanges: nil,
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}

	auth := ss.GetAuthMethod()

	clientConfig := &ssh.ClientConfig{
		Config:  config,
		User:    user,
		Auth:    auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		BannerCallback: BannerCallback,
		HostKeyAlgorithms: []string,
		Timeout: *ss.Timeout,

	}

	return clientConfig

}


func (ss *SSH) NewSSHClient(user string,host string) *ssh.Client {

    clientConfig := ss.GetSSHClientConfig(user)
	sshClient, err := ssh.Dial("tcp", host, clientConfig)
	if err != nil {
		panic(err)
	}

	return sshClient
}

func (ss *SSH) GetSSHSession() *ssh.Session{

	sshClient := ss.NewSSHClient("root","host")

	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return session
}

func (ss *SSH) GetAuthMethod() (auth []ssh.AuthMethod) {
	//
	if ss.AuthMethod == "sshKey"{
		auth = append(auth,ssh.Password("passwd"))
	} else if ss.AuthMethod == "password"{
		auth = append(auth,ssh.Password("passwd"))
	} else{

	}

	return auth
}

func (ss *SSH) NewSSHSession() *ssh.Session{



	sshClient := ss.NewSSHClient("root","host")
	session, err := sshClient.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return session
}



package cmd

import "github.com/8zhiniao/hoesman/pkg/sshutils"

type Config struct {


}


func Exec(){

	sshAgent := sshutils.SSH{

	}
	sshAgent.ExecCommand("192.168.10.1","ls -l")

}
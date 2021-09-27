package sshutils



func (ss *SSH) ExecCommand(host string,cmd string) (resp []byte,err error) {

	session := ss.NewSSHSession()
	output, err1 := session.CombinedOutput(cmd)
	if err1 != nil {
		return nil, err1
	}

	return output, nil
}

func (ss *SSHSession) CombinedOutput(cmd string)   {

}

func (ss *SSHSession) Output(cmd string)  {

}

func (ss *SSHSession) Run(cmd string)  {
	
}

func (ss *SSHSession) Start(cmd string)  {

}

func (ss *SSHSession) SetEnv(cmd string)  {

}


func (ss *SSHSession) Close(cmd string)  {

}


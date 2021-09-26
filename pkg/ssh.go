package pkg

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

type SSHConfig struct {
	User string
	Key string
	Host string
}

func Connect(conf SSHConfig) (*ssh.Client, *ssh.Session, error) {

	key, err := ioutil.ReadFile(conf.Key)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to parse private key: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: conf.User,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", conf.Host, sshConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("can't connect to remote server: %v", err)
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
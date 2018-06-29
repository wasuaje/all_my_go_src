package remote_runner

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"os"
	"path/filepath"
	"bufio"
	"strings"
	"fmt"
	"log"
	//"fmt"
	//"time"
	//"io/ioutil"
	//"log"
	//"os"
	//"fmt"
	//"path/filepath"
	//"bufio"
	//"strings"
	"errors"
	"io/ioutil"

)


func getHostKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
			}
			break
		}
	}

	if hostKey == nil {
		return nil, errors.New(fmt.Sprintf("no hostkey for %s", host))
	}
	return hostKey, nil
}

func Challenge(username, instruction string, questions []string, echos []bool) (answers []string, err error) {
	var suitable_answers = "Asg4rd14!"
	var pwIdx = 0
	answers = make([]string, len(questions))
	for n, _ := range questions {
		//fmt.Printf("Got question: %s\n", q)
		answers[n] = suitable_answers
	}
	pwIdx++

	return answers, nil
}

func ExecuteCmd(cmd string, username string, hostname string) string {
	hostKey, err := getHostKey(hostname)
	if err != nil {
		log.Fatal(err)
	}
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	//key_path := filepath.Join(home + "/.ssh/id_rsa")
	key_path := filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
	//fmt.Println(key_path)
	key, err := ioutil.ReadFile(key_path)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	//fmt.Println(signer)

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(Challenge),
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
			//ssh.Password("Esp4rt4!"),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),


	}

	conn, err := ssh.Dial("tcp", hostname+":22", config)
	if err != nil {
		log.Fatalf("unable to connect to host: %v\n %v", hostname, err)
	}
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("unable to create a new session to host: %v\n %v", hostname, err)
	}

	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	//session.Run(cmd)

	if err := session.Run(cmd); err != nil {
		panic("Failed to run: " + err.Error())
	}

	return stdoutBuf.String()
}
//
//func main() {
//	cmd := os.Args[1]
//	hosts := os.Args[2:]
//
//	results := make(chan string, 10)
//	timeout := time.After(5 * time.Second)
//	config := &ssh.ClientConfig{
//		User: os.Getenv("LOGNAME"),
//		Auth: []ssh.AuthMethod{makeKeyring()},
//	}
//
//	for _, hostname := range hosts {
//		go func(hostname string) {
//			results <- executeCmd(cmd, hostname, config)
//		}(hostname)
//	}
//
//	for i := 0; i < len(hosts); i++ {
//		select {
//		case res := <-results:
//			fmt.Print(res)
//		case <-timeout:
//			fmt.Println("Timed out!")
//			return
//		}
//	}
//}
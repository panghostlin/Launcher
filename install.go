/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Monday 10 February 2020 - 10:52:20
** @Filename:				install.go
**
** @Last modified by:		Tbouder
** @Last modified time:		Tuesday 14 April 2020 - 20:02:57
*******************************************************************************/

package			main

import			"os"
import			"crypto/rand"
import			"encoding/base64"
import			"bufio"
import			"fmt"
import			"os/exec"
import			"strings"

func	generateNonce(n uint32) (string) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if (err != nil) {
        return ``
	}
    ciphertext := base64.RawStdEncoding.EncodeToString(b)
	
    return ciphertext
}

func	prompt() string {
	fmt.Println(`-- DOMAIN --`)
	fmt.Print(`Please, enter the domain to use : `)
	reader := bufio.NewReader(os.Stdin)
	for {
		domainName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return domainName
	}
}
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
		// add another case here for custom commands.
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func	main()	{
	domain := prompt()
	fmt.Println(`You Panghostlin domain will be : ` + domain)

	fmt.Println(`-- GENERATING ENV FILE --`)
	f, err := os.Create(".env")
	if (err != nil) {
		panic(err)
	}
	defer f.Close()
	f.WriteString("DOMAIN=" + strings.TrimSpace(domain) + "\n")
	fmt.Println(`Domain :   [OK]`)

	f.WriteString("POSTGRE_USERNAME=" + `Panghostlin` + "\n")
	f.WriteString("POSTGRE_PWD=" + generateNonce(64) + "\n")
	f.WriteString("POSTGRE_URI=" + `panghostlin-postgre` + "\n")
	f.WriteString("POSTGRE_PORT=" + `54320` + "\n")
	f.WriteString("POSTGRE_DB=" + `panghostlin` + "\n")
	fmt.Println(`Database : [OK]`)

	f.WriteString("MASTER_KEY=" + generateNonce(32) + "\n")
	f.WriteString("MASTER_PUBLIC_KEY=" + generateNonce(32) + "\n")
	f.WriteString("MASTER_KEY_ARGON2=" + generateNonce(32) + "\n")
	f.WriteString("MASTER_KEY_SCRYPT=" + generateNonce(32) + "\n")
	f.WriteString("PRIV_KEY=" + generateNonce(64) + "\n")
	fmt.Println(`Keys :     [OK]`)

	f.WriteString("JWT_ACCESS_TOKEN_KEY=" + generateNonce(64) + "\n")
	f.WriteString("JWT_REFRESH_TOKEN_KEY=" + generateNonce(64) + "\n")
	fmt.Println(`JWT :      [OK]`)
	f.WriteString("IS_DEV=false\n")

	fmt.Println(`-- DONE --`)
}
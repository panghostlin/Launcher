/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Monday 10 February 2020 - 10:52:20
** @Filename:				install.go
**
** @Last modified by:		Tbouder
** @Last modified time:		Monday 10 February 2020 - 14:48:59
*******************************************************************************/

package			main

import			"os"
import			"crypto/rand"
import			"encoding/base64"

func	generateNonce(n uint32) (string) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if (err != nil) {
        return ``
	}
    ciphertext := base64.RawStdEncoding.EncodeToString(b)
	
    return ciphertext
}

func	main()	{
	f, err := os.Create(".env")
	if (err != nil) {
		panic(err)
	}
	defer f.Close()
	f.WriteString("POSTGRE_USERNAME=" + `Panghostlin` + "\n")
	f.WriteString("POSTGRE_PWD=" + generateNonce(64) + "\n")
	f.WriteString("POSTGRE_URI=" + `panghostlin-postgre` + "\n")
	f.WriteString("POSTGRE_PORT=" + `54320` + "\n")
	f.WriteString("POSTGRE_DB=" + `panghostlin` + "\n")
	f.WriteString("MASTER_KEY=" + generateNonce(32) + "\n")
	f.WriteString("MASTER_PUBLIC_KEY=" + generateNonce(32) + "\n")
	f.WriteString("MASTER_KEY_ARGON2=" + generateNonce(32) + "\n")
	f.WriteString("MASTER_KEY_SCRYPT=" + generateNonce(32) + "\n")
	f.WriteString("PRIV_KEY=" + generateNonce(64) + "\n")
	f.WriteString("JWT_ACCESS_TOKEN_KEY=" + generateNonce(64) + "\n")
	f.WriteString("JWT_REFRESH_TOKEN_KEY=" + generateNonce(64) + "\n")
}
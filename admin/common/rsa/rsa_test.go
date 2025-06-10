package rsa

import (
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSA(t *testing.T) {
	privateKey, err := GenerateKey(2048)
	assert.Nil(t, err)

	data := "abcdABCD1234.?/"
	encryptedData, err := Encrypt(data, &privateKey.PublicKey)
	assert.Nil(t, err)

	decryptedData, err := Decrypt(encryptedData, privateKey)
	assert.Nil(t, err)

	assert.Equal(t, data, decryptedData)

	// When the message is too long
	longData := "22069666451008748339503251759148525361164075594487536910385540351176004020504149174230524466483286357625601955249461631036940045196545548748998825476551946788939131757441090884176777283162704817836941071232796987975103116207838036073500403731065117309971693587920709418245140569016005687441340339334167771092774976541127041894556333293062931239505390395517887225078908444644452761548061633204874030229350165442244597118226103203268940371930572470674294929590669333987072539993344600214821864303609288225145665921540649609836791758252738759040105728602555789728921129366844946996496056111759814690356832020328405656999eiojfkdlsklnlmcls;jfdoiewuru"

	_, err = Encrypt(longData, &privateKey.PublicKey)
	assert.Equal(t, rsa.ErrMessageTooLong, err)

	// When encryptedData is tampered
	attckWords := hex.EncodeToString([]byte("attack"))
	encryptedDataTampered := fmt.Sprintf("%s%s", encryptedData, attckWords)

	_, err = Decrypt(encryptedDataTampered, privateKey)
	assert.NotNil(t, err)

	// When private key is not correct
	newPrivateKey, err := GenerateKey(2048)
	assert.Nil(t, err)

	_, err = Decrypt(encryptedData, newPrivateKey)
	assert.Equal(t, rsa.ErrDecryption, err)
}

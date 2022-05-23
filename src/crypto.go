package src

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func getKey(password string) ([]byte, error) {
	keyFile, err := readFile(KEY_PATH)
	if os.IsNotExist(err) {
		key, err := CreateKey(password)
		if err != nil {
			return nil, fmt.Errorf("error creating new key")
		}

		err = writeFile(KEY_PATH, key)
		if err != nil {
			return nil, err
		}

		return getKey(password)
	}
	if err != nil {
		return nil, err
	}

	passwordHash := pbkdf2.Key([]byte(password), make([]byte, HASH_SIZE), SALT_SIZE, KEY_SIZE, sha1.New)
	block, err := aes.NewCipher([]byte(passwordHash))
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, KEY_SIZE)

	cfb := cipher.NewCBCDecrypter(block, make([]byte, HASH_SIZE))
	cfb.CryptBlocks(decrypted, keyFile)
	return decrypted, nil
}

// CreateKey created a new key encrypted with the given password
func CreateKey(password string) ([]byte, error) {
	passwordHash := pbkdf2.Key([]byte(password), make([]byte, HASH_SIZE), SALT_SIZE, KEY_SIZE, sha1.New)

	block, err := aes.NewCipher([]byte(passwordHash))
	if err != nil {
		return nil, err
	}

	key := make([]byte, KEY_SIZE)
	_, err = rand.Read(key) // Randomize the key bytes
	if err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, make([]byte, HASH_SIZE))
	cipherText := make([]byte, len(key))
	cfb.XORKeyStream(cipherText, key)

	return cipherText, nil
}

func encrypt(data []byte, password string) ([]byte, error) {
	key, err := getKey(password)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)

	return cipherText, nil
}

func decrypt(cipherText []byte, password string) ([]byte, error) {
	key, err := getKey(password)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]

	decoded, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

// EncryptPath takes a srcPath and creates a encrypted version at destPath
func CryptPath(srcPath, destPath, password string, decrypt bool) error {
	infile, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	defer infile.Close()

	key, err := getKey(password)
	if err != nil {
		return err
	}

	_, err = rand.Read(key)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	var iv []byte = make([]byte, block.BlockSize())

	if !decrypt {
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return err
		}
	} else {
		fi, err := infile.Stat()
		if err != nil {
			return err
		}

		msgLen := fi.Size() - int64(len(iv))
		_, err = infile.ReadAt(iv, msgLen)
		if err != nil {
			return err
		}
	}

	outfile, err := os.OpenFile(destPath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	buff := make([]byte, FILE_BUFF_SIZE)
	stream := cipher.NewCTR(block, iv)

	for {
		n, err := infile.Read(buff)
		if n > 0 {
			stream.XORKeyStream(buff, buff[:n])
			outfile.Write(buff[:n])
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	_, err = outfile.Write(iv)
	if err != nil {
		return err
	}

	return outfile.Chmod(READ_ONLY_FILE_MODE)
}

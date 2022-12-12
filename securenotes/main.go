package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AccountData struct {
	Username string   `json:"username,omitempty"`
	Password string   `json:"password,omitempty"`
	Notes    []string `json:"notes,omitempty"`
}

var jsonPath = "./container.json"

func getAccounts() (data map[string]AccountData) {
	file, err := os.OpenFile(jsonPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalln(err)
		return
	}

	byteValue, _ := ioutil.ReadAll(file)

	defer file.Close()

	if len(byteValue) <= 0 {
		json.Unmarshal([]byte("{}"), &data)
		return
	}

	json.Unmarshal(byteValue, &data)

	return
}

func getAccountData(username, password string) (data AccountData) {
	accounts := getAccounts()

	data, exists := accounts[username]
	if !exists {
		data = AccountData{
			Username: username,
			Password: password,
			Notes:    []string{},
		}
	}

	if data.Password != password {
		data = AccountData{}
	}

	return
}

func updateAccountFile(data AccountData) {
	accounts := getAccounts()
	accounts[data.Username] = data
	file, _ := json.MarshalIndent(accounts, "", "    ")

	_ = ioutil.WriteFile(jsonPath, file, 0644)
}

func main() {
	var usernameInput string
	var passwordInput string
	var accountData AccountData
	var cipher cipher.Block

	hasher := md5.New()

	a := app.New()
	window := a.NewWindow("Test program REEEEEEEEEEE")

	hello := widget.NewLabel("HIIIII REEEE")

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Input username here!")

	passwordEntry := widget.NewEntry()
	passwordEntry.SetPlaceHolder("Input password here!")
	passwordEntry.Password = true

	addNoteEntry := widget.NewEntry()
	addNoteEntry.SetPlaceHolder("Input note here!")
	addNoteEntry.MultiLine = true

	addNoteButton := widget.NewButton("Add note", func() {
		note := addNoteEntry.Text

		if len(note) <= 0 {
			return
		}

		encryptedNote := EncryptAES(cipher, note)

		accountData.Notes = append(accountData.Notes, encryptedNote)
		updateAccountFile(accountData)
	})

	noteList := widget.NewList(
		func() int {
			return len(accountData.Notes)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			decryptedText := DecryptAES(cipher, accountData.Notes[i])
			o.(*widget.Label).SetText(decryptedText)
		},
	)

	acceptButton := widget.NewButton("Accept", func() {
		usernameInput = usernameEntry.Text
		passwordInput = passwordEntry.Text

		if len(usernameInput) <= 0 {
			hello.SetText("Username must be atleast 1 character!")
			time.Sleep(5 * time.Second)
			log.Fatalln("Invalid username length!")
			return
		}

		if len(passwordInput) <= 0 {
			hello.SetText("Password must be atleast 1 character!")
			time.Sleep(5 * time.Second)
			log.Fatalln("Invalid password length!")
			return
		}

		_, err := hasher.Write([]byte(passwordInput))
		hashedPassword := hex.EncodeToString(hasher.Sum(nil))

		if err != nil {
			log.Fatalln(err)
			return
		}

		accountData = getAccountData(usernameInput, hashedPassword)
		password := accountData.Password

		if password != hashedPassword {
			hello.SetText("Invalid password!")
			time.Sleep(5 * time.Second)
			log.Fatalln("Invalid password!")
			return
		}

		cipher, err = aes.NewCipher([]byte(hashedPassword))
		if err != nil {
			log.Fatalln(err)
			return
		}

		window.SetContent(container.NewVBox(
			addNoteEntry,
			addNoteButton,
			noteList,
		))

	})

	window.SetContent(container.NewVBox(
		hello,
		usernameEntry,
		passwordEntry,
		acceptButton,
	))

	window.Resize(fyne.NewSize(500, 300))
	window.ShowAndRun()
}

func EncryptAES(cipherBlock cipher.Block, plaintext string) string {
	out := make([]byte, aes.BlockSize+len(plaintext))

	iv := out[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalln(err)
		return ""
	}

	stream := cipher.NewCFBEncrypter(cipherBlock, iv)
	stream.XORKeyStream(out[aes.BlockSize:], []byte(plaintext))

	return base64.RawStdEncoding.EncodeToString(out)
}

func DecryptAES(cipherBlock cipher.Block, encrypted string) string {
	cipherText, err := base64.RawStdEncoding.DecodeString(encrypted)

	if err != nil {
		log.Fatalln(err)
		return ""
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(cipherBlock, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}

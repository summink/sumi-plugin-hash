package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"

	"github.com/InkShaStudio/go-command"
	"github.com/atotto/clipboard"
)

func getContent(isFile bool, value string) (io.Reader, error) {
	if isFile {
		return os.Open(value)
	}
	return strings.NewReader(value), nil
}

type NotSupportHashMode struct {
	Message string
}

func (e *NotSupportHashMode) Error() string {
	return e.Message
}

func computedHash(content io.Reader, mode string) (string, error) {
	var hashMode hash.Hash

	switch mode {
	case "md5":
		hashMode = md5.New()
	case "sha1":
		hashMode = sha1.New()
	case "sha256":
		hashMode = sha256.New()
	case "sha512":
		hashMode = sha512.New()
	default:
		return "", &NotSupportHashMode{Message: fmt.Sprintf(`"%s" Is Invalid hash mode`, mode)}
	}

	io.Copy(hashMode, content)

	hashContent := fmt.Sprintf("%x", hashMode.Sum(nil))

	return hashContent, nil
}

func registerCommand() *command.SCommand {
	f := command.NewCommandFlag[bool]("file").ChangeDescription("The file to hash").ChangeValue(false)
	m := command.NewCommandFlag[string]("mode").ChangeDescription("The hash method").ChangeValue("sha256")
	c := command.NewCommandFlag[bool]("copy").ChangeDescription("Copy hash content to clipboard").ChangeValue(false)
	v := command.NewCommandFlag[string]("verify").ChangeDescription("verify hash is equal")

	value := command.NewCommandArg[string]("value").ChangeDescription("Compute hash content")

	cmd := command.
		NewCommand("hash").
		ChangeDescription("Compute hash content").
		AddArgs(value).
		AddFlags(f, m, c, v).
		RegisterHandler(func(cmd *command.SCommand) {
			content, err := getContent(f.Value, value.Value)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			hashContent, err := computedHash(content, m.Value)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if v.Value != "" {
				if hashContent != v.Value {
					fmt.Println("Hash content not equal verify content")
					os.Exit(1)
				} else {
					fmt.Println("Hash content equal verify content")
				}
			} else {
				fmt.Println(hashContent)

				if c.Value {
					clipboard.WriteAll(hashContent)
				}
			}
		})

	return cmd
}

func main() {
	cmd := command.RegisterCommand(registerCommand())
	cmd.Execute()
}

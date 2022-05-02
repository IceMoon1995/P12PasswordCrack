package main

import (
	"PP12PasswordCrack/pkcs12"
	_ "crypto/ecdsa"
	"os"
	"strings"
)

func main() {
	args := os.Args
	println()
	skFile2, _ := os.Open(args[2])
	fileinfo2, _ := skFile2.Stat()
	skBytes2 := make([]byte, fileinfo2.Size())
	skFile2.Read(skBytes2)
	skBytes2str := string(skBytes2)
	skBytes2List := strings.Split(skBytes2str, "\r\n")

	for _, pass := range skBytes2List {
		skFile, err := os.Open(args[1])
		fileinfo, err := skFile.Stat()
		skBytes := make([]byte, fileinfo.Size())
		skFile.Read(skBytes)
		AA, bb, err := pkcs12.DecodeAll(skBytes, pass)
		if err != nil {
			//println("error")
		} else if pass != "" {
			println(pass)
			break
		} else {
			println(AA)
			println(bb)
		}
	}

}

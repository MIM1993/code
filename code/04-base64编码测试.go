package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	src := []byte("两情若是久长时，又岂在朝朝暮暮")

	//标准
	encodeInfo := base64.StdEncoding.EncodeToString(src)
	fmt.Println("encodeInfo : %v\n", encodeInfo)

	//URL专用
	URLencodeInfo := base64.URLEncoding.EncodeToString(src)
	fmt.Println("URLencodeInfo : %v\n", URLencodeInfo)
}

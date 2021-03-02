package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

// go 加解密处理（sha、md5 等）
// 标准库加解-编解码包位于 encoding 包下：https://golang.org/pkg/encoding/

// hash 包：实现了 adler32、crc32、crc64 和 fnv 校验；
// crypto 包：实现了其它的 hash 算法，比如 md4、md5、sha1 等。
//			  以及完整地实现了 aes、blowfish、rc4、rsa、xtea 等加密算法。

func main() {

	str1 := "Whenever you need me, I'll be here."
	str2 := "A sonnet in every tree, a tale in every lifetime, its just for you to see…"

	// 计算 Hash SHA1 校验值
	shaHash := sha1.New()
	b := []byte{}

	io.WriteString(shaHash, str1)
	checksum1 := shaHash.Sum(b)   // 计算 str1 sha1 校验和
	fmt.Printf("%X\n", checksum1) // D57ACBB1116583070B11D8206FC5FC7B25FC6E52

	shaHash.Reset()
	io.WriteString(shaHash, str2)
	checksum2 := shaHash.Sum(b)   // 计算 str2 sha1 校验和
	fmt.Printf("%X\n", checksum2) // 372547E7132D000BE9C515C1FECDF8FD2C56B15E

	// 计算 Hash MD5 校验值
	md5Hash := md5.New()
	io.WriteString(md5Hash, str1)
	checksum3 := md5Hash.Sum(b)   // 计算 str1 md5 校验和
	fmt.Printf("%X\n", checksum3) // 6B5203F0EC116B97B2A9F5811DC26807

}

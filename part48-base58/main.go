package main

import (
	"fmt"
	"crypto/sha256"
	"kongyixueyuan.com/publicChain/part48-base58/BLC"
)

func main()  {

	bytes := []byte("http://liyuechun.orghttp://liyuechun.orghttp://liyuechun.org")

	hasher := sha256.New()
	hasher.Write(bytes)
	hash := hasher.Sum(nil)
	//7b28134d636423d716aaa45227099629783ae6d6012b331049466658b88bf3b5
	//62f556541508147eb4ef2d015a943cfb56f9a65a989ba40b7c3f51491d62a7c7

	fmt.Printf("%x\n",hash)
	//
	bytes58 := BLC.Base58Encode(hash)

	//fmt.Printf("%x\n",bytes58)
	//12TQQJ2URa4LWVGUKRWw8MJvahxzJ
	//1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa

	fmt.Printf("%s\n",bytes58)
	// 19HkbCnjQzGb3hcG1xNFkXBbE5v8FrSoKaywutCgNx3gG
	// 17fHw3nUseFctpqUN6EQ5gsGTfEwwgxiK97WJBF3mcz8v


	bytesStr := BLC.Base58Decode(bytes58)

	fmt.Printf("%x\n",bytesStr[1:])


}

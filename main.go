package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "csrc/increase.c"
#include "csrc/ed25519.h"
#include "csrc/precomp_data.h"
#include "csrc/sha512.c"
#include "csrc/sc.c"
#include "csrc/ge.c"
#include "csrc/fe.c"
#include "csrc/ed25519.c"
*/
import "C"
import (
	"fmt"
	"time"
)

func main() {
	seed := make([]uint8, 32)
	publicKey := make([]uint8, 32)
	privateKey := make([]uint8, 64)
	signature := make([]uint8, 64)
	message := make([]uint8, 16)
	for index := range message {
		message[index] = uint8(index)
	}
	C.ed25519_create_keypair(tranformToUchar(publicKey), tranformToUchar(privateKey), tranformToUchar(seed))
	messageLength := C.ulong(len(message))
	C.ed25519_sign(tranformToUchar(signature), tranformToUchar(message), messageLength, tranformToUchar(publicKey), tranformToUchar(privateKey))
	/* verify the signature */
	if C.ed25519_verify(tranformToUchar(signature), tranformToUchar(message), messageLength, tranformToUchar(publicKey)) == 1 {
		fmt.Println("valid signature")
	} else {
		fmt.Println("invalid signature")
	}
	start := time.Now()
	for i := 0; i < 10000; i++ {
		C.ed25519_verify(tranformToUchar(signature), tranformToUchar(message), messageLength, tranformToUchar(publicKey))
	}
	end := time.Now()
	fmt.Println("each signature verify cost ", end.Sub(start)/10000)
}

func tranformToUchar(src []uint8) *C.uchar {
	//return (*C.uchar)(unsafe.Pointer(&src[0]))
	return (*C.uchar)(&src[0])
}

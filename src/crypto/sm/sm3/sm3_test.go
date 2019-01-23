/*
Copyright Suzhou Tongji Fintech Research Institute 2017 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sm3

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type sm3Test struct {
	out string
	in  string
}

var plainText = []sm3Test{
	{"55e12e91650d2fec56ec74e1d3e4ddbfce2ef3a65890c2a19ecf88a307e76a23", "test"},
	{"1ab21d8355cfa17f8e61194831e81a8f22bec8c728fefb747ed035eb5082aa2b", ""},
	{"00607cd3ffb78125184758bb06d23757beb3d57be9447b1bb58a6a6e67752313", "616263"},
	{"623476ac18f65a2909e43c7fec61b49c7e764a91a18ccb82f1917a29c86c5e88", "a"},
	{"2bb6c53ad20eaf2552425f44e72d96d1b61e63310a1a30f4e5406a103619177d", "He who has a shady past knows that nice guys finish last."},
	{"5ecec640017afd77d00147ef42fdb8e7901f089a62c1888637917e89bb3a6532", "I wouldn't marry him with a ten foot pole."},
	{"26598310dfeea2787829ec21d88fbf9f17c9299adf23de49cfcf26030dbc0e35", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	{"c3555aaf32465c61f681e6dabcc0c95ac93e7c383b1c6eeb621a5ca0eb300508", "The days of the digital watch are numbered.  -Tom Stoppard"},
}

func byteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	//fmt.Println("ret = ", ret)
	return ret
}

func TestSm3(t *testing.T) {
	for i := 0; i < len(plainText); i++ {
		msg := []byte(plainText[i].in)
		err := ioutil.WriteFile("ifile", msg, os.FileMode(0644)) // 生成测试文件
		if err != nil {
			log.Fatal(err)
		}
		msg, err = ioutil.ReadFile("ifile")
		if err != nil {
			log.Fatal(err)
		}
		hw := New()
		hw.Write(msg)
		hash := hw.Sum(nil)
		//fmt.Println(hash)
		//fmt.Printf("hash = %d\n", len(hash))
		hashString := byteToString(hash)
		fmt.Printf("%s\n", hashString)
		if hashString != plainText[i].out {
			t.Fatalf("SumSM3 function: SM3(%s) = %s want %s", plainText[i].in, hashString, plainText[i].out)
		}
		hash1 := Sm3Sum(msg)
		//fmt.Println(hash1)
		fmt.Printf("%s\n", byteToString(hash1))
	}
}

func BenchmarkSm3(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("test1234")
	hw := New()
	for i := 0; i < t.N; i++ {
		hw.Sum(nil)
		Sm3Sum(msg)
	}
}

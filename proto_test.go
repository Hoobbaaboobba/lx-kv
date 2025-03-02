package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/tidwall/resp"
)

func TestProtocol(t *testing.T) {
    raw := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"

    rd := resp.NewReader(bytes.NewBufferString(raw))

    for {
        v, _, err := rd.ReadValue()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Read %s\n", v.Type())

        var cmd Command
        if v.Type() == resp.Array {
            for _, v := range v.Array() {
                switch v.String() {
                case CommandSET:
                    if len(v.Array()) != 3 {
                        panic("yikes")
                    }
                    cmd = SetCommand{
                        key: v.Array()[1].String(),
                        val: v.Array()[2].String(),
                    }
                default:
                }
                if v.String() == CommandSET {
                    panic("works")
                }
                fmt.Printf("%v\n", v)
            }
        }
    }
}

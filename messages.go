package gocec

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type Message []byte

func ParseMessage(message string) (Message, error) {
	message = strings.ReplaceAll(message, ":", "")

	command, err := hex.DecodeString(message)
	if err != nil {
		return Message{}, err
	}

	return Message(command), nil
}

func (message Message) Source() LogicalAddress {
	return LogicalAddress(message[0] >> 4)
}

func (message Message) Destination() LogicalAddress {
	return LogicalAddress(0x0F & message[0])
}

func (message Message) Opcode() Opcode {
	if len (message) < 2 {
		return OpcodeNone
	}

	return Opcode(message[1])
}

func (message Message) Parameters() []byte {
	if len(message) < 3 {
		return nil
	}

	return message[2:]
}

func (message Message) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "from %s to %s, message: %s", message.Source(), message.Destination(), message.Opcode())

	if len(message) > 2 {
		fmt.Fprintf(&b, ", with parameters: %X", message.Parameters())
	}

	return b.String()
}

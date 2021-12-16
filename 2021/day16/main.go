package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occured: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	input, err := inputFromString(string(inputFile))
	if err != nil {
		return err
	}

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))

	return nil
}

func inputFromString(inputStr string) (string, error) {
	binBuilder := strings.Builder{}
	for _, hexVal := range inputStr {
		valToAdd := ""
		switch hexVal {
		case '0':
			valToAdd = "0000"
		case '1':
			valToAdd = "0001"
		case '2':
			valToAdd = "0010"
		case '3':
			valToAdd = "0011"
		case '4':
			valToAdd = "0100"
		case '5':
			valToAdd = "0101"
		case '6':
			valToAdd = "0110"
		case '7':
			valToAdd = "0111"
		case '8':
			valToAdd = "1000"
		case '9':
			valToAdd = "1001"
		case 'A':
			valToAdd = "1010"
		case 'B':
			valToAdd = "1011"
		case 'C':
			valToAdd = "1100"
		case 'D':
			valToAdd = "1101"
		case 'E':
			valToAdd = "1110"
		case 'F':
			valToAdd = "1111"
		}
		binBuilder.WriteString(valToAdd)
	}

	return binBuilder.String(), nil
}

type Packet struct {
	version   int64
	typeID    int64
	value     int
	contained []Packet
}

func valueBitsToNumber(bin string) (int, int) {
	startPos := 0
	numberString := ""
	for bin[startPos:startPos+1] == "1" {
		numberString += bin[startPos+1 : startPos+5]
		startPos += 5
	}
	// read zero bit once again
	numberString += bin[startPos+1 : startPos+5]

	return int(parseBin(numberString)), startPos + 5
}

func valueBitsToPackets(bin string) ([]Packet, int) {
	var packets []Packet
	if bin[0:1] == "0" {
		// next 15 contain bits of all subpackets
		totalBits := int(parseBin(bin[1:16]))
		currentPos := 16
		targetPos := currentPos + totalBits
		// at least 6 bits needed for header
		for targetPos-currentPos > 6 {
			p, length := packetFromStr(bin[currentPos:])
			packets = append(packets, p)
			currentPos += length
		}

		return packets, totalBits + 16
	}

	// next 11 bits total number of subpackets
	totalPackets := int(parseBin(bin[1:12]))
	currentPos := 12

	for i := 0; i < totalPackets; i++ {
		p, length := packetFromStr(bin[currentPos:])
		packets = append(packets, p)
		currentPos += length
	}

	return packets, currentPos
}

func packetFromStr(bin string) (Packet, int) {
	version := parseBin(bin[:3])
	typeID := parseBin(bin[3:6])

	if typeID == 4 {
		number, length := valueBitsToNumber(bin[6:])
		return Packet{
			version: version,
			typeID:  typeID,
			value:   number,
		}, length + 6
	}

	contained, length := valueBitsToPackets(bin[6:])
	return Packet{
		version:   version,
		typeID:    typeID,
		contained: contained,
	}, length + 6
}

func SolvePart1(input string) int {
	versionSum := int(parseBin(input[:3]))
	typeID := parseBin(input[3:6])

	if typeID == 4 {
		panic("nope")
	}

	packets, _ := valueBitsToPackets(input[6:])

	for _, p := range packets {
		versionSum += sumVersion(p)
	}

	return versionSum
}

func sumVersion(packet Packet) int {
	sum := int(packet.version)
	for _, p := range packet.contained {
		sum += sumVersion(p)
	}

	return sum
}

func parseBin(in string) int64 {
	res, err := strconv.ParseInt(in, 2, 0)
	if err != nil {
		panic("test")
	}

	return res
}

func SolvePart2(input string) int {
	version := parseBin(input[:3])
	typeID := parseBin(input[3:6])

	packets, _ := valueBitsToPackets(input[6:])

	rootPacket := Packet{
		version:   int64(version),
		typeID:    typeID,
		contained: packets,
	}

	return applyOperation(rootPacket)
}

func applyOperation(packet Packet) int {
	switch packet.typeID {
	case 0:
		sum := 0
		for _, subp := range packet.contained {
			sum += applyOperation(subp)
		}
		return sum
	case 1:
		product := 1
		for _, subp := range packet.contained {
			product *= applyOperation(subp)
		}
		return product
	case 2:
		min := math.MaxInt
		for _, subp := range packet.contained {
			if res := applyOperation(subp); res < min {
				min = res
			}
		}
		return min
	case 3:
		max := 0
		for _, subp := range packet.contained {
			if res := applyOperation(subp); res > max {
				max = res
			}
		}
		return max
	case 4:
		return packet.value
	case 5:
		if applyOperation(packet.contained[0]) > applyOperation(packet.contained[1]) {
			return 1
		}
		return 0
	case 6:
		if applyOperation(packet.contained[0]) < applyOperation(packet.contained[1]) {
			return 1
		}
		return 0
	case 7:
		if applyOperation(packet.contained[0]) == applyOperation(packet.contained[1]) {
			return 1
		}
		return 0
	}
	panic("unknown type id")
}

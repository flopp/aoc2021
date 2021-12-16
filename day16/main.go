package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/flopp/aoc2021/helpers"
)

type BITS struct {
	transmission []uint8
	pos          int
}

func (bits *BITS) append(b0, b1, b2, b3 uint8) {
	bits.transmission = append(bits.transmission, b0, b1, b2, b3)
}

func (bits *BITS) get(count int) uint64 {
	if count < 0 || count > 64 {
		panic("bad count")
	}
	if count > len(bits.transmission)-bits.pos {
		panic("eot")
	}

	value := uint64(0)
	for i := 0; i < count; i++ {
		value = (value << 1) | uint64(bits.transmission[bits.pos])
		bits.pos++
	}
	return value
}

func (bits *BITS) parseHeader() PacketHeader {
	return PacketHeader{bits.get(3), bits.get(3)}
}

func (bits *BITS) parseLiteralValue() uint64 {
	value := uint64(0)
	for true {
		first := bits.get(1)
		value = value<<4 | bits.get(4)
		if first == 0 {
			break
		}
	}
	return value
}

func (bits *BITS) parseLiteralPacket(header PacketHeader) Packet {
	return LiteralPacket{header, bits.parseLiteralValue()}
}

func (bits *BITS) parseOperatorPacket(header PacketHeader) Packet {
	subpackets := []Packet{}
	lengthTypeID := bits.get(1)
	if lengthTypeID == 0 {
		length := bits.get(15)
		endPos := bits.pos + int(length)
		for bits.pos < endPos {
			subpackets = append(subpackets, bits.parsePacket())
		}
	} else {
		count := bits.get(11)
		for i := uint64(0); i < count; i++ {
			subpackets = append(subpackets, bits.parsePacket())
		}
	}
	return OperatorPacket{header, subpackets}
}

func (bits *BITS) parsePacket() Packet {
	header := bits.parseHeader()
	switch header.typeID {
	case 4:
		return bits.parseLiteralPacket(header)
	default:
		return bits.parseOperatorPacket(header)
	}
}

type Packet interface {
	versionSum() uint64
	value() uint64
}

type PacketHeader struct {
	version uint64
	typeID  uint64
}

type LiteralPacket struct {
	header       PacketHeader
	literalValue uint64
}

func (packet LiteralPacket) versionSum() uint64 {
	return packet.header.version
}

func (packet LiteralPacket) value() uint64 {
	return packet.literalValue
}

type OperatorPacket struct {
	header     PacketHeader
	subpackets []Packet
}

func (packet OperatorPacket) versionSum() uint64 {
	version := packet.header.version
	for _, subpacket := range packet.subpackets {
		version += subpacket.versionSum()
	}
	return version
}

func (packet OperatorPacket) value() uint64 {
	value := uint64(0)

	switch packet.header.typeID {
	case 0:
		for _, p := range packet.subpackets {
			value += p.value()
		}
	case 1:
		value = uint64(1)
		for _, p := range packet.subpackets {
			value *= p.value()
		}
	case 2:
		value = uint64(math.MaxUint64)
		for _, p := range packet.subpackets {
			pv := p.value()
			if pv < value {
				value = pv
			}
		}
	case 3:
		for _, p := range packet.subpackets {
			pv := p.value()
			if pv > value {
				value = pv
			}
		}
	case 5:
		if packet.subpackets[0].value() > packet.subpackets[1].value() {
			value = uint64(1)
		}
	case 6:
		if packet.subpackets[0].value() < packet.subpackets[1].value() {
			value = uint64(1)
		}
	case 7:
		if packet.subpackets[0].value() == packet.subpackets[1].value() {
			value = uint64(1)
		}
	default:
		panic("bad typeID")
	}

	return value
}

func main() {
	bits := BITS{}
	helpers.ReadStdin(func(line string) {
		for _, c := range line {
			hex, _ := strconv.ParseUint(string(c), 16, 8)
			b := uint8(hex)
			bits.append((b>>3)&1, (b>>2)&1, (b>>1)&1, b&1)
		}
	})

	packet := bits.parsePacket()
	if helpers.Part1() {
		fmt.Println(packet.versionSum())
	} else {
		fmt.Println(packet.value())
	}
}

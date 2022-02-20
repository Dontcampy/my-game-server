package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"github.com/dontcampy/my-game-server/mygameserver/utils"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeaderLength() uint32 {
	// id(uint32) + length(uint32)
	return 8
}

// Pack Message |header|data|
func (dp *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	// pack length
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetHeader().ContentLength)
	if err != nil {
		return nil, err
	}

	// pack id
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetHeader().Id)
	if err != nil {
		return nil, err
	}

	// pack data
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetContent())
	if err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (dp *DataPack) UnpackHeader(binaryData []byte) (iface.Header, error) {
	dataBuff := bytes.NewReader(binaryData)

	// unpack length
	var length uint32
	err := binary.Read(dataBuff, binary.LittleEndian, &length)
	if err != nil {
		return iface.Header{}, err
	}

	if length > utils.GlobalObject.MaxPackageSize {
		return iface.Header{}, errors.New("too Large msg data recv")
	}

	// unpack id
	var id uint32
	err = binary.Read(dataBuff, binary.LittleEndian, &id)
	if err != nil {
		return iface.Header{}, err
	}

	return iface.Header{ContentLength: length, Id: id}, nil
}

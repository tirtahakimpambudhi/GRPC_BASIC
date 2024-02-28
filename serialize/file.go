package serialize

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

type File struct {
	FileName     string
	FileNameJSON string
	Message      proto.Message
}

func (f *File) ProtoBufToJSON() (string, error) {
	marshaler := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
	}
	return marshaler.MarshalToString(f.Message)
}

func (f *File) WriteFileJSON() error {
	file, err := os.Create(f.FileNameJSON)
	if err != nil {
		return fmt.Errorf("Error Create File : %v ", err.Error())
	}
	defer file.Close()
	json, err := f.ProtoBufToJSON()
	if err != nil {
		return fmt.Errorf("Error Parse JSON : %v ", err.Error())
	}
	_, err = file.Write([]byte(json))
	if err != nil {
		return fmt.Errorf("Error Write JSON : %v ", err.Error())
	}
	bufio.NewWriter(file)
	return nil
}

func (f *File) WriteFile() error {

	file, err := os.Create(f.FileName)
	if err != nil {
		return fmt.Errorf("Error Create File : %v ", err.Error())
	}
	defer file.Close()
	data, err := proto.Marshal(f.Message)
	if err != nil {
		return fmt.Errorf("Error Marshal Proto: %v", err.Error())
	}
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("Error Write File : %v", err.Error())
	}
	if err != nil {
		return err
	}
	bufio.NewWriter(file)
	return nil
}

func (f *File) ReadFile(message proto.Message) error {
	file, err := ioutil.ReadFile(f.FileName)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(file, message)
	if err != nil {
		return err
	}
	return nil
}

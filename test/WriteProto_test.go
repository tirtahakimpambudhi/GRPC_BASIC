package test

import (
	"github.com/stretchr/testify/require"
	"grpc_course/factory"
	"grpc_course/pb"
	"grpc_course/pkg"
	"grpc_course/serialize"
	"os"
	"path"
	"testing"
)

func TestWriteFile(t *testing.T) {
	workdir, _ := os.Getwd()
	filename := path.Join(workdir, "..", "tmp", "bin", "pc_book.bin")
	filenameJSON := path.Join(workdir, "..", "tmp", "json", "pc_book.json")
	pc1 := factory.NewLaptop()
	pc2 := &pb.Laptop{}
	file := serialize.File{
		FileNameJSON: filenameJSON,
		FileName:     filename,
		Message:      pc1,
	}
	err := file.WriteFile()
	require.NoError(t, err)
	err = file.ReadFile(pc2)
	require.NoError(t, err)
	err = file.WriteFileJSON()
	require.NoError(t, err)
	jsonpc1, err := pkg.ProtoBufToJSON(pc1)
	require.NoError(t, err)
	jsonpc2, err := pkg.ProtoBufToJSON(pc2)
	require.NoError(t, err)
	require.Equal(t, jsonpc1, jsonpc2)
}

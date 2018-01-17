package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/thrift-iterator/go/protocol"
)

func Test_skip_list_of_struct(t *testing.T) {
	should := require.New(t)
	buf := thrift.NewTMemoryBuffer()
	proto := thrift.NewTBinaryProtocol(buf, true, true)
	proto.WriteListBegin(thrift.STRUCT, 2)
	proto.WriteStructBegin("hello")
	proto.WriteFieldBegin("field1", thrift.I64, 1)
	proto.WriteI64(1024)
	proto.WriteFieldEnd()
	proto.WriteFieldStop()
	proto.WriteStructEnd()
	proto.WriteStructBegin("hello")
	proto.WriteFieldBegin("field1", thrift.I64, 1)
	proto.WriteI64(1024)
	proto.WriteFieldEnd()
	proto.WriteFieldStop()
	proto.WriteStructEnd()
	proto.WriteListEnd()
	iter := thrifter.NewIterator(buf.Bytes())
	should.Equal(buf.Bytes(), iter.SkipList())
}

func Test_decode_list_of_struct(t *testing.T) {
	should := require.New(t)
	buf := thrift.NewTMemoryBuffer()
	proto := thrift.NewTBinaryProtocol(buf, true, true)
	proto.WriteListBegin(thrift.STRUCT, 2)
	proto.WriteStructBegin("hello")
	proto.WriteFieldBegin("field1", thrift.I64, 1)
	proto.WriteI64(1024)
	proto.WriteFieldEnd()
	proto.WriteFieldStop()
	proto.WriteStructEnd()
	proto.WriteStructBegin("hello")
	proto.WriteFieldBegin("field1", thrift.I64, 1)
	proto.WriteI64(1024)
	proto.WriteFieldEnd()
	proto.WriteFieldStop()
	proto.WriteStructEnd()
	proto.WriteListEnd()
	iter := thrifter.NewIterator(buf.Bytes())
	should.Equal(map[protocol.FieldId]interface{}{
		protocol.FieldId(1): int64(1024),
	}, iter.ReadList()[0])
}

func Test_encode_list_of_struct(t *testing.T) {
	should := require.New(t)
	stream := thrifter.NewStream(nil)
	stream.WriteList([]interface{}{
		map[protocol.FieldId]interface{}{
			protocol.FieldId(1): int64(1024),
		},
		map[protocol.FieldId]interface{}{
			protocol.FieldId(1): int64(1024),
		},
	})
	iter := thrifter.NewIterator(stream.Buffer())
	should.Equal(map[protocol.FieldId]interface{}{
		protocol.FieldId(1): int64(1024),
	}, iter.ReadList()[0])
}

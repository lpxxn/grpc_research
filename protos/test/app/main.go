package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func main() {
	// StartTime:
	//r := &a.Result{Name: "aaa", D: &a.D{StartTime: ptypes.TimestampNow()}}
	//fmt.Println(r)
	//m, _ := proto.Marshal(r)
	//fmt.Println(string(m))
	//jb, _ := json.Marshal(r)
	//fmt.Println(string(jb))
	//fmt.Println(json.Unmarshal(jb, r))
	msgFactory := dynamic.NewMessageFactoryWithDefaults()
	//
	//rm := r.ProtoReflect()
	//fmt.Println(json.Unmarshal(jb, rm))
	fmt.Println(time.Now())
	parser := protoparse.Parser{
	}
	fds, err := parser.ParseFiles("./a.proto")
	if err != nil {
		panic(err)
	}
	registerFileDescriptors(fds)

	msg := dynamic.NewMessage(fds[0].FindMessage("protos.Result"))

	msgJb, err := json.Marshal(msg)

	fmt.Println(string(msgJb))
	msgJb, err = msg.MarshalJSONPB(&jsonpb.Marshaler{
		Indent:       "",
		EmitDefaults: true,
	})
	fmt.Println(string(msgJb))

	for _, sd := range fds[0].GetServices() {
		pm := msgFactory.NewMessage(sd.GetMethods()[0].GetOutputType())
		fmt.Println(pm)
		fmt.Println(json.Unmarshal([]byte(`{"name":"aaa","d":{"start_time":{"seconds":1623925088,"nanos":206250000}}}`), pm))
		//fmt.Println(json.Unmarshal([]byte(`{"name":"aaa","d":{"start_time":"2021-06-17T18:23:59.919014Z"}}`), pm))
		fmt.Println(pm)
	}
	return
	//pm := msgFactory.NewMessage(fds[0].FindMessage("protos.Result"))
	//pm := msgFactory.NewMessage(fds[0].GetMessageTypes()[0])
	//fmt.Println(pm)
	//fmt.Println(time.Now().String())

	//fmt.Println(json.Unmarshal(jb, pm))
	//fmt.Println(pm)
}

func registerFileDescriptors(fds []*desc.FileDescriptor) (err error) {
	var registry *protoregistry.Files
	fdset := desc.ToFileDescriptorSet(fds...)
	registry, err = protodesc.NewFiles(fdset)
	if err != nil {
		return err
	}
	registry.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if ofd, _ := protoregistry.GlobalFiles.FindFileByPath(fd.Path()); ofd != nil {
			return true
		}

		err = protoregistry.GlobalFiles.RegisterFile(fd)
		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	})
	return
}

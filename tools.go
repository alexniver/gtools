package gtools

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"reflect"
)

// Uint16ToByteArray 将uint16转为[]byte
func Uint16ToByteArray(num uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, num)
	return b
}

// Uint32ToByteArray 将uint32转为[]byte
func Uint32ToByteArray(num uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, num)
	return b
}

// Uint64ToByteArray 将uint64转为[]byte
func Uint64ToByteArray(num uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, num)
	return b
}

// Encode 将struct转为byte
func Encode(e interface{}) ([]byte, error) {
	return json.Marshal(e)
}

// Decode 将byte转为struct
func Decode(b []byte, e interface{}) error {
	return json.Unmarshal(b, e)
}

// EncodeGZip 将struct转为byte
func EncodeGZip(e interface{}) ([]byte, error) {
	b, err := Encode(e)
	if nil != err {
		return b, err
	}

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err = zw.Write(b)
	if err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeUnGzip 解压并将byte转为struct
func DecodeUnGzip(b []byte, e interface{}) error {
	var buf bytes.Buffer
	buf.Write(b)
	zr, err := gzip.NewReader(&buf)
	if err != nil {
		return err
	}

	newB, err := ioutil.ReadAll(zr)
	if nil != err {
		return err
	}

	err = Decode(newB, &e)
	if nil != err {
		return err
	}

	return nil
}

// TypeName get name of Type t
func TypeName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Array,
		reflect.Slice,
		reflect.Map,
		reflect.Chan,
		reflect.Ptr:
		return t.Elem().Name()
	default:
		return t.Name()
	}
}

// 本机ip
func LocalhostIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

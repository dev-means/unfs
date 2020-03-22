package unfs

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type unFs struct {
	client *minio.Client
	Bucket string
	Server string
}

func NewUnFS(server, id, key string) unFs {
	client, err := minio.New(server, id, key, false)
	if err != nil {
		panic(err)
	}
	return unFs{
		client: client,
	}
}

func (fs *unFs) PutObject(fid string, data []byte, cType string) (e error) {
	var exists bool
	exists, e = fs.client.BucketExists(fs.Bucket)
	if e != nil || !exists {
		if e = fs.client.MakeBucket(fs.Bucket, "ASIA"); e != nil {
			return
		}
	}
	_, e = fs.client.PutObject(fs.Bucket, fid, bytes.NewReader(data), int64(len(data)),
		minio.PutObjectOptions{ContentType: cType})
	return
}

func (fs *unFs) Download(fid, host string) (string, error) {
	dl, e := fs.client.PresignedGetObject(fs.Bucket, fid, time.Minute, nil)
	if e != nil {
		return "", e
	} else {
		dl.Host = host
	}
	return dl.String(), nil
}

func SaveMinIO(fs unFs, data *[]byte) (res string) {
	cType := http.DetectContentType(*data)
	spl := strings.Split(cType, "/")
	if len(spl) != 2 {
		return
	}
	filename := fmt.Sprintf("%s.%s", uuId(), spl[1])
	bucket := ""
	if strings.HasPrefix(cType, "image/") {
		bucket = "images"
	} else {
		bucket = spl[1]
	}
	fs.Bucket = bucket
	if e := fs.PutObject(filename, *data, cType); e != nil {
		panic(e)
	}
	res = fmt.Sprintf("%s/%s/%s",
		fs.Server, bucket, filename)
	return
}

var seed = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

func uuId() string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var s []byte
	for {
		i := r.Intn(len(seed))
		s = append(s, seed[i])
		if len(s) == 2 {
			return fmt.Sprintf("%s.%d", s, uuid.New().ID())
		}
	}
}

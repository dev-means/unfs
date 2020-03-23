package main

//const (
//	min_server    = "127.0.0.1:9000"
//	min_accessID  = "minio"
//	min_accessKey = "minioAdmin"
//)
//
//var (
//	unFs = unfs.New(min_server, min_accessID, min_accessKey)
//)
//
//func main() {}

//
//func init() {
//	unFs.Server = "http://test:9000"
//}
//
//func main() {
//	Q := quark.NewHTTP(":9001")
//	Q.Add("PUT", "/unfs/put", objectPut)
//	quark.ShowAPIs(true)
//	if e := Q.Server.ListenAndServe(); e != nil {
//		println(e.Error())
//	}
//}
//
//func objectPut(ct quark.Context) error {
//	var file multipart.File
//	var filename string
//	var data []byte
//	var e error
//	file, filename, e = ct.GetRequestFile("file", 322)
//	if e != nil {
//		panic(e)
//	}
//	data, e = ioutil.ReadAll(file)
//	if e != nil {
//		panic(e)
//	}
//
//	src := unfs.SaveMinIO(unFs, &data)
//	_ = filename
//
//	return ct.ToJson(200, "ok", map[string]interface{}{
//		"src": src,
//	})
//}

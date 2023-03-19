package bcapi

import (
	"archive/zip"
	"flag"
	"net/http"
	"sync/atomic"
	"testing"
)

var (
	needDownloadZip = flag.Bool("download_zip",
		true,
		"if true then will download zip file making http request")
	downloaded = atomic.Bool{}
)

//func Test_getZipFile(t *testing.T) {
//	type args struct {
//		client *http.Client
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		{
//			name: "OK",
//			args: args{
//				client: http.DefaultClient,
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := downloadZipArchive(tt.args.client); (err != nil) != tt.wantErr {
//				t.Errorf("getZipFile() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func Test_openFile(t *testing.T) {
	flag.Parse()
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				fileName: ratesFileName,
			},
			wantErr: false,
		},
	}

	if *needDownloadZip && !downloaded.Load() {
		downloaded.Store(true)
		t.Log("starting download zip archive... 🕔🕔🕔")
		if err := downloadZipArchive(http.DefaultClient); err != nil {
			t.Fatalf(err.Error())
		}
	}

	zipOpened, err := zip.OpenReader(zipFileName)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer zipOpened.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := openFile(zipOpened, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

package OSS

import (
	"MsrmTools/common"
	"bytes"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"golang.org/x/net/context"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

type OssUtil struct {
	QnBucket    string
	QnSecretKey string
	QnAccessKey string
}

/*
分片上传文件
*/
func (this *OssUtil) UpLoadFileInputStream(ctx context.Context, file *multipart.FileHeader, dir string) string {
	// currentTime := common.DateToString(time.Now(), common.YYYYMMddHHmmss)
	uuid, _ := uuid.NewUUID()
	nameKey := strings.ReplaceAll(uuid.String(), "-", "")
	ossKey := filepath.Join(dir, nameKey)
	putPolicy := storage.PutPolicy{Scope: this.QnBucket}
	mac := qbox.NewMac(this.QnAccessKey, this.QnSecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	open, _ := file.Open()
	ByteData := make([]byte, file.Size)
	_, _ = open.Read(ByteData)
	open.Close()
	err := formUploader.Put(ctx, &ret, upToken, ossKey, bytes.NewReader(ByteData), file.Size, &putExtra)
	if err != nil {
		return ""
	}
	return nameKey
}

/*
分片上传文件
*/
func (this *OssUtil) UpLoadFileBlock(ctx context.Context, localFilePath, dir string) string {
	currentTime := common.DateToString(time.Now(), common.YYYYMMddHHmmss)
	uuid, _ := uuid.NewUUID()
	nameKey := strings.ReplaceAll(uuid.String(), "-", "")

	dir += "/" + currentTime + nameKey
	putPolicy := storage.PutPolicy{Scope: this.QnBucket}
	mac := qbox.NewMac(this.QnAccessKey, this.QnSecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputV2Extra{}
	err := resumeUploader.PutFile(ctx, &ret, upToken, nameKey, localFilePath, &putExtra)
	if err != nil {
		return ""
	}
	return nameKey
}

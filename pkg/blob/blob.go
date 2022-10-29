package blob

import (
	"context"
	"fmt"

	"net/url"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/google/uuid"
)

type Blob struct {
	endPoint    string
	container   string
	azrKey      string
	accountName string
}

func New(endPoint, container, azrKey, accountName string) *Blob {
	return &Blob{
		endPoint:    endPoint,
		container:   container,
		azrKey:      azrKey,
		accountName: accountName,
	}
}

func getBlobName() string {
	t := time.Now()
	uuid := uuid.NewString()

	return fmt.Sprintf("%s-%v.jpg", t.Format("20060102"), uuid)
}

func (b Blob) UploadBytesToBlob(data []byte, metadata map[string]string) (string, error) {

	u, _ := url.Parse(fmt.Sprint(b.endPoint, b.container, "/", getBlobName()))

	credential, err := azblob.NewSharedKeyCredential(b.accountName, b.azrKey)
	if err != nil {
		return "", err
	}

	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background()
	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: "image/jpg",
		},
		Metadata: metadata,
	}

	_, err = azblob.UploadBufferToBlockBlob(ctx, data, blockBlobUrl, o)

	return blockBlobUrl.String(), err
}

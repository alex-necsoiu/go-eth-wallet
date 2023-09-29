package s3

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// StoreAccountsIndex stores the account index.
func (s *Store) StoreAccountsIndex(walletID uuid.UUID, data []byte) error {
	var err error

	// Do not encrypt empty index.
	if len(data) != 2 {
		data, err = s.encryptIfRequired(data)
		if err != nil {
			return err
		}
	}

	path := s.walletIndexPath(walletID)
	uploader := s3manager.NewUploader(s.session)
	if _, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
		Body:   bytes.NewReader(data),
	}); err != nil {
		return errors.Wrap(err, "failed to store wallet index")
	}

	return nil
}

// RetrieveAccountsIndex retrieves the account index.
func (s *Store) RetrieveAccountsIndex(walletID uuid.UUID) ([]byte, error) {
	path := s.walletIndexPath(walletID)
	buf := aws.NewWriteAtBuffer([]byte{})
	downloader := s3manager.NewDownloader(s.session)
	if _, err := downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucket),
			Key:    aws.String(path),
		}); err != nil {
		return nil, err
	}
	data := buf.Bytes()
	// Do not decrypt empty index.
	if len(data) == 2 {
		return data, nil
	}
	var err error
	if data, err = s.decryptIfRequired(data); err != nil {
		return nil, err
	}

	return data, nil
}

package utils

import (
	"encoding/base64"
	"io"
	"mime/multipart"
)

func (c *Converter) ImageToByte(file multipart.FileHeader) ([]byte, error) {
	// Open file
	content, openErr := file.Open()
	if openErr != nil {
		return nil, openErr
	}
	defer content.Close()

	// Cast file to byte
	contentByte, byteErr := io.ReadAll(content)
	if byteErr != nil {
		return nil, byteErr
	}
	return contentByte, nil
}

func ByteToBase64(contentByte []byte) string {
	// Cast byte to base64
	content64 := base64.StdEncoding.EncodeToString(contentByte)
	return content64
}

func (c *Converter) ImageToBase64(file multipart.FileHeader) (*string, error) {
	contentByte, err := c.ImageToByte(file)
	if err != nil {
		return nil, err
	}
	content64 := ByteToBase64(contentByte)
	return &content64, nil
}

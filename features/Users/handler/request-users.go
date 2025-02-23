package handler

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type UserRequest struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	FullName       string `json:"fullname" form:"fullname"`
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	PhoneNumber    string `json:"phone_number" form:"phone_number"`
	Address        string `json:"address" form:"address"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (u UserRequest) uploadToCloudinary(file io.Reader, filename string) (string, error) {
	// Konfigurasi Cloudinary
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	// Upload file ke Cloudinary
	uploadParams := uploader.UploadParams{
		Folder:   "user_pictures",
		PublicID: filename,
	}
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}

	// Ambil URL publik dari hasil unggah
	publicURL := uploadResult.URL
	return publicURL, nil
}

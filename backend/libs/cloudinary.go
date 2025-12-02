package libs

import (
	"bytes"
	"context"
	"flower-backend/config"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.uber.org/zap"
)

// NewCloudinary creates a configured Cloudinary client using application config.
func NewCloudinary(cfg *config.Config) (*cloudinary.Cloudinary, error) {
	if cfg == nil {
		return nil, fmt.Errorf("cloudinary config: cfg is nil")
	}

	cld, err := cloudinary.NewFromParams(
		cfg.CloudinaryCloudName,
		cfg.CloudinaryAPIKey,
		cfg.CloudinaryAPISecret,
	)
	if err != nil {
		return nil, fmt.Errorf("cloudinary config: %w", err)
	}

	cld.Config.URL.Secure = cfg.GO_ENV == "production"

	return cld, nil
}

// UploadToCloudinary uploads an image buffer to Cloudinary with optional public ID.
//
// Parameters:
//   - cld: The Cloudinary client instance
//   - buffer: The image data as a byte slice
//   - publicId: Optional public ID for the uploaded image (empty string if not provided)
//
// Returns:
//   - *uploader.UploadResult: The upload result containing the image URL and metadata
//   - error: Any error that occurred during upload
func UploadToCloudinary(cld *cloudinary.Cloudinary, buffer []byte, publicId string) (*uploader.UploadResult, error) {
	logger := zap.L()
	ctx := context.Background()

	uploadParams := uploader.UploadParams{
		AllowedFormats: []string{"png", "jpg", "webp"},
		ResourceType:   "image",
		Folder:         "flower-sharing",
		Transformation: "q_auto",
	}

	// Set public ID only if provided
	if publicId != "" {
		uploadParams.PublicID = publicId
	}

	// Upload the image buffer
	uploadResult, err := cld.Upload.Upload(ctx, bytes.NewReader(buffer), uploadParams)
	if err != nil {
		logger.Error("Error uploading image to Cloudinary", zap.Error(err))
		return nil, fmt.Errorf("error uploading image to Cloudinary: %w", err)
	}

	return uploadResult, nil
}

func DeleteFromCloudinary(cld *cloudinary.Cloudinary, publicId string) error {
	logger := zap.L()
	ctx := context.Background()
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicId,
	})
	if err != nil {
		logger.Error("Error deleting image from Cloudinary", zap.Error(err))
		return fmt.Errorf("error deleting image from Cloudinary: %w", err)
	}
	return err
}

func ExtractPublicId(imageURL string) string {
	parts := strings.Split(imageURL, "/")
	last := parts[len(parts)-1]
	publicId := strings.TrimSuffix(last, filepath.Ext(last))
	return "flower-sharing/" + publicId
}

//go:generate mockery --dir=. --output=../../infrastructure/service/mock --name=ZipCodeService --structname=MockZipCodeService --outpkg=mock --filename=zipcode_service.go --disable-version-string
package gateway

import (
	"climate/internal/business/model"
	"context"
)

type ClimateService interface {
	Get(ctx context.Context, zipCodeID model.ZipCodeID) (*model.Climate, error)
}

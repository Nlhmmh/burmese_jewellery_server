//go:build generate

package generate

import (
	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=models.config.yml ../swagger/swagger.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.config.yml ../swagger/swagger.yml
////go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=client.config.yml ../swagger/swagger.yml

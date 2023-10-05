package common

import (
	"github.com/evgenivanovi/gomart/internal/domain/core"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/pg"
)

/* __________________________________________________ */

func TranslateReadError(err error) error {

	if err == nil {
		return nil
	}

	errCode := pg.ErrorCode(err)
	errEntity := pg.ErrorEntity(err)

	if errCode == pg.ErrorEmptyCode {
		return errx.NewErrorWithEntityCode(
			errEntity, core.ErrorNotFoundCode,
		)
	}

	return errx.NewErrorWithEntityCodeMessage(
		errEntity, errx.ErrorInternalCode, err.Error(),
	)

}

func TranslateWriteError(err error) error {

	if err == nil {
		return nil
	}

	errCode := pg.ErrorCode(err)
	errEntity := pg.ErrorEntity(err)

	if errCode == pg.ErrorUniqueCode {
		return errx.NewErrorWithEntityCode(
			errEntity, core.ErrorExistsCode,
		)
	}

	return errx.NewErrorWithEntityCodeMessage(
		errEntity, errx.ErrorInternalCode, err.Error(),
	)

}

/* __________________________________________________ */

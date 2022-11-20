package usecase

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mayukorin/ebook-merge/app/domain/model"
	"github.com/mayukorin/ebook-merge/app/domain/repository"
	"github.com/mayukorin/ebook-merge/app/domain/service"
	"golang.org/x/oauth2"
)

type EbookUseCase struct {
	db   *sqlx.DB
	conf *oauth2.Config
}

func NewEbookUseCase(db *sqlx.DB, conf *oauth2.Config) *EbookUseCase {
	return &EbookUseCase{
		db:   db,
		conf: conf,
	}
}

func (e *EbookUseCase) Index(userId int64) ([]model.Ebook, error) {
	return repository.AllEbooks(e.db, userId)
}

func (e *EbookUseCase) ScanAllEbooksFromGmail(userId int64) error {
	fmt.Println("ここからusecase")
	gmailAPiOauth2Tokens, err := repository.SelectGmailApiOAuth2TokenByUserID(e.db, userId)
	if err != nil {
		return fmt.Errorf("failed gmail_api_oauth2_tokens by user id:%w", err)
	}

	kindleEbookService, err := repository.FindEbookServiceByName(e.db, "Kindle")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed kindle ebook service:%w", err)
	}

	bookLiveEbookService, err := repository.FindEbookServiceByName(e.db, "BookLive")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed BookLive ebook service:%w", err)
	}
	fmt.Println("ここまでOK")
	for _, gmailOauth2Token := range gmailAPiOauth2Tokens {
		// TODO：allservice で一つずつfindOrEbookやfindOrCreatePurchaseを実行
		scannedBookLiveEbookTitles, err := service.ScanBookLiveEbooks(gmailOauth2Token, e.conf)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("failed scan ebook of booklive:%w", err)
		}
		fmt.Println("ここまでOK2")
		for _, bookLiveTitle := range scannedBookLiveEbookTitles {
			eBookId, err := service.FindOrCreateEbook(bookLiveTitle, bookLiveEbookService.ID, e.db)
			if err != nil {
				fmt.Println(err)
				return fmt.Errorf("failed find or create ebook of book live:%w", err)
			}
			_, err = service.FindOrCreatePurchase(userId, eBookId, e.db)
			if err != nil {
				fmt.Println(err)
				return fmt.Errorf("failed find or create purchase:%w", err)
			}
		}
		fmt.Println("ここまでOK3")
		scannedKindleEbookTitles, err := service.ScanKindleEbooks(gmailOauth2Token, e.conf)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("failed scan ebook of kindle:%w", err)
		}
		fmt.Println("ここまでOK4")
		for _, kindleEbookTitle := range scannedKindleEbookTitles {
			eBookId, err := service.FindOrCreateEbook(kindleEbookTitle, kindleEbookService.ID, e.db)
			if err != nil {
				fmt.Println(err)
				return fmt.Errorf("failed find or create ebook of kindle:%w", err)
			}
			_, err = service.FindOrCreatePurchase(userId, eBookId, e.db)
			if err != nil {
				fmt.Println(err)
				return fmt.Errorf("failed find or create purchase:%w", err)
			}
		}

	}
	return nil
}

func (e *EbookUseCase) TestScanKindleEbooksFromGmail(userId int64) error {
	gmailAPiOauth2Tokens, err := repository.SelectGmailApiOAuth2TokenByUserID(e.db, userId)
	if err != nil {
		return fmt.Errorf("failed gmail_api_oauth2_tokens by user id:%w", err)
	}

	if err != nil {
		return fmt.Errorf("failed BookLive ebook service:%w", err)
	}

	for _, gmailOauth2Token := range gmailAPiOauth2Tokens {
		kidleEbookTitles, err := service.TestScanBookLiveEbooks(gmailOauth2Token, e.conf)
		if err != nil {
			return fmt.Errorf("failed scan ebook of kidle:%w", err)
		}
		for _, title := range kidleEbookTitles {
			fmt.Println(title)
		}
		// fmt.Println((kidleEbookTitles))

	}
	return nil
}

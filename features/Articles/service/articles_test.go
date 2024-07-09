package service_test

import (
	articles "NEWSAPP/features/Articles"
	"NEWSAPP/features/Articles/service"
	"NEWSAPP/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateArticles(t *testing.T) {
	mockArticleData := new(mocks.DataArtikelInterface)
	articleService := service.New(mockArticleData)

	t.Run("success", func(t *testing.T) {
		article := articles.Artikel{
			ArtikelPicture: "articles.png",
			ArtikelName:    "Learn Golang",
			Tag:            "programming",
			Description:    "Learn Golang basic functionality",
		}

		mockArticleData.On("Insert", article).Return(nil).Once()

		err := articleService.Create(article)
		assert.NoError(t, err)
		mockArticleData.AssertExpectations(t)
	})

	t.Run("failed - empty artikel name", func(t *testing.T) {
		article := articles.Artikel{
			ArtikelName: "",
		}

		err := articleService.Create(article)
		assert.Error(t, err)
		assert.Equal(t, "artikel name cannot be empty", err.Error())
	})

	t.Run("failed - data insert error", func(t *testing.T) {
		article := articles.Artikel{
			ArtikelPicture: "articles.png",
			ArtikelName:    "Learn Golang",
			Tag:            "programming",
			Description:    "Learn Golang basic functionality",
		}

		mockArticleData.On("Insert", article).Return(errors.New("insert error")).Once()

		err := articleService.Create(article)
		assert.Error(t, err)
		assert.Equal(t, "insert error", err.Error())
		mockArticleData.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockArticleData := new(mocks.DataArtikelInterface)
	articleService := service.New(mockArticleData)

	t.Run("success", func(t *testing.T) {
		expectedArticles := []articles.Artikel{
			{
				UserID:         1,
				ArtikelPicture: "articles.png",
				ArtikelName:    "Article 1",
				Tag:            "politics",
				Description:    "Content of article 1",
			},
			{
				UserID:         2,
				ArtikelPicture: "articles.png",
				ArtikelName:    "Article 2",
				Tag:            "politics",
				Description:    "Content of article 2",
			},
		}

		mockArticleData.On("GetAll").Return(expectedArticles, nil).Once()

		returnedArticles, err := articleService.GetAllArtikel()

		assert.NoError(t, err)
		assert.Equal(t, expectedArticles, returnedArticles)
		mockArticleData.AssertExpectations(t)
	})

	t.Run("error - get all error", func(t *testing.T) {
		mockArticleData.On("GetAll").Return(nil, errors.New("get all error")).Once()

		returnedArticles, err := articleService.GetAllArtikel()

		assert.Error(t, err)
		assert.Equal(t, "get all error", err.Error())
		assert.Nil(t, returnedArticles)
		mockArticleData.AssertExpectations(t)
	})
}

func TestUpdateArticle(t *testing.T) {
	mockArticleData := new(mocks.DataArtikelInterface)
	articleService := service.New(mockArticleData)

	t.Run("success", func(t *testing.T) {
		articleID := uint(1)
		userID := uint(1)
		article := articles.Artikel{
			ArtikelPicture: "updated.png",
			ArtikelName:    "Updated Article",
			Tag:            "technology",
			Description:    "This is the updated content of the article",
		}

		mockArticleData.On("SelectById", articleID).Return(&articles.Artikel{UserID: userID}, nil).Once()
		mockArticleData.On("Update", articleID, article).Return(nil).Once()

		err := articleService.Update(articleID, userID, article)

		assert.NoError(t, err)
		mockArticleData.AssertExpectations(t)
	})

	t.Run("error - empty artikel name", func(t *testing.T) {
		articleID := uint(1)
		userID := uint(1)
		article := articles.Artikel{
			ArtikelPicture: "updated.png",
			ArtikelName:    "",
			Tag:            "technology",
			Description:    "This is the updated content of the article",
		}

		err := articleService.Update(articleID, userID, article)

		assert.Error(t, err)
		assert.Equal(t, "artikel name cannot be empty", err.Error())
		mockArticleData.AssertExpectations(t)
	})

	t.Run("error - user id not match", func(t *testing.T) {
		articleID := uint(1)
		userID := uint(2) // different userID
		article := articles.Artikel{
			ArtikelPicture: "updated.png",
			ArtikelName:    "Updated Article",
			Tag:            "technology",
			Description:    "This is the updated content of the article",
		}

		mockArticleData.On("SelectById", articleID).Return(&articles.Artikel{UserID: 1}, nil).Once()

		err := articleService.Update(articleID, userID, article)

		assert.Error(t, err)
		assert.Equal(t, "user id not match, cannot update artikel", err.Error())
		mockArticleData.AssertExpectations(t)
	})

	t.Run("error - invalid artikel ID", func(t *testing.T) {
		articleID := uint(0) // invalid ID
		userID := uint(1)
		article := articles.Artikel{
			ArtikelPicture: "updated.png",
			ArtikelName:    "Updated Article",
			Tag:            "technology",
			Description:    "This is the updated content of the article",
		}

		err := articleService.Update(articleID, userID, article)

		assert.Error(t, err)
		assert.Equal(t, "invalid artikel ID", err.Error())
		mockArticleData.AssertExpectations(t)
	})
}

func TestDeleteArticle(t *testing.T) {
	mockArticleData := new(mocks.DataArtikelInterface)
	articleService := service.New(mockArticleData)

	t.Run("success", func(t *testing.T) {
		articleID := uint(1)
		userID := uint(1)

		mockArticleData.On("SelectById", articleID).Return(&articles.Artikel{UserID: userID}, nil).Once()
		mockArticleData.On("Delete", articleID).Return(nil).Once()

		err := articleService.Delete(articleID, userID)

		assert.NoError(t, err)
		mockArticleData.AssertExpectations(t)
	})

	t.Run("error - delete error", func(t *testing.T) {
		articleID := uint(1)
		userID := uint(1)

		mockArticleData.On("SelectById", articleID).Return(&articles.Artikel{UserID: userID}, nil).Once()
		mockArticleData.On("Delete", articleID).Return(errors.New("delete error")).Once()

		err := articleService.Delete(articleID, userID)

		assert.Error(t, err)
		assert.Equal(t, "delete error", err.Error())
		mockArticleData.AssertExpectations(t)
	})

	t.Run("error - user id not match", func(t *testing.T) {
		articleID := uint(1)
		userID := uint(2) // different userID

		mockArticleData.On("SelectById", articleID).Return(&articles.Artikel{UserID: 1}, nil).Once()

		err := articleService.Delete(articleID, userID)

		assert.Error(t, err)
		assert.Equal(t, "user id not match, cannot delete artikel", err.Error())
		mockArticleData.AssertExpectations(t)
	})

	t.Run("error - invalid artikel ID", func(t *testing.T) {
		articleID := uint(0) // invalid ID
		userID := uint(1)

		err := articleService.Delete(articleID, userID)

		assert.Error(t, err)
		assert.Equal(t, "invalid artikel ID", err.Error())
		mockArticleData.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	mockArticleData := new(mocks.DataArtikelInterface)
	articleService := service.New(mockArticleData)

	t.Run("success", func(t *testing.T) {
		articleID := uint(1)
		expectedArticle := &articles.Artikel{
			UserID:         1,
			ArtikelPicture: "articles.png",
			ArtikelName:    "Article 1",
			Tag:            "politics",
			Description:    "Content of article 1",
		}

		mockArticleData.On("SelectById", articleID).Return(expectedArticle, nil).Once()

		returnedArticle, err := articleService.GetById(articleID)

		assert.NoError(t, err)
		assert.Equal(t, expectedArticle, returnedArticle)
		mockArticleData.AssertExpectations(t)
	})

	t.Run("invalid ID", func(t *testing.T) {
		articleID := uint(0)

		returnedArticle, err := articleService.GetById(articleID)

		assert.Error(t, err)
		assert.Equal(t, "id not valid", err.Error())
		assert.Nil(t, returnedArticle)
	})

	t.Run("not found", func(t *testing.T) {
		articleID := uint(1)

		mockArticleData.On("SelectById", articleID).Return(nil, errors.New("not found")).Once()

		returnedArticle, err := articleService.GetById(articleID)

		assert.Error(t, err)
		assert.Equal(t, "not found", err.Error())
		assert.Nil(t, returnedArticle)
		mockArticleData.AssertExpectations(t)
	})
}

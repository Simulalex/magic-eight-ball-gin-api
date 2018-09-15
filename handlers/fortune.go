package handlers

import (
	"github.com/Simulalex/magic-eight-ball-gin-api/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Fortune struct {
	Text string `json:"text"`
}

type FortuneHandler interface {
	Read(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type FortuneHandlerImpl struct {
	db db.FortuneDatabase
}

func Create(dbFilePath string) FortuneHandler {
	return FortuneHandlerImpl{db.Create(dbFilePath)}
}

func (impl FortuneHandlerImpl) Read(c *gin.Context) {
	text, err := impl.db.ReadRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Fortune{text})
}

func (impl FortuneHandlerImpl) Create(c *gin.Context) {
	var fortune Fortune
	if err := c.BindJSON(&fortune); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := impl.db.Create(fortune.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (impl FortuneHandlerImpl) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var fortune Fortune
	if err := c.BindJSON(&fortune); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := impl.db.Update(id, fortune.Text); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (impl FortuneHandlerImpl) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := impl.db.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritimauludin/go-restapi-gin-gorm-mysql/models"
	"gorm.io/gorm"
)

func Index(ctx *gin.Context)  {
	var products []models.Product

	models.DB.Find(&products)
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}
func Show(ctx *gin.Context)  {
	var product models.Product
	id := ctx.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message" : err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"product" : product})
}
func Create(ctx *gin.Context)  {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
			return
	}
	
	models.DB.Create(&product)
	ctx.JSON(http.StatusOK, gin.H{"product" : product})
}
func Update(ctx *gin.Context)  {
	var product models.Product
	id := ctx.Param("id")

	if err :=ctx.ShouldBindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ? ", id).Updates(&product).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot update data, id not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message" : "Product successfuly updated"})
}

func Delete(ctx *gin.Context)  {
	var product models.Product

	// input := map[string]string{"id": "0"} //error if json send integer
	var input struct {
		Id json.Number
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// id, _ := strconv.ParseInt(input["id"], 10, 64) //error if json send integer
	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot delete data, id not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message" : "Product successfuly deleted"})
}

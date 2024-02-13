package user

import "github.com/gin-gonic/gin"

var db *Database

func init() {
	db = NewDb()
}

func GetAllUsers(ctx *gin.Context) {
	ctx.JSON(200, db.GetAll())
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := db.Get(id)
	ctx.JSON(200, user)
}

func CreateUser(ctx *gin.Context) {
	user := &User{}
	ctx.Bind(user)
	result := User{}
	result.ID = "1" + user.Name
	result.Name = user.Name
	result.Email = user.Email
	db.Create(result)
	ctx.JSON(200, &result)
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	result := User{}
	ctx.Bind(&result)
	result.ID = id
	db.Update(id, result)
	ctx.JSON(200, &result)
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	db.Delete(id)
	ctx.JSON(200, gin.H{"message": "delete", "id": id})
}

package controllers

import (
    "github.com/gin-gonic/gin"
    "microservice-uporabniki/initializers"
    "microservice-uporabniki/models"
    "net/http"
)

type UserBody struct{
    Name string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

// AddUser godoc
// @Summary      Add an user
// @Description  add by json user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      UserBody  true  "Add user"
// @Success      200      {object}  models.User
// @Router       /users [post]
func UsersCreate(c *gin.Context) {
    // Parse request
    var body UserBody;

    err := c.ShouldBindJSON(&body)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,
            gin.H{
                "error": "validation_err",
                "message": err.Error(),
            })
        return
    }

    // Create entity
    user := models.User{Name: body.Name, Email: body.Email}

    result := initializers.DB.Create(&user)

    if result.Error != nil{
        c.Status(http.StatusInternalServerError)
        return
    }


    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}

// ListUsers godoc
// @Summary      List users
// @Description  get users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.User
// @Router       /users [get]
func UsersIndex(c *gin.Context){
    var users[] models.User;
    initializers.DB.Find(&users)

    c.JSON(http.StatusOK, gin.H{
        "users": users,
    })
}

// UsersShow ShowUsers godoc
// @Summary      Show user
// @Description  get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Router       /users/{id} [get]
func UsersShow(c *gin.Context){
    // Get ID from URL
    id := c.Param("id")

    var user models.User
    initializers.DB.First(&user, id)

    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}

// UpdateUser godoc
// @Summary      Update an user
// @Description  Update by json user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "User ID"
// @Param        account  body      UserBody true  "Update user"
// @Success      200      {object}  models.User
// @Router       /users/{id} [put]
func UsersUpdate(c *gin.Context){
    id := c.Param("id")

    var body UserBody;

    err := c.ShouldBindJSON(&body)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusBadRequest,
            gin.H{
                "error": "validation_err",
                "message": err.Error(),
            })
        return
    }

    var user models.User;
    initializers.DB.Find(&user, id).Updates(models.User{Name: body.Name, Email: body.Email})

    c.JSON(http.StatusOK, gin.H{
        "user": user,
    })
}

// DeleteUser godoc
// @Summary      Delete an user
// @Description  Delete by user ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"  Format(int64)
// @Router       /users/{id} [delete]
func UsersDelete(c *gin.Context){
    id := c.Param("id")


    initializers.DB.Delete(&models.User{}, id)

    c.JSON(http.StatusOK, gin.H{
    })
}

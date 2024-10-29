package api

import (
	"fmt"
	"local/fin/controllers"
	_ "local/fin/docs"
	"local/fin/forms"
	"local/fin/models"
	"local/fin/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
)

// @Title crud APIs
// @Version 0.0.2
// @Description For local container 'maria'(mariaDB) database CRUD
// @host 127.0.0.1:11011
// @BasePath /

func PostFile(c echo.Context, key string) (users []models.UserModel, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.rest.PostFile: %v", r)
		}
	}()
	file, err := c.FormFile(key)
	if err != nil {
		c.Set("error", err)
		return nil, err
	}
	src, err := file.Open()
	if err != nil {
		c.Set("error", err)
		return nil, err
	}
	defer src.Close()
	if err := gocsv.UnmarshalMultipartFile(&src, &users); err != nil {
		c.Set("error", err)
		return nil, err
	}
	return users, nil
}

// @Summary        	Get one record
// @Description    	Search record using primary key(id)
// @Tags           	CRUD
// @Accept         	json
// @Produce        	json
// @Param          	id query uint64 true "user Id"
// @Success        	200 {object} ReturnUserModel "success processing your request"
// @Failure         400 {object} ReturnError "invalid data(Query)"
// @Router         	/get [get]
func GetUser(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.rest.GetUser: %v", r)
			utils.ErrLogging()
		}
	}()
	id := c.QueryParam("id")
	conv_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Set("error", err)
		return c.JSONPretty(http.StatusBadRequest, InvalidParams("id"), "  ")
	}

	user, err := controllers.Get(conv_id)
	if err != nil {
		c.Set("error", err)
		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
	}
	return c.JSONPretty(http.StatusOK, ReturnUserModel{User: user}, "  ")
}

// @Summary        	Get all record
// @Description    	Print all record in 'user_models' table
// @Tags           	CRUD
// @Accept         	json
// @Produce        	json
// @Success        	200 {object} []models.UserModel "success processing your request"
// @Failure         500 {object} ReturnError "errors in database"
// @Router         	/list [get]
func ListUser(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.rest.ListUser: %v", r)
			utils.ErrLogging()
		}
	}()
	users, err := controllers.List()
	if err != nil {
		c.Set("error", err)
		return c.JSONPretty(http.StatusInternalServerError, ReturnError{ErrorMessage: err.Error()}, "  ")
	}
	return c.JSONPretty(http.StatusOK, users, "  ")
}

// @Summary        	Create new record
// @Description    	Create new record except primary key(id)
// @Tags           	CRUD
// @Accept         	json
// @Produce        	json
// @Success        	201 {object} ReturnMessage "success processing your request"
// @Failure         400 {object} ReturnError "invalid data(Body)"
// @Router         	/create [post]
func CreateUser(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.rest.CreateUser: %v", r)
			utils.ErrLogging()
		}
	}()

	var user = new(forms.UserForm)
	if err := c.Bind(user); err != nil {
		c.Set("error", err)
		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
	}

	if err := controllers.Create(user); err != nil {
		c.Set("error", err)
		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
	}
	return c.JSONPretty(http.StatusCreated, ReturnMessage{Message: fmt.Sprintf("User '%v' successfuly created", user.Email)}, "  ")
}

// @Summary        	Update exist record
// @Description    	Replace existing record(id) with new record(body)
// @Tags           	CRUD
// @Accept         	json
// @Produce        	json
// @Param          	id query uint64 true "user Id"
// @Param						body_data body string true "will be replaced data(json)"
// @Success        	201 {object} ReturnMessage "success processing your request"
// @Failure         400 {object} ReturnError "invalid data(Query or Body)"
// @Failure         500 {object} ReturnError "errors in database"
// @Router         	/updadte [put]
// func UpdateUser(c echo.Context) error {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			log.Printf("api.rest.UpdateUser: %v", r)
// 			utils.ErrLogging()
// 		}
// 	}()

// 	id := c.QueryParam("id")
// 	conv_id, err := strconv.ParseUint(id, 10, 64)
// 	if err != nil {
// 		c.Set("error", err)
// 		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
// 	}
// 	user := new(models.UserModel)
// 	if err := c.Bind(user); err != nil {
// 		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
// 	}

// 	if err := controllers.Update(conv_id, user); err != nil {
// 		return c.JSONPretty(http.StatusInternalServerError, ReturnError{ErrorMessage: err.Error()}, "  ")
// 	}
// 	return c.JSONPretty(http.StatusCreated, ReturnMessage{Message: fmt.Sprintf("User (id = %v) successfuly updated", id)}, "  ")
// }

// @Summary        	Delete exist record
// @Description    	Delete record using primary key(id)
// @Tags           	CRUD
// @Accept         	json
// @Produce        	json
// @Param          	id query uint64 true "user Id"
// @Success        	200 {object} ReturnMessage "success processing your request"
// @Failure         400 {object} ReturnError "invalid data(Query)"
// @Router         	/delete [delete]
func DeleteUser(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.rest.DeleteUser: %v", r)
			utils.ErrLogging()
		}
	}()

	id := c.QueryParam("id")
	conv_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
	}

	if err := controllers.Delete(conv_id); err != nil {
		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
	}
	return c.JSONPretty(http.StatusOK, ReturnMessage{Message: fmt.Sprintf("User (id = %v) has deleted", id)}, "  ")
}

// @Summary        	Batch data save
// @Description    	Save(Create or Update) records using csv file
// @Tags           	CRUD
// @Accept         	json
// @Produce        	json
// @Param          	file formData file true "user records"
// @Success        	200 {object} []models.UserModel "success processing your request"
// @Failure         400 {object} ReturnError "invalid data(Query)"
// @Router         	/batch-save [post]
func BatchSave(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.rest.BatchSave: %v", r)
			utils.ErrLogging()
		}
	}()
	users, err := PostFile(c, "file")
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
	}
	if err := controllers.BSave(users); err != nil {
		return c.JSONPretty(http.StatusInternalServerError, ReturnError{ErrorMessage: err.Error()}, "  ")
	}
	return c.JSONPretty(http.StatusCreated, users, "  ")
}

// @Summary        	Batch data delete
// @Description    	Delete records using csv file
// @Tags           	CRUD
// @Accept         	json
// @Produce        	json
// @Param          	file formData file true "user records"
// @Success        	200 {object} []models.UserModel "success processing your request"
// @Failure         400 {object} ReturnError "invalid data(Query)"
// @Router         	/batch-delete [delete]
func BatchDelete(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.rest.BatchDelete: %v", r)
			utils.ErrLogging()
		}
	}()
	users, err := PostFile(c, "file")
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
	}

	if err := controllers.BDelete(users); err != nil {
		return c.JSONPretty(http.StatusBadRequest, ReturnError{ErrorMessage: err.Error()}, "  ")
	}
	return c.JSONPretty(http.StatusNoContent, users, "  ")
}

func Reset(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.rest.Reset: %v", r)
			utils.ErrLogging()
		}
	}()
	if err := controllers.Reset(); err != nil {
		return err
	}
	return c.JSONPretty(http.StatusNoContent, ReturnMessage{Message: "reset table"}, "  ")
}
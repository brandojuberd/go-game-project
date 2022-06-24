package users

import (
	"brandos-lair/models"
	"fmt"
	"reflect"
)

type usersService struct {
	model models.Model
	user  User
	users []User
}

func InitService() usersService {
	model := models.Model{}
	usersService := usersService{
		model: model.InitModel(&User{}),
		user:  User{},
		users: []User{},
	}
	return usersService
}

func (usersService usersService) Read() []User {
	fmt.Println(usersService)
	usersService.model.Read(&usersService.users)
	return usersService.users
}

func (usersService usersService) Write(data User) User {
	usersService.model.Read(&usersService.users)
	usersDb := usersService.users
	if data.Id != "" {
		for i := range usersDb {
			userDb := &usersDb[i]
			if userDb.Id == data.Id {
				dataReflect := reflect.ValueOf(&data).Elem()
				userDbReflect := reflect.ValueOf(userDb).Elem()
				for i := 0; i < dataReflect.NumField(); i++ {
					key := dataReflect.Type().Field(i).Name
					if dataReflect.Field(i).IsValid() && userDbReflect.FieldByName(key).CanSet() {
						switch dataReflect.Field(i).Kind() {
						case reflect.Int:
							userDbReflect.FieldByName(key).Set(dataReflect.Field(i))
						case reflect.String:
							userDbReflect.FieldByName(key).Set(dataReflect.Field(i))
						}
					}
				}
			}
		}
		usersService.users = usersDb
	} else {
		usersService.users = append(usersService.users, data)
	}
	usersService.model.Save(usersService.users)
	return data
}



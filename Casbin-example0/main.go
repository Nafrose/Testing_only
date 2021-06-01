package main

import (
	"github.com/Nafrose/Testing_only/Casbin-example0/model"
	"fmt"
	"github.com/casbin/casbin/v2"
	"log"
	"net/http"
)

func main() {
	// setup casbin auth rules
	_, casbinErr := casbin.NewEnforcer("./authorizationModel.conf", "./policy.csv")

	if casbinErr != nil {
		fmt.Println("Error in Casbin enforcer")
		log.Fatalln(casbinErr)
	}

	// setup session store


	// setup Users
	users := createUsers()

	// setup routes
	casMux := http.NewServeMux()
	casMux.HandleFunc("/login", loginHandler(users))
	casMux.HandleFunc("/logout", logoutHandler())
	//casMux.HandleFunc("/member/current", currentMemberHandler())
	//casMux.HandleFunc("/member/role", memberRoleHandler())
	//casMux.HandleFunc("/admin/stuff", adminHandler())

	log.Fatalln(http.ListenAndServe(":9060", casMux))
}

func createUsers() model.Users {
	users := model.Users{}
	users = append(users, model.User{Id: 1, Name: "Admin", Role: "admin"})
	users = append(users, model.User{Id: 2, Name: "Sabine", Role: "member"})
	users = append(users, model.User{Id: 3, Name: "Sepp", Role: "member"})

	return users
}

func loginHandler(users model.Users) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")
		user, err := users.FindByName(name)
		if err != nil {
			log.Fatalln(http.StatusBadRequest, "WRONG_CREDENTIALS", w, err)
			return
		}

		log.Println(http.StatusOK, user)
	})
}

func logoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(http.StatusOK, "logout")
	})
}
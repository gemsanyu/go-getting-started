package container

import (
	"net/http"

	"github.com/heroku/go-getting-started/httphandler/extra"
	"github.com/heroku/go-getting-started/httphandler/ping"
	"github.com/heroku/go-getting-started/httphandler/profile"
	"github.com/heroku/go-getting-started/httphandler/user"
	"github.com/heroku/go-getting-started/role"
	"github.com/gorilla/mux"
	jwt "github.com/mmuflih/go-httplib/httplib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-02-09 20:49
**/
func InvokeRoute(route *mux.Router,
	pingH ping.PingHandler, p404H extra.P404Handler, userAdd user.AddHandler,
	userEdit user.EditHandler, userGet user.GetHandler, userList user.ListHandler,
	getToken user.GetTokenHandler, getMyUser user.GetMyHandler,
	voidUser user.VoidHandler, profileAdd profile.AddHandler, profileEdit profile.EditHandler,
	profileGet profile.GetHandler, profileAvatar profile.UploadAvatarHandler,
) {
	route.NotFoundHandler = http.HandlerFunc(p404H.Handle)
	/** api v1 route */
	apiV1 := route.PathPrefix("/api/v1").Subrouter()
	pingRoute := apiV1.PathPrefix("/ping").Subrouter()
	userRoute := apiV1.PathPrefix("/user").Subrouter()
	myUserRoute := apiV1.PathPrefix("/my/user").Subrouter()
	profileRoute := apiV1.PathPrefix("/profile").Subrouter()

	/** ping */
	pingRoute.HandleFunc("", pingH.Handle).Methods("GET")

	/** user */
	userRoute.HandleFunc("/login", getToken.Handle).Methods("POST")
	userRoute.HandleFunc("", jwt.JWTMidWithRole(userAdd.Handle, role.OWNER)).Methods("POST")
	userRoute.HandleFunc("", jwt.JWTMidWithRole(userList.Handle, role.OWNER)).Methods("GET")
	userRoute.HandleFunc("/{id}", jwt.JWTMidWithRole(userGet.Handle, role.OWNER)).Methods("GET")
	userRoute.HandleFunc("/{id}", jwt.JWTMidWithRole(userEdit.Handle, role.OWNER)).Methods("PUT")
	userRoute.HandleFunc("/{id}", jwt.JWTMidWithRole(voidUser.Handle, role.OWNER)).Methods("DELETE")

	/** my/user **/
	myUserRoute.HandleFunc("", jwt.JWTMidWithRole(getMyUser.Handle, role.USER)).Methods("GET")

	/** profile */
	profileRoute.HandleFunc("", jwt.JWTMidWithRole(profileAdd.Handle, role.OWNER)).Methods("POST")
	profileRoute.HandleFunc("", jwt.JWTMidWithRole(profileEdit.Handle, role.USER)).Methods("PUT")
	profileRoute.HandleFunc("", jwt.JWTMidWithRole(profileGet.Handle, role.USER)).Methods("GET")
	profileRoute.HandleFunc("/avatar", jwt.JWTMidWithRole(profileAvatar.Handle, role.USER)).Methods("PUT")
}

package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
	"user-service/internal/core/port/service"
)

type UserControll struct {
	Mux         *http.ServeMux
	userService service.UserService
}

func NewUserControll(userService service.UserService) UserControll {
	return UserControll{
		Mux:         http.NewServeMux(),
		userService: userService,
	}
}

func (u UserControll) Router() {
	u.Mux.HandleFunc("/", u.Handle)
	u.Mux.HandleFunc("/signup", u.SignUp)
}

func (u UserControll) Handle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/index.html")
}

func (u UserControll) SignUp(w http.ResponseWriter, r *http.Request) {
	request, _ := ParseRequest(r)

	var response *response.Response = u.userService.SignUp(request)

	data, _ := json.Marshal(response)
	w.Write([]byte(data))
}

func ParseRequest(r *http.Request) (*request.SignUpRequest, error) {
	u, err := url.Parse(r.URL.String())

	if err != nil {
		return nil, err
	}

	query := u.Query()

	return &request.SignUpRequest{
		Username: query.Get("username"),
		Password: query.Get("password"),
	}, nil
}

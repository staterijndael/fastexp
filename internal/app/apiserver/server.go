package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Oringik/fastexp/internal/app/model"
	"github.com/Oringik/fastexp/internal/app/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "fastexp"
	ctxKeyUser  ctxKey = iota
)

var (
	errIncorrectEmailOrPassword = errors.New("Incorrect email or password")
	errNotAuthenticated         = errors.New("Not authenticated")
)

type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type", "set-cookie"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)

	s.router.Use(cors)
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST", "OPTIONS")

	private := s.router.PathPrefix("/private").Subrouter()

	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET", "OPTIONS")
	private.HandleFunc("/addtags", s.handleAddTags()).Methods("POST", "OPTIONS")
	private.HandleFunc("/createtheme", s.handleCreateTheme()).Methods("POST", "OPTIONS")
	private.HandleFunc("/generatethemes", s.handleGenerateThemes()).Methods("POST", "OPTIONS")
	private.HandleFunc("/addcard", s.handleAddCard()).Methods("POST", "OPTIONS")
	private.HandleFunc("/deletecard", s.handleDeleteCard()).Methods("POST", "OPTIONS")
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxKeyUser).(*model.User)
		s.setTags(user)
		s.respond(w, r, http.StatusOK, user)
	}

}

func (s *server) handleAddCard() http.HandlerFunc {
	type request struct {
		Name      string
		ShortDesc string
		FullDesc  string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		Card := model.Card{
			Name:      req.Name,
			ShortDesc: req.ShortDesc,
			FullDesc:  req.FullDesc,
		}
		err := s.store.User().CreateCard(&Card)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		s.respond(w, r, http.StatusOK, Card)
	}
}

func (s *server) handleDeleteCard() http.HandlerFunc {
	type request struct {
		Name string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.store.User().DeleteCard(req.Name)

		s.respond(w, r, http.StatusOK, "OK")
	}
}

func (s *server) handleGenerateThemes() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxKeyUser).(*model.User)
		tags, err := s.store.User().GetTags(user.ID)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if len(tags) == 0 {
			s.error(w, r, http.StatusBadRequest, store.TagsNotfound)
			return
		}

		themes, err := s.store.User().GetAllThemes()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var needThemes []model.Theme

		for _, theme := range themes {
			tags2, err := s.store.User().GetThemeTags(theme.ID)
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			for _, tag2 := range tags2 {
				for _, tag := range tags {
					if tag2.Text == tag.Text {
						needThemes = append(needThemes, theme)
					}
				}
			}

		}

		s.respond(w, r, http.StatusOK, needThemes)
	}

}

func (s *server) handleCreateTheme() http.HandlerFunc {
	type request struct {
		Title       string
		Description string
		Tags        []string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		theme := &model.Theme{
			Title:       req.Title,
			Description: req.Description,
		}

		err := s.store.User().CreateTheme(theme)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		err = s.store.User().AddThemeTags(theme.ID, req.Tags)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

	}
}

func (s *server) handleAddTags() http.HandlerFunc {
	type request struct {
		Tags []string `json:"tags"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user := r.Context().Value(ctxKeyUser).(*model.User)

		if err := s.store.User().AddTags(user.ID, req.Tags); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, req.Tags)

	}

}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.logger.Info("User" + req.Email + " successfully created!")

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) setTags(user *model.User) (*model.User, error) {

	tags, err := s.store.User().GetTags(user.ID)

	if err != nil {
		return nil, err
	}

	var tagsText []string

	for _, tag := range tags {
		tagsText = append(tagsText, tag.Text)
	}

	user.Tags = tagsText

	return user, err

}

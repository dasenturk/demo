package web

import (
	"html/template"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"
	"github.com/gorilla/csrf"
	"github.com/dasenturk/demo"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	store    petsy.Store
	sessions *scs.SessionManager
}

func (h *UserHandler) Register() http.HandlerFunc {
	type data struct {
		SessionData
		CSRF template.HTML
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/user_register.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data{
			SessionData: GetSessionData(h.sessions, r.Context()),
			CSRF:        csrf.TemplateField(r),
		})
	}
}

func (h *UserHandler) RegisterSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := RegisterForm{
			Username:      r.FormValue("username"),
			Password:      r.FormValue("password"),
			Name:		   r.FormValue("name"),
			Surname:	   r.FormValue("surname"),
			Email:		   r.FormValue("email"),
			CityName:	   r.FormValue("cityname"),
			EmailTaken: false,
			UsernameTaken: false,
		}
		if _, err := h.store.PetOwnerByUsername(form.Username); err == nil {
			form.UsernameTaken = true
		}
		if _, err := h.store.PetOwnerByEmail(form.Email); err == nil {
			form.EmailTaken = true
		}
		if !form.Validate() {
			h.sessions.Put(r.Context(), "form", form)
			http.Redirect(w, r, r.Referer(), http.StatusFound)
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := h.store.CreatePetOwner(&petsy.PetOwner{
			ID:       uuid.New(),
			Username: form.Username,
			Password: string(password),
			Name: form.Name,
			Surname: form.Surname,
			Email: form.Email,
			CityName: form.CityName,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.sessions.Put(r.Context(), "flash", "Your registration was successful. Please log in.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h *UserHandler) Login() http.HandlerFunc {
	type data struct {
		SessionData
		CSRF template.HTML
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/user_login.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data{
			SessionData: GetSessionData(h.sessions, r.Context()),
			CSRF:        csrf.TemplateField(r),
		})
	}
}

func (h *UserHandler) LoginSubmit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := LoginForm{
			Username:             r.FormValue("username"),
			Password:             r.FormValue("password"),
			IncorrectCredentials: false,
		}
		user, err := h.store.PetOwnerByUsername(form.Username)
		if err != nil {
			form.IncorrectCredentials = true
		} else {
			compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
			form.IncorrectCredentials = compareErr != nil
		}
		if !form.Validate() {
			h.sessions.Put(r.Context(), "form", form)
			http.Redirect(w, r, r.Referer(), http.StatusFound)
			return
		}

		h.sessions.Put(r.Context(), "user_id", user.ID)
		h.sessions.Put(r.Context(), "flash", "You have been logged in sucessfully.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h *UserHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.sessions.Remove(r.Context(), "user_id")
		h.sessions.Put(r.Context(), "flash", "You have been logged out sucessfully.")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

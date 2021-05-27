package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/lpiegas25/go_store/internal/data"
	"github.com/lpiegas25/go_store/pkg/model/account"
	"github.com/lpiegas25/go_store/pkg/model/client"
	"github.com/lpiegas25/go_store/pkg/model/role"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	ar := &AccountRouter{Repository: &account.AccountRepository{Data: data.New()}}
	cr := &ClientRouter{Repository: &client.ClientRepository{Data: data.New()}}
	rr := &RoleRouter{Repository: &role.RoleRepository{Data: data.New()}}

	r.Mount("/accounts", ar.Routes())
	r.Mount("/clients", cr.Routes())
	r.Mount("/roles", rr.Routes())

	return r
}

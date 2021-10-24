package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"project/internal/models"
	"project/internal/store"
	"strconv"
	"time"
)

type Server struct {
	ctx context.Context
	idleConnsCh chan struct{}
	store store.Store

	Address string
}

func NewServer(ctx context.Context, address string, store store.Store) *Server {
	return &Server{
		ctx: ctx,
		idleConnsCh: make(chan struct{}),
		store: store,

		Address: address,
	}
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	r.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		product := new(models.ProductDto)
		if err := json.NewDecoder(r.Body).Decode(product); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Products().Create(r.Context(), product)
		render.JSON(w, r, product)
	})
	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		catIdStr := r.URL.Query().Get("category-id")
		if len(catIdStr) == 0 {
			catIdStr = "0"
		}
		catId, err := strconv.Atoi(catIdStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		brIdStr := r.URL.Query().Get("brand-id")
		if len(brIdStr) == 0 {
			brIdStr = "0"
		}
		brId, err := strconv.Atoi(brIdStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		products, err := s.store.Products().All(r.Context(), catId, brId)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, products)
	})
	r.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		product, err := s.store.Products().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, product)
	})
	r.Put("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		product := new(models.ProductDto)
		if err := json.NewDecoder(r.Body).Decode(product); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Products().Update(r.Context(), product, id)
		render.JSON(w, r, product)
	})
	r.Delete("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Products().Delete(r.Context(), id)
	})


	r.Post("/categories", func(w http.ResponseWriter, r *http.Request) {
		category := new(models.Category)
		if err := json.NewDecoder(r.Body).Decode(category); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Categories().Create(r.Context(), category)
		render.JSON(w, r, category)
	})
	r.Get("/categories", func(w http.ResponseWriter, r *http.Request) {
		categories, err := s.store.Categories().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, categories)
	})
	r.Get("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		category, err := s.store.Categories().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, category)
	})
	r.Put("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		category := new(models.Category)
		if err := json.NewDecoder(r.Body).Decode(category); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Categories().Update(r.Context(), category, id)
		render.JSON(w, r, category)
	})
	r.Delete("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Categories().Delete(r.Context(), id)
	})


	r.Post("/brands", func(w http.ResponseWriter, r *http.Request) {
		brand := new(models.Brand)
		if err := json.NewDecoder(r.Body).Decode(brand); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Brands().Create(r.Context(), brand)
		render.JSON(w, r, brand)
	})
	r.Get("/brands", func(w http.ResponseWriter, r *http.Request) {
		brands, err := s.store.Brands().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, brands)
	})
	r.Get("/brands/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		brand, err := s.store.Brands().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, brand)
	})
	r.Put("/brands/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		brand := new(models.Brand)
		if err := json.NewDecoder(r.Body).Decode(brand); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Brands().Update(r.Context(), brand, id)
		render.JSON(w, r, brand)
	})
	r.Delete("/brands/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Brands().Delete(r.Context(), id)
	})


	r.Post("/cart-items", func(w http.ResponseWriter, r *http.Request) {
		cartItem := new(models.CartItem)
		if err := json.NewDecoder(r.Body).Decode(cartItem); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.CartItems().Create(r.Context(), cartItem)
		render.JSON(w, r, cartItem)
	})
	r.Get("/cart-items/{userId}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "userId")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		cartItems, err := s.store.CartItems().ListByUserId(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, cartItems)
	})
	r.Put("/cart-items/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		cartItem := new(models.CartItem)
		if err := json.NewDecoder(r.Body).Decode(cartItem); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.CartItems().Update(r.Context(), cartItem, id)
		render.JSON(w, r, cartItem)
	})
	r.Delete("/cart-items/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.CartItems().Delete(r.Context(), id)
	})

	r.Post("/orders", func(w http.ResponseWriter, r *http.Request) {
		order := new(models.OrderDto)
		if err := json.NewDecoder(r.Body).Decode(order); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		orderFull, err := s.store.Orders().Create(r.Context(), order)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}
		render.JSON(w, r, orderFull)
	})
	r.Get("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		orders, err := s.store.Orders().ListByUserId(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, orders)
	})

	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr: s.Address,
		Handler: s.basicHandler(),
		ReadTimeout: time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	<-s.idleConnsCh
}

package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/thaessaep/billingSystem/internal/model"
	"github.com/thaessaep/billingSystem/internal/storage"
)

type HttpServer struct {
	config  *Config
	router  *mux.Router
	storage storage.Storage
}

func New(config *Config) *HttpServer {
	return &HttpServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *HttpServer) Start() error {
	println("Start server <3")
	s.configureRouter()

	err := s.configureStorage()
	if err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *HttpServer) configureRouter() {
	s.router.HandleFunc("/addBalance", s.addBalance())
	s.router.HandleFunc("/getBalance", s.getBalance())
	s.router.HandleFunc("/reserve", s.reserve())
	s.router.HandleFunc("/report", s.report())
}

func (s *HttpServer) configureStorage() error {
	st := storage.New(s.config.Storage)

	if err := st.Open(); err != nil {
		return err
	}

	s.storage = *st

	return nil
}

// @Tags AddBalance
// @Accept json
// @Produce json
// @Param input body model.User true "Balance info"
// @Success 200
// @Failure 400,422
// @Router /addBalance [post]
func (s *HttpServer) addBalance() http.HandlerFunc {
	type request struct {
		Id      int `json:"user_id"`
		Balance int `json:"balance"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			UserId:  req.Id,
			Balance: req.Balance,
		}

		if err := s.storage.User().AddBalance(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		user, err := s.storage.User().FindById(u.UserId)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusAccepted, user)
	}
}

// @Tags GetBalance
// @Accept json
// @Produce json
// @Description Use without balance!
// @Param input body model.User true "user_id"
// @Success 200
// @Failure 400,422
// @Router /getBalance [post]
func (s *HttpServer) getBalance() http.HandlerFunc {
	type request struct {
		Id int `json:"user_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			UserId: req.Id,
		}

		balance, err := s.storage.User().FindBalanceById(u.UserId)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusAccepted, balance)
	}
}

// @Tags Reserve
// @Accept json
// @Produce json
// @Description Can use without balance, success
// @Param input body model.ReserveBills true "Reserve info"
// @Success 200
// @Failure 400
// @Router /reserve [post]
func (s *HttpServer) reserve() http.HandlerFunc {
	type request struct {
		Success   *bool `json:"success"`
		OrderId   int   `json:"order_id"`
		ServiceId int   `json:"service_id"`
		Cost      int   `json:"cost"`
		UserId    int   `json:"user_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		rB := &model.ReserveBills{
			OrderId:   req.OrderId,
			ServiceId: req.ServiceId,
			Cost:      req.Cost,
			User: model.User{
				UserId: req.UserId,
			},
			Success:  req.Success,
			Datetime: time.Now().Unix(),
		}

		if err := s.storage.ReserveBills().AddReserveBill(rB); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusAccepted, rB)
	}
}

// @Tags Report
// @Accept json
// @Produce json
// @Param year path integer true "Year"
// @Param month path integer true "Month"
// @Success 200
// @Failure 400
// @Router /report [post]
func (s *HttpServer) report() http.HandlerFunc {
	type request struct {
		Year  int `json:"year"`
		Month int `json:"month"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		report, err := s.storage.ReserveBills().Report(req.Year, req.Month)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		s.respond(w, r, http.StatusAccepted, report)
	}
}

func (s *HttpServer) error(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	s.respond(w, r, statusCode, map[string]string{"error": err.Error()})
}

func (s *HttpServer) respond(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

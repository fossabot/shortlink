package balance

import (
	"net/http"

	"github.com/segmentio/encoding/json"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/encoding/protojson"

	payment_application "github.com/batazor/shortlink/internal/services/billing/application/payment"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
)

type BalanceAPI struct {
	jsonpb protojson.MarshalOptions // nolint:structcheck

	paymentService *payment_application.PaymentService
}

func New(paymentService *payment_application.PaymentService) (*BalanceAPI, error) {
	return &BalanceAPI{
		paymentService: paymentService,
	}, nil
}

// Routes creates a REST router
func (api *BalanceAPI) Routes(r chi.Router) {
	r.Put("/balance/{payment_id}", api.update)
}

// Update ...
func (api *BalanceAPI) update(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// inject spanId in response header
	w.Header().Add("trace-id", trace.LinkFromContext(r.Context()).SpanContext.TraceID().String())

	// Parse request
	decoder := json.NewDecoder(r.Body)
	var request billing.Payment
	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	updatePayment, err := api.paymentService.UpdateBalance(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := api.jsonpb.Marshal(updatePayment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint:errcheck
}

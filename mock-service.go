package mockservice

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Mock certificate response
type MockCertificateResponse struct {
	CertificatePEM string `json:"certificate_pem"`
	CAPEM          string `json:"ca_pem"`
}

func main() {
	http.HandleFunc("/api/v1/certificates", func(w http.ResponseWriter, r *http.Request) {
		// Handle certificate request (could include CSR validation, etc.)
		if r.Method == http.MethodPost {
			// Simulate certificate issuance delay
			time.Sleep(2 * time.Second)

			// Always return 200 OK with a self-signed certificate
			mockCert := MockCertificateResponse{
				CertificatePEM: "-----BEGIN CERTIFICATE-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7G8wGGF/ykN/h0zUEMKT\n-----END CERTIFICATE-----",
				CAPEM:          "-----BEGIN CERTIFICATE-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7G8wGGF/ykN/h0zUEMKT\n-----END CERTIFICATE-----",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mockCert)
			return
		}

		// Return 405 for unsupported methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("Starting mock server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

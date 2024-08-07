package controllers

import (
	"alpha-issuer/api"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"

	"github.com/cert-manager/issuer-lib/api/v1alpha1"
	"github.com/cert-manager/issuer-lib/controllers"
	"github.com/cert-manager/issuer-lib/controllers/signer"
)

var apiURL = "http://localhost:8080/api/v1/certificates"

// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests,verbs=get;list;watch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificaterequests/status,verbs=patch

// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=get;list;watch
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests/status,verbs=patch
// +kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,verbs=sign,resourceNames=alphaissuers.certmanager.alpha-issuer.io/*;alphaclusterissuers.certmanager.alpha-issuer.io/*

// +kubebuilder:rbac:groups=alphaissuers.certmanager.alpha-issuer.io,resources=aplhaissuers;alphaclusterissuers,verbs=get;list;watch
// +kubebuilder:rbac:groups=alphaissuers.certmanager.alpha-issuer.io,resources=alphaissuers/status;alphaclusterissuers/status,verbs=patch

// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch

type Signer struct{}

func (s Signer) SetupWithManager(ctx context.Context, mgr ctrl.Manager) error {
	return (&controllers.CombinedController{
		IssuerTypes:        []v1alpha1.Issuer{&api.AlphaIssuer{}},
		ClusterIssuerTypes: []v1alpha1.Issuer{&api.AlphaClusterIssuer{}},

		FieldOwner:       "alphaissuers.certmanager.alpha-issuer.io",
		MaxRetryDuration: 1 * time.Minute,

		Sign:          s.Sign,
		Check:         s.Check,
		EventRecorder: mgr.GetEventRecorderFor("alphaissuers.certmanager.alpha-issuer.io"),
	}).SetupWithManager(ctx, mgr)
}

func (Signer) Check(ctx context.Context, issuerObject v1alpha1.Issuer) error {
	return nil
}

type ApiCSR struct {
	CSR []byte
}

type ApiIssuedCert struct {
	CertificatePEM []byte
	CAPEM          []byte
}

func (Signer) Sign(ctx context.Context, cr signer.CertificateRequestObject, issuerObject v1alpha1.Issuer) (signer.PEMBundle, error) {

	_, _, csr, err := cr.GetRequest()

	// Prepare your REST API request using the certificate request data
	apiCSR := &ApiCSR{
		CSR: csr,
		// Populate other fields as necessary
	}

	// Send the request to your CA
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(apiCSR.CSR))
	if err != nil {
		return signer.PEMBundle{}, err
	}
	defer resp.Body.Close()

	// Handle the response and update the CertificateRequest status
	if resp.StatusCode == http.StatusOK {
		// Parse the response body to get the issued certificate
		var issuedCert ApiIssuedCert
		if err := json.NewDecoder(resp.Body).Decode(&issuedCert); err != nil {
			return signer.PEMBundle{}, err
		}

		return signer.PEMBundle{
			ChainPEM: issuedCert.CertificatePEM,
		}, nil
	} else {
		return signer.PEMBundle{}, fmt.Errorf("failed to issue certificate: %v", resp.Status)
	}
}

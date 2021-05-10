package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	fluxcdv1 "github.com/fluxcd/helm-operator/pkg/apis/helm.fluxcd.io/v1"
	"github.com/tom721/from-helm-server/internal/utils"
	"github.com/tom721/from-helm-server/pkg/schema"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var logger = logf.Log.WithName("HelmRelease")

func HelmRelease(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hb schema.HelmBody

	//get body
	if err := json.NewDecoder(r.Body).Decode(&hb); err != nil {
		respondError(w, http.StatusInternalServerError, &schema.Error{
			Error:       "InternalServerError",
			Description: "Error occurs while decoding request body",
		})
		return
	}

	defer r.Body.Close()

	// initialize client
	s := scheme.Scheme
	if err := fluxcdv1.AddToScheme(s); err != nil {
		respondError(w, http.StatusInternalServerError, &schema.Error{
			Error:       "InternalServerError",
			Description: "Error occurs while add HelmRelease scheme",
		})
		return
	}

	c, err := utils.Client(s)
	if err != nil {
		respondError(w, http.StatusInternalServerError, &schema.Error{
			Error:       "InternalServerError",
			Description: "Error occurs while connect to k8s api server",
		})
		return
	}

	if isHelmReleaseExist(c, types.NamespacedName{
		Namespace: hb.NameSpace,
		Name:      hb.Name,
	}) {
		respondError(w, http.StatusInternalServerError, &schema.Error{
			Error:       "InternalServerError",
			Description: fmt.Sprintf("HelmRelease %s already exists in %s namespace", hb.Name, hb.NameSpace),
		})
		return
	}

	logger.Info("start createHelmRelease API")
	if err := createHelmRelease(c, hb); err != nil {
		logger.Error(err, "occured")
		respondError(w, http.StatusInternalServerError, &schema.Error{
			Error:       "InternalServerError",
			Description: "Error occurs while creates HelmRelease resource",
		})
		return
	}
}

func respondError(w http.ResponseWriter, statusCode int, message *schema.Error) {
	logger.Error(fmt.Errorf(message.Description), "error occurred")

	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		logger.Error(err, "error occurs while encoding error message to json format")
	}
}

func isHelmReleaseExist(c client.Client, name types.NamespacedName) bool {
	hr := &fluxcdv1.HelmRelease{}
	if err := c.Get(context.TODO(), name, hr); err != nil && errors.IsNotFound(err) {
		return false
	}
	return true
}

func createHelmRelease(c client.Client, hb schema.HelmBody) error {
	logger.Info("start define HR...")
	helmRelease := &fluxcdv1.HelmRelease{
		ObjectMeta: metav1.ObjectMeta{
			Name:      hb.Name,
			Namespace: hb.NameSpace,
		},
		Spec: fluxcdv1.HelmReleaseSpec{
			ChartSource: fluxcdv1.ChartSource{
				RepoChartSource: &fluxcdv1.RepoChartSource{
					RepoURL: hb.Chart["repository"],
					Name:    hb.Chart["name"],
					Version: hb.Chart["version"],
				},
			},
			Values: fluxcdv1.HelmValues{
				Data: hb.Values,
			},
		},
	}

	logger.Info("start create HR..")
	if err := c.Create(context.TODO(), helmRelease); err != nil {
		return err
	}

	return nil
}

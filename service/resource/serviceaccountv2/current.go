package serviceaccountv2

import (
	"context"

	"github.com/giantswarm/microerror"
	apiv1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/kvm-operator/service/keyv2"
)

func (r *Resource) GetCurrentState(ctx context.Context, obj interface{}) (interface{}, error) {
	customObject, err := keyv2.ToCustomObject(obj)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	r.logger.LogCtx(ctx, "debug", "looking for a service account in the Kubernetes API")

	namespace := keyv2.ClusterNamespace(customObject)
	var currentServiceAccount *apiv1.ServiceAccount
	currentServiceAccount, err = r.k8sClient.CoreV1().ServiceAccounts(namespace).Get(keyv2.ServiceAccountName(customObject), apismetav1.GetOptions{})
	if err != nil {
		return nil, microerror.Mask(err)
	}

	r.logger.LogCtx(ctx, "debug", "found a service account in the Kubernetes API")

	return currentServiceAccount, nil
}
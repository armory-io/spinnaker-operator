package deployer

import (
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

func TestFindLoadBalancerUrlStandardPort(t *testing.T) {
	svc := buildSvc("spin-deck", "LoadBalancer", nil)
	svc.Status.LoadBalancer.Ingress = []corev1.LoadBalancerIngress{
		{Hostname: "abc.com"},
	}
	fakeClient := fake.NewFakeClient(svc)
	lbUrl, err := FindLoadBalancerUrl("spin-deck", "ns1", fakeClient)
	assert.Nil(t, err)
	assert.Equal(t, "http://abc.com", lbUrl)
}

func TestFindLoadBalancerUrlCustomPort(t *testing.T) {
	svc := buildSvc("spin-deck", "LoadBalancer", nil)
	svc.Spec.Ports[0].Port = 8084
	svc.Status.LoadBalancer.Ingress = []corev1.LoadBalancerIngress{
		{Hostname: "abc.com"},
	}
	fakeClient := fake.NewFakeClient(svc)
	lbUrl, err := FindLoadBalancerUrl("spin-deck", "ns1", fakeClient)
	assert.Nil(t, err)
	assert.Equal(t, "http://abc.com:8084", lbUrl)
}

func TestFindLoadBalancerUrlHttps(t *testing.T) {
	svc := buildSvc("spin-deck", "LoadBalancer", nil)
	svc.Spec.Ports[0].Port = 443
	svc.Status.LoadBalancer.Ingress = []corev1.LoadBalancerIngress{
		{Hostname: "abc.com"},
	}
	fakeClient := fake.NewFakeClient(svc)
	lbUrl, err := FindLoadBalancerUrl("spin-deck", "ns1", fakeClient)
	assert.Nil(t, err)
	assert.Equal(t, "https://abc.com", lbUrl)
}

func TestFindLoadBalancerUrlIP(t *testing.T) {
	svc := buildSvc("spin-deck", "LoadBalancer", nil)
	svc.Status.LoadBalancer.Ingress = []corev1.LoadBalancerIngress{
		{Hostname: "", IP: "10.0.0.14"},
	}
	fakeClient := fake.NewFakeClient(svc)
	lbUrl, err := FindLoadBalancerUrl("spin-deck", "ns1", fakeClient)
	assert.Nil(t, err)
	assert.Equal(t, "http://10.0.0.14", lbUrl)
}
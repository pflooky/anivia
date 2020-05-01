package task

import (
	"github.com/kube-sailmaker/template-gen/model"
	"github.com/kube-sailmaker/template-gen/test"
	"testing"
)

func TestGetResource(t *testing.T) {
	resourceDir := "../sample-manifest/provider"
	resource := &model.Resource{}
	err := GetResource("cassandra", resource, resourceDir)
	test.Null(t, err)
	test.NotNull(t, resource)
	test.EqualTo(t, "Resource", resource.Kind)
	test.EqualTo(t, "v1", resource.ApiVersion)
	test.EqualTo(t, "cassandra-a", resource.Metadata["name"])
	test.EqualTo(t, "test1", resource.Spec.ResourceTemplate[0].Name)
}

func TestGetInfrastructure(t *testing.T) {
	resourceDir := "../sample-manifest/provider"
	infrastructure := &model.Infrastructure{}
	err := GetInfrastructure("cassandra-cluster-a", infrastructure, resourceDir)
	test.Null(t, err)
	test.NotNull(t, infrastructure)
	test.EqualTo(t, "v1", infrastructure.ApiVersion)
	test.EqualTo(t, "Infrastructure", infrastructure.Kind)
	test.EqualTo(t, "cassandra-a", infrastructure.Metadata["name"])
	test.EqualTo(t, "test", infrastructure.Spec.Template[0].Name)
}

func TestGetMixin(t *testing.T) {
	resourceDir := "../sample-manifest/provider"
	mixinList := &model.MixinList{}
	err := GetMixin("java", mixinList, resourceDir)
	test.Null(t, err)
	test.NotNull(t, mixinList)
	test.EqualTo(t, "java-default", mixinList.Mixin[0].Name)
}

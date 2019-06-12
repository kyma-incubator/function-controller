package utils_test

import (
	"testing"

	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
)

func TestGetServiceSpec(t *testing.T) {
	//imageName := "foo-image"
	//fn := runtimev1alpha1.Function{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name: "foo",
	//	},
	//	Spec: runtimev1alpha1.FunctionSpec{
	//		Function:            "main() {}",
	//		FunctionContentType: "plaintext",
	//		Size:                "L",
	//		Runtime:             "nodejs8",
	//	},
	//}
	//
	//rnInfo := &utils.RuntimeInfo{
	//	RegistryInfo: "test",
	//	AvailableRuntimes: []utils.RuntimesSupported{
	//		{
	//			ID:             "nodejs8",
	//			DockerFileName: "testnodejs8",
	//		},
	//	},
	//}
	//serviceSpec := utils.GetServiceSpec(imageName, fn, rnInfo)
	//build := (*serviceSpec.RunLatest.Configuration.Build)
	//tempBuildByte, err := build.MarshalJSON()
	//if err != nil {
	//	t.Fatalf("Error while marshaling build object: %v", err)
	//}
	//
	//// Testing BuildSpec
	//var buildSpec buildv1alpha1.BuildSpec
	//err = json.Unmarshal(tempBuildByte, &buildSpec)
	//if err != nil {
	//	t.Fatalf("Error while unmarshaling buildSpec: %v", err)
	//}
	//
	//if len(buildSpec.Steps[0].Args) != 2 {
	//	t.Fatalf("Expected length of args: %d Got: %d", 2, len(buildSpec.Steps[0].Args))
	//}
	//
	//for _, vol := range buildSpec.Steps[0].VolumeMounts {
	//	if vol.Name != "dockerfile-vol" && vol.Name != "func-vol" {
	//		t.Fatalf("Got incorrect values for volumemounts names: %v", vol.Name)
	//	}
	//	if vol.MountPath != "/workspace" && vol.MountPath != "/src" {
	//		t.Fatalf("Got incorrect values for volumemounts mountpaths: %v", vol.MountPath)
	//	}
	//}
	//
	//for _, vol := range buildSpec.Volumes {
	//	if vol.Name != "dockerfile-vol" && vol.Name != "func-vol" {
	//		t.Fatalf("Got incorrect values for build.spec.volumes.names: %v", vol.Name)
	//	}
	//	if vol.ConfigMap.Name != "foo" && vol.ConfigMap.Name != "testnodejs8" {
	//		t.Fatalf("Got incorrect values for build.spec.volumes.configmap.name: %v", vol.ConfigMap.Name)
	//	}
	//}
	//
	//for _, v := range buildSpec.Steps[0].Args {
	//	if v != "--dockerfile=/workspace/Dockerfile" && v != "--destination=foo-image" {
	//		t.Fatalf("Got an unacceptable buildSpec.Steps[0].Args in args %s", v)
	//	}
	//}
	//
	//if buildSpec.Steps[0].Image != "gcr.io/kaniko-project/executor" {
	//	t.Fatalf("Expected build.Spec.Image: %v Got: %v", "gcr.io/kaniko-project/executor", buildSpec.Steps[0].Image)
	//}
	//
	//if buildSpec.Steps[0].Name != "build-and-push" {
	//	t.Fatalf("Expected build.Spec.Name: %v Got: %v", "build-and-push", buildSpec.Steps[0].Name)
	//}
	//
	//// Testing Revision
	//if serviceSpec.RunLatest.Configuration.RevisionTemplate.Spec.Container.Image != "foo-image" {
	//	t.Fatalf("Expected image for RevisionTemplate.Spec.Container.Image: %v Got: %v", "foo-image", serviceSpec.RunLatest.Configuration.RevisionTemplate.Spec.Container.Image)
	//}
	//expectedEnv := []corev1.EnvVar{
	//	{
	//		Name:  "FUNC_HANDLER",
	//		Value: "main",
	//	},
	//	{
	//		Name:  "MOD_NAME",
	//		Value: "handler",
	//	},
	//	{
	//		Name:  "FUNC_TIMEOUT",
	//		Value: "180",
	//	},
	//	{
	//		Name:  "FUNC_RUNTIME",
	//		Value: "nodejs8",
	//	},
	//	{
	//		Name:  "FUNC_MEMORY_LIMIT",
	//		Value: "128Mi",
	//	},
	//	{
	//		Name:  "FUNC_PORT",
	//		Value: "8080",
	//	},
	//	{
	//		Name:  "NODE_PATH",
	//		Value: "$(KUBELESS_INSTALL_VOLUME)/node_modules",
	//	},
	//}
	//if !compareEnv(t, expectedEnv, serviceSpec.RunLatest.Configuration.RevisionTemplate.Spec.Container.Env) {
	//	expectedEnvStr, err := getString(expectedEnv)
	//	gotEnvStr, err := getString(expectedEnv)
	//	t.Fatalf("Expected value in Env: %v Got: %v", expectedEnvStr, gotEnvStr)
	//	if err != nil {
	//		t.Fatalf("Error while unmarshaling expectedBuildSpec: %v", err)
	//	}
	//}
}

func compareEnv(t *testing.T, source, dest []corev1.EnvVar) bool {
	for i, _ := range source {
		found := false
		for j, _ := range dest {
			if source[i].Name == dest[j].Name && source[i].Value == dest[j].Value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func getString(obj interface{}) (string, error) {
	output, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

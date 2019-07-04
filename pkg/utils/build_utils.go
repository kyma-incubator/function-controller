package utils

import (
	"os"
	"time"

	buildv1alpha1 "github.com/knative/build/pkg/apis/build/v1alpha1"
	runtimev1alpha1 "github.com/kyma-incubator/runtime/pkg/apis/runtime/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/*
apiVersion: build.knative.dev/v1alpha1
kind: Build
metadata:
  name: sample
spec:
  serviceAccountName: runtime-controller
  template:
    name: function-kaniko
    kind: BuildTemplate
    arguments:
      - name: IMAGE
        value: sample:latest
      - name: DOCKERFILE
        value: dockerfile-nodejs-8
  volumes:
  - configmap:
      name: sample
    name: source
*/

type Build struct {
	Name               string
	Namespace          string
	ServiceAccountName string
	BuildtemplateName  string
	Args               map[string]string
	Envs               map[string]string
	Timeout            time.Duration
}

var dockerFileName = map[string]string{
	"nodejs6": "dockerfile-nodejs-6",
	"nodejs8": "dockerfile-nodejs-8",
}

var buildTimeout = os.Getenv("BUILD_TIMEOUT")

var defaultMode = int32(420)

func NewBuild(rnInfo *RuntimeInfo, fn *runtimev1alpha1.Function, imageName string, buildTemplateName string) *Build {

	argsMap := make(map[string]string)
	argsMap["IMAGE"] = imageName
	argsMap["DOCKERFILE"] = dockerFileName[fn.Spec.Runtime]

	envMap := make(map[string]string)

	timeout, err := time.ParseDuration(buildTimeout)
	if err != nil {
		timeout = 30 * time.Minute
	}

	// TODO: do we need the extra build struct or not ?
	return &Build{
		Name:               fn.Name,
		Namespace:          fn.Namespace,
		ServiceAccountName: rnInfo.ServiceAccount,
		BuildtemplateName:  buildTemplateName,
		Args:               argsMap,
		Envs:               envMap,
		Timeout:            timeout,
	}
}

func GetBuildResource(build *Build, fn *runtimev1alpha1.Function) *buildv1alpha1.Build {

	args := []buildv1alpha1.ArgumentSpec{}
	for k, v := range build.Args {
		args = append(args, buildv1alpha1.ArgumentSpec{Name: k, Value: v})
	}

	envs := []corev1.EnvVar{}
	for k, v := range build.Envs {
		envs = append(envs, corev1.EnvVar{Name: k, Value: v})
	}

	vols := []corev1.Volume{
		{
			Name: "source",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: &defaultMode,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: fn.Name,
					},
				},
			},
		},
	}

	b := buildv1alpha1.Build{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "build.knative.dev/v1alpha1",
			Kind:       "Build",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      build.Name,
			Namespace: build.Namespace,
			Labels:    fn.Labels,
		},
		Spec: buildv1alpha1.BuildSpec{
			ServiceAccountName: build.ServiceAccountName,
			Template: &buildv1alpha1.TemplateInstantiationSpec{
				Name:      "function-kaniko",
				Kind:      buildv1alpha1.BuildTemplateKind,
				Arguments: args,
				Env:       envs,
			},
			Volumes: vols,
		},
	}

	if b.Spec.Timeout == nil {
		b.Spec.Timeout = &metav1.Duration{Duration: build.Timeout}
	}

	return &b
}

func GetBuildTemplateSpec(fn *runtimev1alpha1.Function) buildv1alpha1.BuildTemplateSpec {

	parameters := []buildv1alpha1.ParameterSpec{
		{
			Name:        "IMAGE",
			Description: "The name of the image to push",
		},
		{
			Name:        "DOCKERFILE",
			Description: "name of the configmap that contains the Dockerfile",
		},
	}

	destination := "--destination=${IMAGE}"
	steps := []corev1.Container{
		{
			Name:  "build-and-push",
			Image: "gcr.io/kaniko-project/executor",
			Args: []string{
				"--dockerfile=/workspace/Dockerfile",
				destination,
			},
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      "${DOCKERFILE}",
					MountPath: "/workspace",
				},
				{
					Name:      "source",
					MountPath: "/src",
				},
			},
		},
	}

	volumes := []corev1.Volume{
		{
			Name: "dockerfile-nodejs-6",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: &defaultMode,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "dockerfile-nodejs-6",
					},
				},
			},
		},
		{
			Name: "dockerfile-nodejs-8",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: &defaultMode,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "dockerfile-nodejs-8",
					},
				},
			},
		},
	}

	bt := buildv1alpha1.BuildTemplateSpec{
		Parameters: parameters,
		Steps:      steps,
		Volumes:    volumes,
	}

	return bt
}


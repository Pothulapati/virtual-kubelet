package framework

import (
	corev1 "k8s.io/api/core/v1"
)


//Create a dummy pod with mandatory configmaps
func (f *Framework) CreatePodWithMandatoryConfigMaps() (*corev1.Pod){

	//Check with both Optional and Mandatory ConfigMaps
	var mandatory bool = false
	var optional bool = true
	env := []corev1.EnvVar {
		{
			Name: "zero",
			ValueFrom: &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{Name: "Configname0"},
					Key:"key0",
					Optional: &optional ,
				},
			},

		},
		{
			Name: "one",
			ValueFrom: &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{Name: "Configname1"},
					Key:"key1",
					Optional: &mandatory,
				},

			},
		},
	}

	return f.CreatePodWithEnv(env)
}


//Create a dummy pod with optional configmaps
func (f *Framework) CreatePodWithOptionalConfigMaps() (*corev1.Pod){
	//Check with both Optional ConfigMaps
	var optional bool = true
	env := []corev1.EnvVar {
		{
			Name: "ENVIRONMENTVARIABLE",
			ValueFrom: &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{Name: "configname0"},
					Key:"key",
					Optional:&optional,
				},
			},

		},
	}

	return f.CreatePodWithEnv(env)
}


//Create a dummy pod with optional Secrets
func (f *Framework) CreatePodWithOptionalSecrets() (*corev1.Pod){
	//Check with both Optional
	//var mandatory bool = false
	var optional bool = true
	env := []corev1.EnvVar {
		{
			Name: "ENVIRONMENTVARIABLE",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{Name: "secretname0"},
					Key:"key",
					Optional:&optional,
				},
			},

		},
	}

	return f.CreatePodWithEnv(env)
}

//Create a dummy pod with optional Secrets
func (f *Framework) CreatePodWithMandatorySecrets() (*corev1.Pod){
	//Check with both Optional and not optional Secrets
	var mandatory bool = false
	var optional bool = true
	env := []corev1.EnvVar {
		{
			Name: "zero",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{Name: "secretname0"},
					Key:"secretkey0",
					Optional: &optional ,
				},
			},

		},
		{
			Name: "one",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{Name: "secretname1"},
					Key:"secretkey1",
					Optional: &mandatory,
				},

			},
		},
	}
	return f.CreatePodWithEnv(env)
}

//Creates a dummy pod with the environment variables passed
func (f *Framework) CreatePodWithEnv(env []corev1.EnvVar) (*corev1.Pod)  {

	pod := f.CreateDummyPodObjectWithPrefix("nginxtest","foo")
	pod.Spec.Containers[0].Env = env
	return pod
}



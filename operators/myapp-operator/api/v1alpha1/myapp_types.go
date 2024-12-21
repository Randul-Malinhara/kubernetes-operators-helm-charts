package v1alpha1

import (
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MyAppSpec struct {
  Size     int32  `json:"size"`
  LogLevel string `json:"logLevel"`
}

type MyAppStatus struct {
  Nodes []string `json:"nodes"`
}

type MyApp struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec   MyAppSpec   `json:"spec,omitempty"`
  Status MyAppStatus `json:"status,omitempty"`
}

type MyAppList struct {
  metav1.TypeMeta `json:",inline"`
  metav1.ListMeta `json:"metadata,omitempty"`
  Items           []MyApp `json:"items"`
}

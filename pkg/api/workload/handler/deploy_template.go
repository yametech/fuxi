package handler

type metadataTemplate []struct {
	Base struct {
		Name            string `json:"name"`
		Image           string `json:"image"`
		ImagePullPolicy string `json:"imagePullPolicy"`
		Resource        struct {
			Limits struct {
				CPU    string `json:"cpu"`
				Memory string `json:"memory"`
			} `json:"limits"`
			Requests struct {
				CPU    string `json:"cpu"`
				Memory string `json:"memory"`
			} `json:"requests"`
		} `json:"resource"`
	} `json:"base"`
	Commands    []string `json:"commands"`
	Args        []string `json:"args"`
	Environment []struct {
		Type         string `json:"type"`
		OneEnvConfig struct {
			Name          string `json:"name"`
			ConfigureName string `json:"configureName"`
			SecretName    string `json:"secretName"`   // Secret name
			SecretKey     string `json:"secretKey"`    // Secret key
			EnterCommand  string `json:"enterCommand"` //other
			Key           string `json:"key"`
			Value         string `json:"value"`
		} `json:"oneEnvConfig"`
	} `json:"environment"`
	ReadyProbe struct {
		Status     bool   `json:"status"`
		Timeout    string `json:"timeout"`
		Cycle      string `json:"cycle"`
		RetryCount string `json:"retryCount"`
		Delay      string `json:"delay"`
		Pattern    struct {
			Type     string `json:"type"`
			HTTPPort string `json:"httpPort"`
			URL      string `json:"url"`
			TCPPort  string `json:"tcpPort"`
			Command  string `json:"command"`
		} `json:"pattern"`
	} `json:"readyProbe"`
	LiveProbe struct {
		Status     bool   `json:"status"`
		Timeout    string `json:"timeout"`
		Cycle      string `json:"cycle"`
		RetryCount string `json:"retryCount"`
		Delay      string `json:"delay"`
		Pattern    struct {
			Type     string `json:"type"`
			HTTPPort string `json:"httpPort"`
			URL      string `json:"url"`
			TCPPort  string `json:"tcpPort"`
			Command  string `json:"command"`
		} `json:"pattern"`
	} `json:"liveProbe"`
	LifeCycle struct {
		Status    bool `json:"status"`
		PostStart struct {
			Type     string `json:"type"`
			HTTPPort string `json:"httpPort"`
			URL      string `json:"url"`
			TCPPort  string `json:"tcpPort"`
			Command  string `json:"command"`
		} `json:"postStart"`
		PreStop struct {
			Type     string `json:"type"`
			HTTPPort string `json:"httpPort"`
			URL      string `json:"url"`
			TCPPort  string `json:"tcpPort"`
			Command  string `json:"command"`
		} `json:"preStop"`
	} `json:"lifeCycle"`
	VolumeMounts struct {
		Status bool `json:"status"`
		Items  []struct {
			Name      string `json:"name"`
			MountPath string `json:"mountPath"`
		} `json:"items"`
	} `json:"volumeMounts"`
}

type serviceTemplate struct {
	Type  string `json:"type"`
	Ports []struct {
		Name       string `json:"name"`
		Protocol   string `json:"protocol"`
		Port       string `json:"port"`
		TargetPort string `json:"targetPort"`
	} `json:"ports"`
}

type volumeClaimsTemplate []struct {
	Metadata struct {
		IsUseDefaultStorageClass bool   `json:"isUseDefaultStorageClass"`
		Name                     string `json:"name"`
		Annotations              struct {
			VolumeAlphaKubernetesIoStorageClass string `json:"volume.alpha.kubernetes.io/storage-class"`
		} `json:"annotations"`
	} `json:"metadata"`
	Spec struct {
		AccessModes      []string `json:"accessModes"`
		StorageClassName string   `json:"storageClassName"`
		Resources        struct {
			Requests struct {
				Storage string `json:"storage"`
			} `json:"requests"`
		} `json:"resources"`
	} `json:"spec"`
}

type workloadsTemplate struct {
	Metadata     metadataTemplate     `json:"metadata"`
	Service      serviceTemplate      `json:"service"`
	VolumeClaims volumeClaimsTemplate `json:"volumeClaims"`
}

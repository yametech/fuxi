package ops

import (
	pipelineClient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1alpha1"
	informers "github.com/tektoncd/pipeline/pkg/client/informers/externalversions"
	"github.com/yametech/fuxi/pkg/db"
	"github.com/yametech/fuxi/pkg/gits"
	"github.com/yametech/fuxi/pkg/tekton"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"time"
)

type OpsService interface {
	ListRepos(username, namespace string) ([]string, error)
	ListBranchs(username, namespace string) ([]string, error)
	LogService
	PipelineService
	ResourceService
	TaskService
	PipelineRunService
}

type Ops struct {
	client pipelineClient.TektonV1alpha1Interface
	log *Logger
	informer informers.SharedInformerFactory
}

func NewOps(defaultResync time.Duration) *Ops {
	return &Ops{
		client: tekton.TektonClient.TektonV1alpha1(),
		log:new(Logger),
		informer:informers.NewSharedInformerFactory(tekton.TektonClient, defaultResync),
	}
}


func (ops *Ops)Start(stopCh <-chan struct{}) {
	ops.informer.Start(stopCh)
}


var _ OpsService = (*Ops)(nil)

//ListRepo according username and dep to select all repos
func (o *Ops) ListRepos(username, namespace string) ([]string, error) {
	uid, err := db.FindUserIdByName(username)
	if err != nil {
		return nil, err
	}
	git, err := db.FindGitByUserId(uid)
	if err != nil {
		return nil, err
	}
	gitArgs := &gits.GitArgs{
		Username: git.Username,
		ApiToken: git.Token,
		Url:      git.Url,
	}
	c := gits.NewGiteaClient(gitArgs)

	repos, err := c.ListRepositories(namespace)
	if err != nil {
		return nil, err
	}
	var repoName []string
	for _, value := range repos {
		repoName = append(repoName, value.CloneURL)
	}

	return repoName, nil
}

//ListBranch  according username and repo name select all branch name
func (o *Ops) ListBranchs(username, namespace string) ([]string, error) {
	uid, err := db.FindUserIdByName(username)
	if err != nil {
		return nil, err
	}
	git, err := db.FindGitByUserId(uid)
	if err != nil {
		return nil, err
	}
	c := gits.NewGiteaClient(&gits.GitArgs{
		Username: git.Username,
		ApiToken: git.Token,
		Url:      git.Url,
	})

	branchs, err := c.ListBranchs(namespace)
	if err != nil {
		return nil, err
	}

	return branchs, nil
}

////BuidProject build a project
//func (ops *Ops) BuidProject(ctx context.Context, req *ops.BuildRequest, rsp *ops.BuildResponse) error {
//
//	if err := ops.findPreFixName(req); err != nil {
//		return err
//	}
//
//	gitName := ops.preFixName + string(v1alpha1.PipelineResourceTypeGit)
//	if err := ops.checkIfPipelineResource(req,
//		v1alpha1.PipelineResourceTypeGit,
//		gitName); err != nil {
//		return err
//	}
//
//	imageType := v1alpha1.PipelineResourceTypeImage
//	imageName := ops.preFixName + string(v1alpha1.PipelineResourceTypeImage)
//	if err := ops.checkIfPipelineResource(req,
//		imageType, imageName); err != nil {
//		return err
//	}
//	if err := ops.checkIfCreateTask(gitName, imageName); err != nil {
//		return err
//	}
//	if err := ops.checkIfCreateTaskRun(gitName, imageName); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (ops *Ops) findPreFixName(req *ops.BuildRequest) error {
//	// accroding user name find uid
//	uid, err := db.FindUserIdByName(req.Username)
//	if err != nil {
//		return err
//	}
//	//accroding uid find depart
//	d, err := db.FindDepartment(uid)
//	if err != nil {
//		return err
//	}
//	ops.deptName = d.DeptName
//	ops.preFixName = fmt.Sprintf("%s-%s-%s-", req.Username, d.DeptName, req.Branchname)
//	return nil
//}
//
//func (ops *Ops) PreCheck() error {
//	//check if namcesapce exist,not exist will create it
//	if err := ops.checkIfCreateNameSapce(); err != nil {
//		return err
//	}
//	if err := ops.checkIfCreateSecret(); err != nil {
//		return err
//	}
//	if err := ops.checkIfCreateServiceAccount(); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (ops *Ops) getPipelineResources(name, namesapce string) (*v1alpha1.PipelineResource, error) {
//	res, err := ops.client.PipelineResources(namesapce).
//		Get(name, v1.GetOptions{})
//	if err != nil {
//		return res, err
//	}
//	return res, err
//}
//
////func toPipelineResource(resourceParams []v1alpha1.ResourceParam,
////	pipelineResourceType v1alpha1.PipelineResourceType,
////	name string) *v1alpha1.PipelineResource {
////	return &v1alpha1.PipelineResource{
////		TypeMeta:   v1.TypeMeta{},
////		ObjectMeta: v1.ObjectMeta{Name: name},
////		Spec: v1alpha1.PipelineResourceSpec{
////			Type:         pipelineResourceType,
////			Params:       resourceParams,
////			SecretParams: nil,
////		},
////		Status: v1alpha1.PipelineResourceStatus{},
////	}
////}
//
//func toTaskResource(name string, pipelineResourceType v1alpha1.PipelineResourceType) v1alpha1.TaskResource {
//	return v1alpha1.TaskResource{
//		ResourceDeclaration: v1alpha1.ResourceDeclaration{
//			Name:       name,
//			Type:       pipelineResourceType,
//			TargetPath: "",
//		},
//		// change by dxp 20200114
//		//OutputImageDir: "",
//	}
//}
//
//func toTask(taskResource v1alpha1.TaskResource,
//	taskParamSpec []v1alpha1.ParamSpec,
//	outputTaskResource v1alpha1.TaskResource,
//	buildSetp v1alpha1.Step,
//	name string) *v1alpha1.Task {
//	return &v1alpha1.Task{
//		TypeMeta:   v1.TypeMeta{},
//		ObjectMeta: v1.ObjectMeta{Name: name},
//		Spec: v1alpha1.TaskSpec{
//			Inputs: &v1alpha1.Inputs{
//				Resources: []v1alpha1.TaskResource{taskResource},
//				Params:    taskParamSpec,
//			},
//			Outputs: &v1alpha1.Outputs{
//				Results:   nil,
//				Resources: []v1alpha1.TaskResource{outputTaskResource},
//			},
//			Steps:        []v1alpha1.Step{buildSetp},
//			StepTemplate: nil,
//			Sidecars:     nil,
//		},
//	}
//}
//
//func toTaskRun(inputTaskResourceBinding v1alpha1.TaskResourceBinding,
//	ps []v1alpha1.Param,
//	outPutName string,
//	outPutresourceName string,
//	metaName string,
//	taskName string,
//) *v1alpha1.TaskRun {
//	return &v1alpha1.TaskRun{
//		TypeMeta:   v1.TypeMeta{},
//		ObjectMeta: v1.ObjectMeta{Name: metaName},
//		Spec: v1alpha1.TaskRunSpec{
//			Inputs: v1alpha1.TaskRunInputs{
//				Resources: []v1alpha1.TaskResourceBinding{inputTaskResourceBinding},
//				Params:    ps,
//			},
//			Outputs: v1alpha1.TaskRunOutputs{
//				Resources: []v1alpha1.TaskResourceBinding{{
//					PipelineResourceBinding: v1alpha1.PipelineResourceBinding{
//						Name: outPutresourceName,
//						// change by dxp 20200114
//						ResourceRef: &v1alpha1.PipelineResourceRef{
//							Name: outPutresourceName,
//						},
//					},
//				}},
//			},
//
//			ServiceAccountName: "build-bot",
//			TaskRef: &v1alpha1.TaskRef{
//				Name: taskName,
//			},
//			TaskSpec: nil,
//			Status:   "",
//			Timeout:  nil,
//			//todo: will give a pod  lable and then acording the lable selector node run pileline
//			PodTemplate: v1alpha1.PodTemplate{},
//		},
//		Status: v1alpha1.TaskRunStatus{},
//	}
//}
//
//func (ops *Ops) checkIfCreateNameSapce() error {
//	_, err := kubeclient.K8sClient.CoreV1().Namespaces().Get(ops.namesapce, v1.GetOptions{})
//	if err != nil && errors.IsNotFound(err) {
//		nsSpec := &corev1.Namespace{ObjectMeta: v1.ObjectMeta{Name: ops.namesapce}}
//		_, err := kubeclient.K8sClient.CoreV1().Namespaces().Create(nsSpec)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//	return err
//}
//
//func (ops *Ops) checkIfCreateSecret() error {
//	//todo: hard code,may be config it from etcd
//	s := "basic-user-pass"
//	basics := make(map[string]string)
//	basics["username"] = "withlin"
//	basics["password"] = ""
//	annotations := make(map[string]string)
//	annotations["tekton.dev/git-0"] = "http://10.200.100.219:10080/"
//	annotations["tekton.dev/docker-0"] = "10.200.100.200"
//	_, err := kubeclient.K8sClient.CoreV1().Secrets(ops.namesapce).Get(s, v1.GetOptions{})
//	if err != nil && errors.IsNotFound(err) {
//		_, err := kubeclient.K8sClient.CoreV1().Secrets(ops.namesapce).Create(&corev1.Secret{
//			TypeMeta: v1.TypeMeta{},
//			ObjectMeta: v1.ObjectMeta{
//				Name:        s,
//				Annotations: annotations,
//			},
//			Data:       nil,
//			StringData: basics,
//			Type:       corev1.SecretTypeBasicAuth,
//		})
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//	//todo: if find it. should  update it.
//	return err
//}
//
//func (ops *Ops) checkIfCreateServiceAccount() error {
//	name := "build-bot"
//	_, err := kubeclient.K8sClient.CoreV1().ServiceAccounts(ops.namesapce).Get(name, v1.GetOptions{})
//	if err != nil && errors.IsNotFound(err) {
//		_, err := kubeclient.K8sClient.CoreV1().ServiceAccounts(ops.namesapce).Create(&corev1.ServiceAccount{
//			TypeMeta: v1.TypeMeta{},
//			ObjectMeta: v1.ObjectMeta{
//				Name: name,
//			},
//			Secrets: []corev1.ObjectReference{{
//				Name: "basic-user-pass",
//			}},
//		})
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//	//todo: if find it. should  update it.
//	return err
//}
//
//func (ops *Ops) checkIfPipelineResource(req *ops.BuildRequest,
//	resourceType v1alpha1.PipelineResourceType,
//	typeName string) error {
//	var resourceParam v1alpha1.ResourceParam
//	var resourceParams []v1alpha1.ResourceParam
//	if resourceType == v1alpha1.PipelineResourceTypeGit {
//		resourceParam.Name = "dev"
//		resourceParam.Value = req.Branchname
//		resourceParams = append(resourceParams, resourceParam)
//		resourceParam.Name = "url"
//		resourceParam.Value = req.Reponame
//		resourceParams = append(resourceParams, resourceParam)
//	} else {
//		resourceParam.Name = "url"
//		gitUrl := strings.Split(req.Reponame, "/")
//		repoName := strings.Split(gitUrl[4], ".")
//		resourceParam.Value = ops.registerUrl + ops.deptName + "/" + repoName[0]
//		resourceParams = append(resourceParams, resourceParam)
//	}
//
//	oldGitPipelineResource, err := ops.getPipelineResources(typeName, ops.namesapce)
//	if err != nil {
//		//todo:need to create secretparam by kube  client
//		if errors.IsNotFound(err) {
//			_, err := ops.client.PipelineResources(ops.namesapce).Create(toPipelineResource(resourceParams, resourceType, typeName))
//			if err != nil {
//				return err
//			}
//		} else {
//			return err
//		}
//
//	}
//
//	if oldGitPipelineResource.Spec.DeepCopy().Type != "" {
//		newGitPipelineResource := toPipelineResource(resourceParams, v1alpha1.PipelineResourceTypeGit, typeName)
//		if !reflect.DeepEqual(oldGitPipelineResource.Spec, newGitPipelineResource.Spec) {
//			oldGitPipelineResource.Spec = newGitPipelineResource.Spec
//			_, err = ops.client.PipelineResources(ops.namesapce).Update(oldGitPipelineResource)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
//
//func (ops *Ops) checkIfCreateTask(gitName, imageName string) error {
//	ops.taskName = ops.preFixName + "task"
//	taskInputResource := toTaskResource(gitName, v1alpha1.PipelineResourceTypeGit)
//	taskPathToDockerFileInputParamSpec := v1alpha1.ParamSpec{
//		Name:        "pathToDockerFile",
//		Type:        "string",
//		Description: "The path to the dockerfile to build",
//		Default: &v1alpha1.ArrayOrString{
//			Type:      v1alpha1.ParamTypeString,
//			StringVal: "/workspace/docker-source/Dockerfile",
//			ArrayVal:  nil,
//		},
//	}
//
//	taskPathToContextInputParamSpec := v1alpha1.ParamSpec{
//		Name:        "pathToContext",
//		Type:        "string",
//		Description: "The build context used by Kaniko",
//		Default: &v1alpha1.ArrayOrString{
//			Type:      v1alpha1.ParamTypeString,
//			StringVal: "/workspace/docker-source",
//			ArrayVal:  nil,
//		},
//	}
//	var paramSpecs []v1alpha1.ParamSpec
//	paramSpecs = append(paramSpecs, taskPathToDockerFileInputParamSpec)
//	paramSpecs = append(paramSpecs, taskPathToContextInputParamSpec)
//
//	taskOutPutResource := toTaskResource(imageName, v1alpha1.PipelineResourceTypeImage)
//	env := corev1.EnvVar{
//		Name:      "DOCKER_CONFIG",
//		Value:     "/builder/home/.docker/",
//		ValueFrom: nil,
//	}
//	//withlin-cloud-master-git
//	dockerFile := fmt.Sprintf("--dockerfile=/workspace/%s/Dockerfile", gitName)
//	destination := fmt.Sprintf("--destination=$(outputs.resources.%s.url)", imageName)
//	ct := fmt.Sprintf("--context=/workspace/%s/", gitName)
//	setupName := "build-and-push"
//	buildImage := `gcr.io/kaniko-project/executor:latest`
//	excuteCommand := "/kaniko/executor"
//
//	buildStep := v1alpha1.Step{
//		Container: corev1.Container{
//			Name:    setupName,
//			Image:   buildImage,
//			Command: []string{excuteCommand},
//			Args:    []string{dockerFile, destination, ct, "--insecure=true"},
//			Env:     []corev1.EnvVar{env},
//		},
//	}
//
//	oldTask, err := ops.client.Tasks(ops.namesapce).Get(ops.taskName, v1.GetOptions{})
//	if err != nil {
//		if errors.IsNotFound(err) {
//			_, err := ops.client.Tasks(ops.namesapce).Create(toTask(taskInputResource,
//				paramSpecs, taskOutPutResource, buildStep, ops.taskName))
//			if err != nil {
//				return err
//			}
//		} else {
//			return err
//		}
//	}
//
//	if oldTask.Name != "" {
//		newTask := toTask(taskInputResource,
//			paramSpecs, taskOutPutResource, buildStep, ops.taskName)
//
//		if reflect.DeepEqual(oldTask.Spec, newTask.Spec) {
//			oldTask.Spec = newTask.Spec
//			_, err = ops.client.Tasks(ops.namesapce).Update(oldTask)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
//
//func (ops *Ops) checkIfCreateTaskRun(gitName, imageName string) error {
//	taskRunName := ops.preFixName + "taskrun"
//	//todo: need to  delete oldTaskRun
//	taskResourceBindingName := gitName
//	taskResourceBinding := v1alpha1.TaskResourceBinding{
//		PipelineResourceBinding: v1alpha1.PipelineResourceBinding{
//			Name: taskResourceBindingName,
//			// change by dxp 20200114
//			ResourceRef: &v1alpha1.PipelineResourceRef{
//				Name:       gitName,
//				APIVersion: "",
//			},
//			ResourceSpec: nil,
//		},
//		Paths: nil,
//	}
//	var ps []v1alpha1.Param
//	var pm v1alpha1.Param
//	pm.Name = "pathToDockerFile"
//	pm.Value = v1alpha1.ArrayOrString{
//		Type:      v1alpha1.ParamTypeString,
//		StringVal: fmt.Sprintf("/workspace/%s/Dockerfile", gitName),
//		ArrayVal:  nil,
//	}
//	ps = append(ps, pm)
//	pm.Name = "pathToContext"
//	pm.Value = v1alpha1.ArrayOrString{
//		Type: v1alpha1.ParamTypeString,
//		//todo:lets user chose build dir to build,which a project will mutilple cmd dir.
//		StringVal: fmt.Sprintf("/workspace/%s/", gitName),
//		ArrayVal:  nil,
//	}
//	ps = append(ps, pm)
//	pipelineResourceBinding := "builtImage"
//	oldTaskRun, err := ops.client.TaskRuns(ops.namesapce).Get(taskRunName, v1.GetOptions{})
//	if err != nil {
//		if errors.IsNotFound(err) {
//			taskRun := toTaskRun(taskResourceBinding,
//				ps, pipelineResourceBinding, imageName, taskRunName, ops.taskName)
//			_, err = ops.client.TaskRuns(ops.namesapce).Create(taskRun)
//			if err != nil {
//				return err
//			}
//			return nil
//		} else {
//			return err
//		}
//	}
//
//	err = ops.client.TaskRuns(ops.namesapce).Delete(oldTaskRun.Name, nil)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func(ops *Ops)createPipelineResource() error{
//
//	return nil
//}
//
////type ResourceType  = string
////
////
////const(
////	Git ResourceType = "git"
////	Image ResourceType = "image"
////)
////
////type Param struct{
////	Name  string `json:"name"`
////	Value string `json:"value"`
////}
////
////type Resource struct{
////	ResourceType ResourceType
////	params []Param
////}
////
////
//
//
//
//
////toPipelineResource map Pipeline  enity
//func toPipeline(name string,
//	resources []v1alpha1.PipelineDeclaredResource,
//	tasks []v1alpha1.PipelineTask,
//	paramSpecs []v1alpha1.ParamSpec) *v1alpha1.Pipeline {
//	return &v1alpha1.Pipeline{
//		TypeMeta:   v1.TypeMeta{
//			Kind:       "",
//			APIVersion: "",
//		},
//		ObjectMeta: v1.ObjectMeta{
//			Name:                       name,
//		},
//		Spec:       v1alpha1.PipelineSpec{
//			Resources: resources,
//			Tasks:     tasks,
//			Params:    paramSpecs,
//		},
//		Status:     v1alpha1.PipelineStatus{},
//	}
//}
//
////func (ops *Ops) checkIfCreatePipeline()  error {
////	pipelineName := ops.preFixName + "pipeline"
////	var resources []v1alpha1.PipelineDeclaredResource
////	var tasks []v1alpha1.PipelineTask
////	var paramSpecs []v1alpha1.ParamSpec
////	taskInputResource := toPipeline(pipelineName,resources,tasks,paramSpecs)
////
////
////	return nil
////}

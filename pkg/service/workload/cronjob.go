package workload

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type CronJob struct {
	defaultImplWorkloadsResourceHandler
}

func NewCronJob() *CronJob {
	return &CronJob{defaultImplWorkloadsResourceHandler{}}
}

func (c *CronJob) GetOne(namespace, name string) map[string]string {

	return nil
}

func (c *CronJob) List(res, pos int64, size int64, flag string) (list *unstructured.UnstructuredList, err error) {
	return nil, nil
}

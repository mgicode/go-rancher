package client

const (
	RESTART_POLICY_TYPE = "restartPolicy"
)

type RestartPolicy struct {
	Resource `yaml:"-"`

	MaximumRetryCount int64 `json:"maximumRetryCount,omitempty" yaml:"maximum_retry_count,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type RestartPolicyCollection struct {
	Collection
	Data   []RestartPolicy `json:"data,omitempty"`
	client *RestartPolicyClient
}

type RestartPolicyClient struct {
	rancherClient *RancherClient
}

type RestartPolicyOperations interface {
	List(opts *ListOpts) (*RestartPolicyCollection, error)
	Create(opts *RestartPolicy) (*RestartPolicy, error)
	Update(existing *RestartPolicy, updates interface{}) (*RestartPolicy, error)
	ById(id string) (*RestartPolicy, error)
	Delete(container *RestartPolicy) error
}

func newRestartPolicyClient(rancherClient *RancherClient) *RestartPolicyClient {
	return &RestartPolicyClient{
		rancherClient: rancherClient,
	}
}

func (c *RestartPolicyClient) Create(container *RestartPolicy) (*RestartPolicy, error) {
	resp := &RestartPolicy{}
	err := c.rancherClient.doCreate(RESTART_POLICY_TYPE, container, resp)
	return resp, err
}

func (c *RestartPolicyClient) Update(existing *RestartPolicy, updates interface{}) (*RestartPolicy, error) {
	resp := &RestartPolicy{}
	err := c.rancherClient.doUpdate(RESTART_POLICY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RestartPolicyClient) List(opts *ListOpts) (*RestartPolicyCollection, error) {
	resp := &RestartPolicyCollection{}
	err := c.rancherClient.doList(RESTART_POLICY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RestartPolicyCollection) Next() (*RestartPolicyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RestartPolicyCollection{}
		err := cc.client.rancherClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RestartPolicyClient) ById(id string) (*RestartPolicy, error) {
	resp := &RestartPolicy{}
	err := c.rancherClient.doById(RESTART_POLICY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RestartPolicyClient) Delete(container *RestartPolicy) error {
	return c.rancherClient.doResourceDelete(RESTART_POLICY_TYPE, &container.Resource)
}

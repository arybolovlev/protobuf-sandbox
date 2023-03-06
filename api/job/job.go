package job

type Job struct {
	ID   string `hcl:"id,optional"`
	Name string `hcl:"name,label"`
}

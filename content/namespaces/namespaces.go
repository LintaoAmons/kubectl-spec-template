package namespaces

const (
	Ref = "https://kubernetes.io/docs/tasks/administer-cluster/namespaces/#creating-a-new-namespace"
	Simple = `
apiVersion: v1
kind: Namespace
metadata:
  name: <insert-namespace-name-here>
	`
)

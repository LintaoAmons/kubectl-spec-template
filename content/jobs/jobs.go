package jobs

const (
	Ref = "https://kubernetes.io/docs/concepts/workloads/controllers/job/"
	Simple = `
apiVersion: batch/v1
kind: Job
metadata:
  name: pi
spec:
  template:
    spec:
      containers:
      - name: pi
        image: perl:5.34.0
        command: ["perl",  "-Mbignum=bpi", "-wle", "print bpi(2000)"]
      restartPolicy: Never
  backoffLimit: 4 # the number of retries that Kubernetes will attempt in case the job fails. 
`
)

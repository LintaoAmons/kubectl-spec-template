package persistentvolumeclaims

const (
	Ref = "https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/"

	Simply = `
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: task-pv-claim
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
`
)

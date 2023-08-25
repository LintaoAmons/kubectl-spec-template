package persistentvolumes

const (
	Simple = `
apiVersion: v1
kind: PersistentVolume
metadata:
  name: task-pv-volume
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/mnt/data"
	`
	
	StaticBinding = `
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  volumeName: my-pv     # Name of the PersistentVolume to bind
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
	`
)



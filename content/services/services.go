package services

const (
	Ref = "https://kubernetes.io/docs/concepts/services-networking/service/"
	Simple = `
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app.kubernetes.io/name: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
	`
)

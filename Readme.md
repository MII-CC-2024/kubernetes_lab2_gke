# K8S - GCP Kubernetes Engine (GKE)

En GKE, Google administra el plano de control y los componentes del sistema; en el modo Autopilot, que es la forma recomendada, Google también administra los nodos trabajadores.

## Requisitos

### Instala gcloud CLI

https://cloud.google.com/sdk/docs/install?hl=es-419

Para Linux: https://cloud.google.com/sdk/docs/install?hl=es-419#linux

Identificaté en la consola con:

```shell
gcloud init
```
Con el comando anterior puedes hacer login, pero si no lo haces, utiliza el comando:

```shell
gcloud auth login
```

#### NOTA
Si gcloud no encuentra Python define la variable de entorno CLOUDSDK_PYTHON apuntando al fichero python.exe con la ruta al directorio de instalación de python (por ejemplo, C:\Program Files\Python310\python.exe).

### Instala kubectl

```shell
gcloud components install kubectl
```

O utiliza alguno de los otros métodos disponibles:

https://kubernetes.io/docs/tasks/tools/


### Instala el componente gke-gcloud-auth-plugin

```shell
gcloud components install gke-gcloud-auth-plugin
```
(Este comando requiere privilegios de administrador)


## Crea el cluster en GKE 

Utiliza la consola Web de GCP para crear el cluster: 
- Selecciona Kubernetes Engine
- |+| Create
- Selecciona el nombre y la región
- Deja el resto de opciones por defecto y pulsa en Create

Para obtener las credenciales para conexión al cluster, accede al cluster haciendo click en su nombre y pulsa "Establecer Conexión", copia y ejecuta el comando en una consola domde estés identificado en GCP. 

```
$ gcloud container clusters get-credentials <nombre_cluster> --region <region> --project <proyecto>
Fetching cluster endpoint and auth data.
kubeconfig entry generated for <nombre_cluster>.

```
Este comando creará o, si ya existe, actualizará el fichero .kube/config para incluir las credenciales de acceso al clúster, estableciéndolo como el contexto por defecto.

Podemos probar si tenemos las credenciales y la conexión es correcta mediante el comando:

```
$ kubectl cluster-info 
Kubernetes control plane is running at https://<IP-CLUSTER>
GLBCDefaultBackend is running at https://<IP-CLUSTER>/api/v1/namespaces/kube-system/services/default-http-backend:http/proxy        
KubeDNS is running at https://<IP-CLUSTER>/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://<IP-CLUSTER>/api/v1/namespaces/kube-system/services/https:metrics-server:/proxy
```
Además, podemos obtener información sobre las versiones del clientes y el servidor con el comando:

```
$ kubectl version -o=yaml  
clientVersion:
  buildDate: "2024-02-14T10:40:49Z"
  compiler: gc
  gitCommit: 4b8e819355d791d96b7e9d9efe4cbafae2311c88
  gitTreeState: clean
  gitVersion: v1.29.2
  goVersion: go1.21.7
  major: "1"
  minor: "29"
  platform: windows/amd64
kustomizeVersion: v5.0.4-0.20230601165947-6ce0bf390ce3
serverVersion:
  buildDate: "2024-02-26T09:16:36Z"
  compiler: gc
  gitCommit: 9a6e4cf95a19739cf5c154bdb6ed9cca39051829
  gitTreeState: clean
  gitVersion: v1.28.7-gke.1026000
  goVersion: go1.21.7 X:boringcrypto
  major: "1"
  minor: "28"
  platform: linux/amd64
```
Si todo es correcto, veremos la versión del cliente y del servidor; si no tenemos las credenciales solo veremos la versión del cliente y un error de conexión al servidor.


Para más información sobre las principales opciones del comando kubectl consulta alguno de estos enlaces:

```
https://kubernetes.io/docs/reference/kubectl/cheatsheet/

https://jamesdefabia.github.io/docs/user-guide/kubectl/kubectl/

```
Por ejemplo, los siguientes comandos, permiten ver la configuración del fichero .kube/config, obtener los contextos (clusters) disponibles, ver el contexto actual y cambiar el contexto por defecto 

```
kubectl config view                         # Muestra fichero kube/config
kubectl config get-contexts                 # Muestra los contextos (clusters) disponibles
kubectl config current-context              # Muestra el contexto actual 
kubectl config use-context <cluster-name>   # Asigna el contexto por defecto a <cluster-name>
``` 
### Más información sobre el cluster

Versiones de las API de kubernetes disponibles:

```
$ kubectl api-versions
admissionregistration.k8s.io/v1
apiextensions.k8s.io/v1
apiregistration.k8s.io/v1
apps/v1
authentication.k8s.io/v1
authorization.k8s.io/v1
auto.gke.io/v1
auto.gke.io/v1alpha1
autoscaling.gke.io/v1beta1
autoscaling.k8s.io/v1
autoscaling.k8s.io/v1beta2
autoscaling/v1
autoscaling/v2
batch/v1
certificates.k8s.io/v1
cilium.io/v2
cilium.io/v2alpha1
cloud.google.com/v1
cloud.google.com/v1beta1
coordination.k8s.io/v1
discovery.k8s.io/v1
events.k8s.io/v1
flowcontrol.apiserver.k8s.io/v1beta2
flowcontrol.apiserver.k8s.io/v1beta3
gateway.networking.k8s.io/v1alpha2
gateway.networking.k8s.io/v1beta1
ha.gke.io/v1
hub.gke.io/v1
internal.autoscaling.gke.io/v1
monitoring.googleapis.com/v1
monitoring.googleapis.com/v1alpha1
networking.gke.io/v1
networking.gke.io/v1alpha1
networking.gke.io/v1beta1
networking.gke.io/v1beta2
networking.k8s.io/v1
node.k8s.io/v1
nodemanagement.gke.io/v1alpha1
policy/v1
rbac.authorization.k8s.io/v1
scheduling.k8s.io/v1
snapshot.storage.k8s.io/v1
snapshot.storage.k8s.io/v1beta1
storage.k8s.io/v1
v1
vulnerabilities.protect.gke.io/v1
warden.gke.io/v1
```

Recursos disponibles para crear en el clúster:

```
$ kubectl api-resources
NAME                              SHORTNAMES          APIVERSION                             NAMESPACED   KIND
bindings                                              v1                                     true         Binding
componentstatuses                 cs                  v1                                     false        ComponentStatus
configmaps                        cm                  v1                                     true         ConfigMap
endpoints                         ep                  v1                                     true         Endpoints
events                            ev                  v1                                     true         Event
limitranges                       limits              v1                                     true         LimitRange
namespaces                        ns                  v1                                     false        Namespace
nodes                             no                  v1                                     false        Node
persistentvolumeclaims            pvc                 v1                                     true         PersistentVolumeClaim     
persistentvolumes                 pv                  v1                                     false        PersistentVolume
pods                              po                  v1                                     true         Pod
podtemplates                                          v1                                     true         PodTemplate
replicationcontrollers            rc                  v1                                     true         ReplicationController     
resourcequotas                    quota               v1                                     true         ResourceQuota
secrets                                               v1                                     true         Secret
serviceaccounts                   sa                  v1                                     true         ServiceAccount
services                          svc                 v1                                     true         Service
mutatingwebhookconfigurations                         admissionregistration.k8s.io/v1        false        MutatingWebhookConfiguration
validatingwebhookconfigurations                       admissionregistration.k8s.io/v1        false        ValidatingWebhookConfiguration
customresourcedefinitions         crd,crds            apiextensions.k8s.io/v1                false        CustomResourceDefinition  
apiservices                                           apiregistration.k8s.io/v1              false        APIService
controllerrevisions                                   apps/v1                                true         ControllerRevision
daemonsets                        ds                  apps/v1                                true         DaemonSet
deployments                       deploy              apps/v1                                true         Deployment
replicasets                       rs                  apps/v1                                true         ReplicaSet
statefulsets                      sts                 apps/v1                                true         StatefulSet
selfsubjectreviews                                    authentication.k8s.io/v1               false        SelfSubjectReview
tokenreviews                                          authentication.k8s.io/v1               false        TokenReview
localsubjectaccessreviews                             authorization.k8s.io/v1                true         LocalSubjectAccessReview  
selfsubjectaccessreviews                              authorization.k8s.io/v1                false        SelfSubjectAccessReview   
selfsubjectrulesreviews                               authorization.k8s.io/v1                false        SelfSubjectRulesReview    
subjectaccessreviews                                  authorization.k8s.io/v1                false        SubjectAccessReview       
allowlistedv2workloads                                auto.gke.io/v1                         false        AllowlistedV2Workload     
allowlistedworkloads                                  auto.gke.io/v1                         false        AllowlistedWorkload       
horizontalpodautoscalers          hpa                 autoscaling/v2                         true         HorizontalPodAutoscaler   
multidimpodautoscalers            mpa                 autoscaling.gke.io/v1beta1             true         MultidimPodAutoscaler     
verticalpodautoscalers            vpa                 autoscaling.k8s.io/v1                  true         VerticalPodAutoscaler     
cronjobs                          cj                  batch/v1                               true         CronJob
jobs                                                  batch/v1                               true         Job
certificatesigningrequests        csr                 certificates.k8s.io/v1                 false        CertificateSigningRequest 
ciliumendpoints                   cep,ciliumep        cilium.io/v2                           true         CiliumEndpoint
ciliumendpointslices              ces                 cilium.io/v2alpha1                     false        CiliumEndpointSlice       
ciliumexternalworkloads           cew                 cilium.io/v2                           false        CiliumExternalWorkload    
ciliumidentities                  ciliumid            cilium.io/v2                           false        CiliumIdentity
ciliumlocalredirectpolicies       clrp                cilium.io/v2                           true         CiliumLocalRedirectPolicy 
ciliumnodes                       cn,ciliumn          cilium.io/v2                           false        CiliumNode
backendconfigs                    bc                  cloud.google.com/v1                    true         BackendConfig
leases                                                coordination.k8s.io/v1                 true         Lease
endpointslices                                        discovery.k8s.io/v1                    true         EndpointSlice
events                            ev                  events.k8s.io/v1                       true         Event
flowschemas                                           flowcontrol.apiserver.k8s.io/v1beta3   false        FlowSchema
prioritylevelconfigurations                           flowcontrol.apiserver.k8s.io/v1beta3   false        PriorityLevelConfiguration
gatewayclasses                    gc                  gateway.networking.k8s.io/v1beta1      false        GatewayClass
gateways                          gtw                 gateway.networking.k8s.io/v1beta1      true         Gateway
httproutes                                            gateway.networking.k8s.io/v1beta1      true         HTTPRoute
referencegrants                   refgrant            gateway.networking.k8s.io/v1beta1      true         ReferenceGrant
highavailabilityapplications                          ha.gke.io/v1                           true         HighAvailabilityApplication
memberships                                           hub.gke.io/v1                          false        Membership
capacityrequests                  capreq              internal.autoscaling.gke.io/v1         true         CapacityRequest
clusterpodmonitorings                                 monitoring.googleapis.com/v1           false        ClusterPodMonitoring      
clusterrules                                          monitoring.googleapis.com/v1           false        ClusterRules
globalrules                                           monitoring.googleapis.com/v1           false        GlobalRules
operatorconfigs                                       monitoring.googleapis.com/v1           true         OperatorConfig
podmonitorings                                        monitoring.googleapis.com/v1           true         PodMonitoring
rules                                                 monitoring.googleapis.com/v1           true         Rules
frontendconfigs                                       networking.gke.io/v1beta1              true         FrontendConfig
gcpbackendpolicies                                    networking.gke.io/v1                   true         GCPBackendPolicy
gcpgatewaypolicies                                    networking.gke.io/v1                   true         GCPGatewayPolicy
gkenetworkparamsets                                   networking.gke.io/v1                   false        GKENetworkParamSet        
healthcheckpolicies                                   networking.gke.io/v1                   true         HealthCheckPolicy
lbpolicies                                            networking.gke.io/v1                   true         LBPolicy
managedcertificates               mcrt                networking.gke.io/v1                   true         ManagedCertificate        
networkloggings                   nl                  networking.gke.io/v1alpha1             false        NetworkLogging
networks                                              networking.gke.io/v1                   false        Network
redirectservices                  rds                 networking.gke.io/v1alpha1             true         RedirectService
serviceattachments                                    networking.gke.io/v1                   true         ServiceAttachment
servicenetworkendpointgroups      svcneg              networking.gke.io/v1beta1              true         ServiceNetworkEndpointGroup
ingressclasses                                        networking.k8s.io/v1                   false        IngressClass
ingresses                         ing                 networking.k8s.io/v1                   true         Ingress
networkpolicies                   netpol              networking.k8s.io/v1                   true         NetworkPolicy
runtimeclasses                                        node.k8s.io/v1                         false        RuntimeClass
updateinfos                       updinf              nodemanagement.gke.io/v1alpha1         true         UpdateInfo
poddisruptionbudgets              pdb                 policy/v1                              true         PodDisruptionBudget       
clusterrolebindings                                   rbac.authorization.k8s.io/v1           false        ClusterRoleBinding        
clusterroles                                          rbac.authorization.k8s.io/v1           false        ClusterRole
rolebindings                                          rbac.authorization.k8s.io/v1           true         RoleBinding
roles                                                 rbac.authorization.k8s.io/v1           true         Role
priorityclasses                   pc                  scheduling.k8s.io/v1                   false        PriorityClass
volumesnapshotclasses             vsclass,vsclasses   snapshot.storage.k8s.io/v1             false        VolumeSnapshotClass       
volumesnapshotcontents            vsc,vscs            snapshot.storage.k8s.io/v1             false        VolumeSnapshotContent     
volumesnapshots                   vs                  snapshot.storage.k8s.io/v1             true         VolumeSnapshot
csidrivers                                            storage.k8s.io/v1                      false        CSIDriver
csinodes                                              storage.k8s.io/v1                      false        CSINode
csistoragecapacities                                  storage.k8s.io/v1                      true         CSIStorageCapacity        
storageclasses                    sc                  storage.k8s.io/v1                      false        StorageClass
volumeattachments                                     storage.k8s.io/v1                      false        VolumeAttachment
extractionresults                 er                  vulnerabilities.protect.gke.io/v1      true         ExtractionResult
audits                                                warden.gke.io/v1                       false        Audit
```

Podemos solicitar información sobre un determinado recurso o de alguna de sus opciones:

```
$ kubectl explain --api-version=apps/v1 Deployment

$ kubectl explain --api-version=apps/v1 Deployment.metadata | more
$ kubectl explain --api-version=apps/v1 Deployment.spec.template | more
$ kubectl explain --api-version=apps/v1 Deployment.spec.template.spec.containers | more
```

## Consultar recursos del cluster

Podemos consutar los recursos del Clúster con: kubectl get

Por ejemplo, los espacios de nombre (ns: namespaces) con 

```
$ kubectl get ns   
NAME                       STATUS   AGE
default                    Active   66m
gke-gmp-system             Active   64m
gke-managed-cim            Active   65m
gke-managed-filestorecsi   Active   65m
gke-managed-system         Active   65m
gmp-public                 Active   64m
kube-node-lease            Active   66m
kube-public                Active   66m
kube-system                Active   66m
```
El espacio de nombre por defecto es default, todos los comandos se ejecutan sobre ese espacio de nombre a menos que se indique lo contrario. Así, para obtener los pods del espacio de nombre por defecto:

```
$ kubectl get pods      
No resources found in default namespace.
```
Podemos especifiar el espacio de nombre con: --namespace <espacio_nombres>

```
$ kubectl get pods --namespace kube-system 
NAME                                                       READY   STATUS    RESTARTS   AGE
antrea-controller-horizontal-autoscaler-6dfb5849b6-5dxf2   0/1     Pending   0          68m
egress-nat-controller-54bddbd686-dg8q7                     0/1     Pending   0          68m
event-exporter-gke-7d996c57bf-jf7nr                        0/2     Pending   0          68m
konnectivity-agent-6c74695575-q75bp                        0/2     Pending   0          68m
konnectivity-agent-autoscaler-5847cf65c7-2knz5             0/1     Pending   0          68m
kube-dns-6f955b858b-dvz96                                  0/4     Pending   0          69m
kube-dns-autoscaler-755c7dfdf5-sl668                       0/1     Pending   0          69m
l7-default-backend-6779bb6c8d-dv549                        0/1     Pending   0          68m
metrics-server-v0.6.3-7cb4458849-m5mwm                     0/2     Pending   0          68m
```

## Crear el primer Pod

De forma imperativa, podemos crear un Pod con:

```
$ kubectl run nginx --image=nginx --port=80
pod/nginx created
```

También, podemos usar un fichero de manifiesto .yaml (por ejemplo, nginx-pod.yaml) con la configuración del Pod:

```
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    ports:
    - containerPort: 80
```
y ejecutar, el comando imperativo:

```
$ kubectl create -f nginx-pod.yaml
pod/nginx created
```

O podemos ejecutarlo de forma declarativa con:

```
$ kubectl apply -f nginx-pod.yaml  
pod/nginx created
```

Ver los Pods:

```
$ kubectl get pods -o=wide
NAME    READY   STATUS    RESTARTS   AGE   IP          NODE             NOMINATED NODE   READINESS GATES
nginx   1/1     Running   0          29s   10.1.0.29   nodo1            <none>           <none>
```

Ver los logs de un Pod con kubectl logs <nombre_pod>

```
$ kubectl logs nginx
...
```

Ejecutar un comando en un Pod con kubectl exec <nombre_pod> -- <comando>  

```
$ kubectl exec -it nginx -- bash
root@nginx:/#

```

Obtener la descripción de un recurso con kubectl describe, en este caso de un Pod con:

```
kubectl describe pod nginx
...
```

Borrar recursos con kubectl delete

```
kubectl delete pod nginx
```
También podemos hacerlo usando el fichero yaml:

```
kubectl delete -f nginx-pod.yaml
```

## Exponer un Pod mediante un servicio de forma imperativa

### Servicio ClusterIP

```
$ kubectl expose pod nginx --port=80 --name=frontend
```

Ver los servicios (services):

```
$ kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
frontend     ClusterIP   10.100.18.125   <none>        80/TCP    2s
...
```

Ver los endpoints, a los que apunta un servicio:

```
kubectl get ep 
NAME         ENDPOINTS           AGE
frontend     10.1.0.30:80        103s
...
```

Ahora tenemos accesible el Pod (con IP 10.1.0.30) desde la IP del servicio (10.100.18.125).

Podemos entrar en cualquier Pod y comprobar que el servicio accede al Pod nginx.


### Servicio NodePort

```
$ kubectl expose pod nginx --port=80 --type=NodePort --name=frontend
```

Ver el servicio

```
$ kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
frontend     NodePort    10.108.59.115   <none>        80:31930/TCP   29s
```

Ahora el Pod es accesible desde la IP del Host mediante el puerto 31930.


```
$ ipconfig
Dirección IPv4. . . . . . . . . . . . . . : 192.168.18.34

$ curl 192.168.18.34:31930

```

### Servicio LoadBalancer

```
$ kubectl expose pod nginx --port=80 --type=LoadBalancer --name=frontend
NAME               TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
service/frontend   LoadBalancer   10.104.13.74   localhost     80:32523/TCP   2m6s
```
Ahora el Pod es accesible mediante la IP externa (EXTERNAL-IP) asignada al balanceador.

### Mostrar recursos

Todos los recursos (all):

```
kubectl get all    
NAME        READY   STATUS    RESTARTS   AGE
pod/nginx   1/1     Running   0          136m

NAME               TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
service/frontend   LoadBalancer   10.104.13.74   localhost     80:32523/TCP   2m6s

```

Por etiqueta (-l), en este caso run=nginx:

```
kubectl get all -l run=nginx     
NAME        READY   STATUS    RESTARTS   AGE
pod/nginx   1/1     Running   0          136m

NAME               TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
service/frontend   LoadBalancer   10.104.13.74   localhost     80:32523/TCP   2m6s

```

### Eliminar recursos por etiquetas

```
$ kubectl delete all -l run=nginx 
pod "nginx" deleted
service "frontend" deleted
```

## Implementar Aplicación con Go

https://go.dev/doc/tutorial/getting-started

Crear la carpeta app para el proyecto y habilitar las dependencias:

```
$ mkdir app
$ cd app

$ go mod init example/hello
```

Crear un fichero hello.go con el siguiente contenido:

```go

package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func handler(w http.ResponseWriter, r *http.Request) {

        hostname, err := os.Hostname()
        if err != nil {
                fmt.Fprintln(w, "Error:", err)
        } else {
                fmt.Fprintln(w, "Hello from", hostname, "(version 2023)")
        }
}

func main() {
        http.HandleFunc("/", handler)
        log.Println("Go Hello is listening on port 8888")
        log.Fatal(http.ListenAndServe(":8080", nil))
}

```

Comprobemos el funcionamiento en local:

```
$ go run app\hello.go
2024/05/16 09:15:38 Go Hello is listening on port 8080
```


## Contenerizar la App

Dockerfile

```dockerfile
FROM golang:1.22.3-alpine AS build
WORKDIR /src/
RUN go mod init example/hello
COPY app/main.go /src/
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /bin/hello

FROM scratch
EXPOSE 8080
COPY --from=build /bin/hello /bin/hello
ENTRYPOINT ["/bin/hello"]
```

Construir la imagen:

```shell 
docker build -t jluisalvarez/go_hello:2023 .
```

Comprobar funcionamiento de la imagen:

```shell 
docker run -d -p 8080:8080 --name hello_go jluisalvarez/go_hello:2023   
```

## Crear Deployment en Kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-deployment
  labels:
    app: hello
spec:
  replicas: 3
  selector:
    matchLabels:
      app: hello
  template:
    metadata:
      labels:
        app: hello
    spec:
      containers:
      - name: go-hello
        image: jluisalvarez/go_hello:2023
        ports:
        - containerPort: 8080
        resources:
          limits:
            cpu: 500m
          requests:
            cpu: 200m
```

Despliega en Kubernetes:

```shell
kubectl apply -f k8s/hello_deploy.yaml
```

## Crear servicio

Crear el fichero hello_service.yaml para exponer el deployment en un servicio tipo load balancer.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: hello-service
spec:
  selector:
    app: hello
  ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
  type: LoadBalancer
```

Aplica el manifiesto con:

```shell
$ kubectl apply -f k8s/hello_service.yaml
```

## Escalar la aplicación

Podemos escalar la aplicación de forma manual con:

A patir del fichero

```
kubectl scale --replicas=5 -f hello_deploy.yaml
```

El recurso concreto:

```
kubectl scale --replicas=3 deployment/hello-deployment
```

### Autoescalado

También, podemos definir un escalado automático en función de parámetros como el uso de la CPU. Para ello,
debemos tener disponibles en nuestro clúster las métricas. 

```
kubectl autoscale deployment hello-deployment --max=3 --min=1 --cpu-percent=50
```

#### NOTA

Simula peticiones en Linux:

```
for ((i=1;i<=10000000;i++)); do   curl --header "Connection: keep-alive" "http://localhost:8080/"; done
```
Simula peticiones en Windows:

```
@echo off
for /L %%i in (1,1,10000000) do (
curl --header "Connection: keep-alive" "http://localhost:8080/"
)
```

## Desplegar una nueva versión

En kubernetes un Deployment nos permite definir la estrategia para publicar las actualizaciones, lo que se conoce como rollout. En la estrategia por defecto, RollingUpdate, la actualización sigue un enfoque gradual.

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-deployment
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 2
      maxUnavailable: 0
  selector:
    ...
  ```

Etiqueta la versión actual

```
kubectl annotate deployment/hello-deployment kubernetes.io/change-cause="version 2023"
```

Provoca un cambio de imagen, imperativo: 

```
kubectl set image deployments/hello-deployment go-hello=jluisalvarez/go_hello:2024
```

(Esto también, podría haberse realizado modificando el fichero yaml y aplicando o editando directamente el deployment)

Ver versiones
```
kubectl rollout history deployment/hello-deployment
```

Anotaciones del motivo del cambio
```
kubectl annotate deployment/hello-deployment kubernetes.io/change-cause="New version 2024"
```

Comprobar estado de la actualización:
```
kubectl rollout status deployments/hello-deployment
```

Deshacer actualización:
```
kubectl rollout undo deployments/hello-deployment
```
```
kubectl rollout undo --to-revision=2 deployments/hello-deployment
```

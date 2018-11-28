package dockerswarm

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"bytes"
	"bufio"

	"github.com/fsouza/go-dockerclient"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/api/types/mount"
)

type Config struct {
	GameId string
	Name string
	Ports []swarm.PortConfig
	Namespace string 
	Image string 
	DataVol1 string 
	Maxplayers string 
	Version string
	Envs	[]string
}

type Dockerswarm struct {
	Client *docker.Client
}

func NewBackend() Dockerswarm {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		panic(err)
	}
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
	}

	return Dockerswarm{Client: client}
}

// func (d Dockerswarm) getAuthConfigurations() docker.AuthConfigurations{

// 	configs := docker.AuthConfigurations{}

// 	gcloudFile := os.Getenv("OGP_GCLOUD_JSON")

// 	if gcloudFile != "" {
// 		gcloudJson, err := ioutil.ReadFile(gcloudFile)
// 		if err != nil {
// 			fmt.Printf("Error Reading Gcloud file %+v\n",err)
// 		}

// 		configs.Configs["gcloud"] = docker.AuthConfiguration{
// 			Username: "_json_key",
// 			Password: string(gcloudJson),
// 			ServerAddress: "https://gcr.io",
// 		}
// 	}

// 	return configs
// }

func (d Dockerswarm) getAuthConfiguration() docker.AuthConfiguration{

	gcloudFile := os.Getenv("OGP_GCLOUD_JSON")

	if gcloudFile != "" {
		gcloudJson, err := ioutil.ReadFile(gcloudFile)
		if err != nil {
			fmt.Printf("Error Reading Gcloud file %+v\n",err)
		}

		return docker.AuthConfiguration{
			Username: "_json_key",
			Password: string(gcloudJson),
			ServerAddress: "https://gcr.io",
		}
	}

	return docker.AuthConfiguration{}
}

func (d Dockerswarm) getMount(GameId string, path string) mount.Mount {

	fmt.Printf("Getting MOUNTS. The path: %+v\n", path)
	fmt.Printf("Getting MOUNTS. The GameId: %+v\n", GameId)

	storageType := strings.ToUpper(os.Getenv("STORAGE"))

	driverOptions := make(map[string]string)

// TODO, IF NFS, we could delete the volume before mounting it (to ensure all changes are there)

	switch storageType {
	case "NFS":
		fmt.Printf("Getting MOUNTS. Building NFS GameId: %+v\n", GameId)

		driverOptions["type"] = "nfs"
		driverOptions["o"] = "addr="+os.Getenv("STORAGE_NFS_ADDRESS") +",rw"
		driverOptions["device"] = ":"+os.Getenv("STORAGE_NFS_PATH") +"/"+ GameId
	case "LOCAL":
	}

	driver := mount.Driver{
		Name: "local",
		Options: driverOptions,
	}
	volumeOptions :=mount.VolumeOptions{
		NoCopy: true,
		DriverConfig: &driver,
	}
	mountObj := mount.Mount{
		Type: mount.TypeVolume,
		Source: "volume_"+GameId+"_1", // TODO: This hard codes the volume name with a 1 on the end. 
		Target: "/home/steam/linuxgsm/serverfiles/" + path,
		VolumeOptions: &volumeOptions,
	}

	fmt.Printf("Getting MOUNTS. The MountObg: %+v\n", mountObj)

	return mountObj
}

// This function is kinda dirty
func (d Dockerswarm) Start(config Config) {
	fmt.Println("Docker Starting!")

	serviceSpec := swarm.ServiceSpec{}
	updating := false
	serviceId := d.getServiceId(config.Name)
	var version uint64 

	if d.DoesServiceExist(config.Name) {
		updating = true

		serviceDetails, err := d.Client.InspectService(serviceId)

		if err != nil {
			fmt.Printf("Error InspectService %+v\n",err)
			return
		}
		serviceSpec = serviceDetails.Spec
		version = serviceDetails.Meta.Version.Index		
	}

	imagePath := config.Namespace +"/"+config.Image

	mounts := []mount.Mount{
		d.getMount(config.GameId,config.DataVol1),
	}

	containerSpec := swarm.ContainerSpec{
		Image: imagePath,
		Env: config.Envs,
		Mounts: mounts,
	}

	replicas := uint64(1)

	serviceSpec = swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: config.Name,
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: &containerSpec,
		},
		EndpointSpec: &swarm.EndpointSpec{
			Ports: config.Ports,
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: &replicas,
			},
		},
	}

	if updating {
		fmt.Println("Updating existing Docker Service!")

		serviceOpts := docker.UpdateServiceOptions{
			Auth: docker.AuthConfiguration{},
			ServiceSpec: serviceSpec,
			Version: version,
		}
	
		if err := d.Client.UpdateService(serviceId,serviceOpts); err != nil {
			fmt.Printf("Error Updating service: %v\n", err)
		}
	} else {
		fmt.Println("Creating new Docker Service!")
		serviceOpts := docker.CreateServiceOptions{
			Auth: d.getAuthConfiguration(),
			ServiceSpec: serviceSpec,
		}
	
		_, err := d.Client.CreateService(serviceOpts)
		if err != nil {
			fmt.Printf("Error Creating service %+v\n",err)
		}
	}
}

func (d Dockerswarm) Stop(name string) {
	fmt.Println("Docker Stopping!")
	d.scaleService(name,uint64(0))
}

func (d Dockerswarm) scaleService(name string, replicas uint64) {
	serviceId := d.getServiceId(name)

	if serviceId == "" {
		fmt.Println("Service Doesn't exist")
		return
	}

	serviceDetails, err := d.Client.InspectService(serviceId)

	if err != nil {
		fmt.Printf("Error InspectService %+v\n",err)
		return
	}

	serviceSpec := serviceDetails.Spec
	serviceSpec.Mode.Replicated.Replicas = &replicas
	serviceSpec.TaskTemplate.ForceUpdate = uint64(1)
	
	serviceOpts := docker.UpdateServiceOptions{
		Auth: docker.AuthConfiguration{},
		ServiceSpec: serviceSpec,
		Version: serviceDetails.Meta.Version.Index,
	}

	if err := d.Client.UpdateService(serviceId,serviceOpts); err != nil {
		fmt.Printf("Error updating service: %v\n", err)
	}
}

func (d Dockerswarm) IsRunning(name string)  bool {

	serviceId := d.getServiceId(name)

	if serviceId == "" {
		fmt.Println("Service Doesn't exist")
		return false
	}

	serviceDetails, err := d.Client.InspectService(serviceId)

	if err != nil {
		fmt.Printf("Error InspectService %+v\n",err)
		return false
	}

	if *serviceDetails.Spec.Mode.Replicated.Replicas == 0 {
		fmt.Println("Service Not Running")
		return false
	} else {
		fmt.Println("Service Running")
		return true
	}
}

func (d Dockerswarm) DoesServiceExist(name string)  bool {
	serviceId := d.getServiceId(name)
	if serviceId == "" {
		fmt.Println("Service Doesn't exist")
		return false
	}
	return true
}


func (d Dockerswarm) getServiceId(name string) string {
	filter := map[string][]string{}

	filter["name"]=[]string{name}

	opts := docker.ListServicesOptions{
		Filters: filter,
	}

	serviceList, err := d.Client.ListServices(opts)

	if err != nil {
		fmt.Printf("Error Listing services %+v\n",err)
		return ""
	}

	if len(serviceList) == 0 {
		return ""
	}

	return serviceList[0].ID
}

func (d Dockerswarm) Logs(name string) bytes.Buffer {

	serviceId := d.getServiceId(name)

	var logs bytes.Buffer
    writer := bufio.NewWriter(&logs)

	opts := docker.LogsServiceOptions{
		Service: serviceId,
		OutputStream: writer,
		Stdout:     true,
		Stderr:     true,
		Timestamps: true,
		Tail: "100",
	}

	d.Client.GetServiceLogs(opts)

	return logs
}



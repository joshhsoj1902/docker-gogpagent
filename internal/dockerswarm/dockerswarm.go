package dockerswarm

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
	"github.com/docker/docker/api/types/swarm"
)

type Config struct {
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

func (d Dockerswarm) Start(config Config) {
	fmt.Println("Docker Starting!")

	if d.DoesServiceExist(config.Name) {
		d.scaleService(config.Name,uint64(1))
		return
	}

	imagePath := config.Namespace +"/"+config.Image

	containerSpec := swarm.ContainerSpec{
		Image: imagePath,
		Env: config.Envs,
	}

	replicas := uint64(1)

	serviceSpec := swarm.ServiceSpec{
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
	
	serviceOpts := docker.CreateServiceOptions{
		Auth: docker.AuthConfiguration{},
		ServiceSpec: serviceSpec,
	}

	_, err := d.Client.CreateService(serviceOpts)
	if err != nil {
		fmt.Printf("Error Creating service %+v\n",err)
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



// func (d Dockerswarm) Pull(image string) {
// 	//Pull image from Registry
// 	opts := docker.PullImageOptions{Repository: image}
// 	err := d.Client.PullImage(opts, docker.AuthConfiguration{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("Docker Image Downloaded")
// }

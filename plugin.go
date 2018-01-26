package main

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-plugins-helpers/authorization"
)

func NewNoRootPlugin() (*noroot, error) {

	return &noroot{}, nil
}

type ContainerCreateConfig struct {
	container.Config
	HostConfig       container.HostConfig
	NetworkingConfig network.NetworkingConfig
}

type noroot struct {
}

func (p *noroot) AuthZReq(req authorization.Request) authorization.Response {
	ruri, err := url.QueryUnescape(req.RequestURI)
	if err != nil {
		return authorization.Response{Err: err.Error()}
	}

	log.Println("RequestMethod:", req.RequestMethod)
	log.Println("ruri:", ruri)
	if !strings.HasSuffix(ruri, "containers/create") {
		return authorization.Response{Allow: true}
	}

	var create ContainerCreateConfig
	if err := json.Unmarshal(req.RequestBody, &create); err != nil {
		return authorization.Response{Err: err.Error()}
	}

	if create.HostConfig.UsernsMode == "host" {
		return authorization.Response{
			Allow: false,
			Err:   "--userns=host not allowed"}
	}

	if create.HostConfig.UTSMode == "host" {
		return authorization.Response{
			Allow: false,
			Err:   "--uts=host not allowed"}
	}

	if create.HostConfig.PidMode == "host" {
		return authorization.Response{
			Allow: false,
			Err:   "--pid=host not allowed"}
	}

	if create.HostConfig.NetworkMode == "host" {
		return authorization.Response{
			Allow: false,
			Err:   "--net=host not allowed"}
	}

	if len(create.HostConfig.LogConfig.Type) > 0 {
		return authorization.Response{
			Allow: false,
			Err:   "--log-driver not allowed"}
	}

	if len(create.HostConfig.LogConfig.Config) > 0 {
		return authorization.Response{
			Allow: false,
			Err:   "--log-opt not allowed"}
	}

	if len(create.HostConfig.CapAdd) > 0 {
		return authorization.Response{
			Allow: false,
			Err:   "--cap-add not allowed"}
	}

	if len(create.HostConfig.Devices) > 0 {
		return authorization.Response{
			Allow: false,
			Err:   "--device not allowed"}
	}

	if len(create.HostConfig.SecurityOpt) > 0 {
		return authorization.Response{
			Allow: false,
			Err:   "--security-opt not allowed"}
	}

	if create.HostConfig.Privileged {
		return authorization.Response{
			Allow: false,
			Err:   "--privileged not allowed"}
	}

	return authorization.Response{Allow: true}
}

func (p *noroot) AuthZRes(req authorization.Request) authorization.Response {
	return authorization.Response{Allow: true}
}

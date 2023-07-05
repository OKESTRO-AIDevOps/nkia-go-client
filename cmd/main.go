package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/OKESTRO-AIDevOps/npia-go-client/goclient"
)

func InitGoClient() error {

	cmd := exec.Command("mkdir", "-p", "srv")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed init npia go client: %s", err.Error())
	}

	get_kubeconfig_path_command_string :=
		`#!/bin/bash
[[ ! -z "$KUBECONFIG" ]] && echo "$KUBECONFIG" || echo "$HOME/.kube/config"`

	get_kubeconfig_path_command_b := []byte(get_kubeconfig_path_command_string)

	err = os.WriteFile("srv/get_kubeconfig_path", get_kubeconfig_path_command_b, 0755)

	if err != nil {

		return fmt.Errorf("failed init npia go client: %s", err.Error())
	}

	return nil

}

func main() {

	err := InitGoClient()

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	// interactive cmd

}

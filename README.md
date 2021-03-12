## Task 0: Install a ubuntu 16.04 server 64-bit
1. First Virtualbox package was downloaded on the host by using the following command:
sudo apt install virtualbox
2. Next, the iso image of ubuntu server 16.04 was downloaded from the given link.
Using virtualbox, a new virtual machine with the type of Linux and version Ubuntu(64-bit) and 4GB of memory and a virtual hard disk of 10GB(dynamically allocated) was created.
3. I started the newly created virtual machine from VirtualBox and selected the ubuntu server image .iso file to boot from the virtual optical disk of the virtual machine and followed the instructions to set up ubuntu server 16.04. 
4. In the network settings for the virtual machine, inside the Adapter 1 tab and also Port Forwarding, I set the rules as given.



## Task 1: Update system
1. I ssh’d into the vm using :  ssh -p 2222 ramtin@127.0.0.1 and then did
sudo apt upgrade
2. I checked the linux version : uname -mrs
3. Installed the latest kernel version(4.15) using: 
   sudo apt-get install --install-recommends \ linux-generic-hwe-16.04 xserver-xorg-hwe-16.04 -y
4. Then updated the grub boot loader and restarted the vm:
        a. sudo update-grub
        b. sudo reboot

## Task 2: Install gitlab-ce version in the host
1. I installed gitlab-ce from the given link
2. Trying to access localhost:8080 from my host was unsuccessful, after some inspection, it turned out that port 8080 on my host was already being used by puma.
3. Instead of 8080, I forwarded 8585 to 80 of the VM for gitlab.
4. Made sure the external_url of the gitlab is correctly set in /etc/gitlab/gitlab.rb to “http://127.0.0.1” so that it did not have to use https and certificates 
5. Reconfigured the gitlab(sudo gitlab-ctl reconfigure) and it was ready at localhost:8585
6. Then I set the new password for my gitlab account
7. Lastly, I added my public SSH key to gitlab account to be able to push/pull


## Task 3: Create a demo group/project in gitlab
1. I created a project and added main.go to the project repo
2. I also cloned the repo on the vm: 
  git clone http://127.0.0.1/demo/go-web-hello-world.git

## Task 4: Build the app and expose ($ go run) the service to 8081 port
1. First I needed to install the latest golang-go(v1.16.2) 
a. wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz
b. sudo tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz
2. I ran "go run main.go" and then from my host did curl http://127.0.0.1:8081 ==> it showed "Go Web Hello World!" right under the last command.

## Task 5: Install docker
I installed Docker according to the given link and ran the hello-world docker image. It executed successfully.

## Task 6: Run the app in container
1. First I built a Dockerfile in the repo for my main.go. 
2. After building the image and trying to run it on port 8082, it was revealed that this port was being used by another program called "sidekiq". So I changed portforwarding rule to 8083 and 8083. 
3. Next problem when building the image was that the golang needed a main module, so I ran: go mod init app/main
4. Finally I created the image from the docker file and ran it successfully.
  a. sudo docker build -t my-go-app .
  b. sudo docker run -it -p 8083:8083 my-go-app
  c. did curl http://127.0.0.1:8083 from my host and saw the "Go Web Hello World!" on my vm terminal 
4. I pushed the Dockerfile to the repo

## Task 7: Push image to dockerhub
Lastly, I tagged the image using  dockram/go-web-hello-world:v0.1 and pushed it to docker hub.

## Task 8: Document the procedure in a MarkDown file
I documented steps 0-7 as instructed.

## Task 9: Install a single node Kubernetes cluster using kubeadm
I followed the instructions in the given link and finally ran the command “kubeadm init” to start a control plane and cluster. 

## Task 10: Deploy the hello world container
In this step, after creating the yaml file and creating the pod, I realized that the status of the pod was stuck at pending. I created another pod using a different yaml to check whether the problem was the container/yaml file or something else. After using “kubectl describe” I realized that the master node had gone into DiskPressure status and also the pod for the hello-world app had failed in scheduling. I checked my disk usage and also saw that the 10GB limit was almost reached. There was a message saying "not enough disk" Therefore, I tried to increase the size of my VM storage using an alternative method with no guaranteed results: 
1. I used the command “VBoxManage clonehd UbuntuServer16.04.vdi UbuntuServer16.04-new.vdi --format VDI --variant Standard” to make a new copy of the disk,
2. enlarged the new copy to 25GB using : “VBoxManage modifyhd UbuntuServer16.04-new.vdi --resize 25000” 
3. Replaced the old storage with the new one in the virtual box manager however it did not work as the VM would not detect the increase in size
4. In the end I had to create a another VM from scratch with larger size of storage. 

After setting up the new VM, I installed kubernetes again. Next, I installed Calico first for pod networking:
1. curl https://docs.projectcalico.org/manifests/calico.yaml -O
2. kubectl apply -f calico.yaml
I tried to create a sample deployment from a yaml file from the web, but it got stuck in the pending status. After getting the description of the pod for that deployment, it said that the reason was “FailedScheduling” and the master node had taint which had to be removed. So I removed it by using :
sudo kubectl taint nodes myubuntuserver16 node-role.kubernetes.io/master-

Finally, I got my container up and running on a pod. I faced some challenges:
1. I tried to use kubernete’s port-forwarding to be able to curl http://127.0.0.1:31080, but still could not reach the container from my local host. It would give the error of connection reset by peer.
2. Next, I tried creating a service instead of the pod, and then communicating through the kubernetes proxy with the container, yet using curl http://127.0.0.1:31080 would give an error of connection refused.
3. I checked the port 31080 to make sure it was available on both ends and it was all ok.
4. I am sure that I had missed something while learning about pod communication and kubeadm server API, and made a mistake in the setup. 
5. Unfortunately I ran out of time at this point and did not get to dive deeper in debugging and could not get to tasks 11,12.
 
The deployment yaml is pushed to the repo.





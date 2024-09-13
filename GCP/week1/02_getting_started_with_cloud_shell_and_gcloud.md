# Getting started with cloud shell and gcloud

* Set the zone to \<ZONE>

```sh
gcloud config set compute/zone ZONE
```

* Get the zone

```sh
gcloud config get-value compute/zone
```

* How to get the project id using gcloud

```sh
gcloud config get-value project
```

* In Cloud Shell, run the following gcloud command to view details about the project:

```sh
gcloud compute project-info describe --project $(gcloud config get-value project)
```

 Find the zone and region metadata values in the output. You'll use the zone (google-compute-default-zone) from the output later in this lab.

 Note: When the google-compute-default-region and google-compute-default-zone keys and values are missing from the output, no default zone or region is set. The output includes other useful information regarding your project. Take some time to explore this in more detail.


 ### Setting environment variables

1. Create an environment variable to store your Project ID:
```sh
export PROJECT_ID=$(gcloud config get-value project)
```

2. Create an environment variable to store your Zone:

```sh
export ZONE=$(gcloud config get-value compute/zone)
```

3. To verify that your variables were set properly, run the following commands: 

```sh
echo -e "PROJECT ID: $PROJECT_ID\nZONE: $ZONE"
```

## Creating a virtual machine with the gcloud tool

1. To create a VM instance, run the following command

```sh
gcloud compute instances create gcelab2 --machine-type e2-medium --zone $ZONE
```

**Command details**

* gcloud compute allows you to manage your Compute Engine resources in a format that's simpler than the Compute Engine API.
* instances create creates a new instance.
* gcelab2 is the name of the VM.
* The --machine-type flag specifies the machine type as e2-medium.
* The --zone flag specifies where the VM is created.
* If you omit the --zone flag, the gcloud tool can infer your desired zone based on your default properties. Other required instance settings, such as machine type and image, are set to default values if not specified in the create command.



To open help for the create command, run the following command:

```sh
gcloud compute instances create --help
```

### Exploring gcloud commands

The gcloud tool offers simple usage guidelines that are available by adding the -h flag (for help) onto the end of any gcloud command.

1. Run the following command:

```sh
gcloud -h
```

You can access more verbose help by appending the --help flag onto a command or running the gcloud help command.

2. Run the following command:

```sh
gcloud config --help
```

```sh
gcloud help config
```

The results of the gcloud config --help and gcloud help config commands are equivalent. Both return long, detailed help.

3. View the list of configurations in your environment:

```sh
gcloud config list
```

4. To see all properties and their settings:

```sh
gcloud config list --all
```

5. List your components:

```sh
gcloud components list
```

## Filtering command-line output

The gcloud command-line interface (CLI) is a powerful tool for working at the command line. You may want specific information to be displayed.

1. List the compute instance available in the project:

```sh
gcloud compute instances list
```

**Note:** Having multiple resources deployed in a project is very common. Fortunately gcloud has some clever formatting that can help identify specific resources.

2. List the gcelab2 virtual machine:

```sh
gcloud compute instances list --filter="name=('gcelab2')"
```
In the above command, you asked gcloud to only show the information matching the criteria i.e. a virtual instance name matching the criteria.


3. List the firewall rules in the project:

```sh
gcloud compute firewall-rules list
```

4. List the firewall rules for the default network:

```sh
gcloud compute firewall-rules list --filter="network='default'"
```

5. List the firewall rules for the default network where the allow rule matches an ICMP rule:


```sh
gcloud compute firewall-rules list --filter="NETWORK:'default' AND ALLOW:'icmp'"
```

## Connecting to your VM instance

gcloud compute makes connecting to your instances easy. 
The gcloud compute ssh command provides a wrapper around SSH, which takes care of authentication and the mapping of instance names to IP addresses.

1. To connect to your VM with SSH, run the following command:

```sh
gcloud compute ssh gcelab2 --zone $ZONE
```

2. Install nginx web server on to virtual machine:

```sh
sudo apt install -y nginx
```

## Updating the firewall

When using compute resources such as virtual machines, it's important to understand the associated firewall rules.

1. List the firewall rules for the project:

```sh
gcloud compute firewall-rules list
```

**Output**

```
NAME                         NETWORK      DIRECTION  PRIORITY  ALLOW                         DENY  DISABLED
default-allow-icmp           default      INGRESS    65534     icmp                                False
default-allow-internal       default      INGRESS    65534     tcp:0-65535,udp:0-65535,icmp        False
default-allow-rdp            default      INGRESS    65534     tcp:3389                            False
default-allow-ssh            default      INGRESS    65534     tcp:22                              False
dev-net-allow-ssh            dev-network  INGRESS    1000      tcp:22                              False
serverless-to-vpc-connector  dev-network  INGRESS    1000      icmp,udp:665-666,tcp:667            False
vpc-connector-egress         dev-network  INGRESS    1000      icmp,udp,tcp                        False
vpc-connector-health-check   dev-network  INGRESS    1000      tcp:667                             False
vpc-connector-to-serverless  dev-network  EGRESS     1000      icmp,udp:665-666,tcp:667            False

```

From the above you can see there are two networks available. The default network is where the virtual machine gcelab2 is located.

2. Try to access the nginx service running on the gcelab2 virtual machine.

 **Note:** Communication with the virtual machine will fail as it does not have an appropriate firewall rule. The nginx web server is expecting to communicate on tcp:80. To get communication working you need to:

* Add a tag to the gcelab2 virtual machine
* Add a firewall rule for http traffic

3. Add a tag to the virtual machine:

```sh
gcloud compute instances add-tags gcelab2 --tags http-server,https-server
```

4. Update the firewall rule to allow:

```sh
gcloud compute firewall-rules create default-allow-http \
--direction=INGRESS \
--priority=1000 \
--network=default \
--action=ALLOW \
--rules=tcp:80 \
--source-ranges=0.0.0.0/0 \
--target-tags=http-server
```

5. List the firewall rules for the project:

```sh
gcloud compute firewall-rules list --filter=ALLOW:'80'
```

6. Verify communication is possible for http to the virtual machine:

```sh
curl http://$(gcloud compute instances list --filter=name:gcelab2 --format='value(EXTERNAL_IP)')
```

## Viewing the system logs

Viewing logs is essential to understanding the working of your project. Use gcloud to access the different logs available on Google Cloud.

1. View the available logs on the system:

```sh
gcloud logging logs list
```

2. View the logs that relate to compute resources:

```sh
gcloud logging logs list --filter="compute"
```

3. Read the logs related to the resource type of gce_instance:

```sh
gcloud logging read "resource.type=gce_instance" --limit 5
```

4. Read the logs for a specific virtual machine:

```sh
gcloud logging read "resource.type=gce_instance AND labels.instance_name='gcelab2'" --limit 5
```
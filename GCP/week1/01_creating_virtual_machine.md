# Creating a virtual machine

## Understanding Regions and Zones

* Certain Compute Engine resources live in regions or zones. A region is a specific geographical location where you can run your resources. Each region has one or more zones. For example, the us-central1 region denotes a region in the Central United States that has zones us-central1-a, us-central1-b, us-central1-c, and us-central1-f.


| Regions | Zones |
|---------|---------|
|Western US|us-west1-a, us-west1-b |
|Central US|us-central1-a, us-central1-b, us-central1-d, us-central1-f|
|Eastern US	 |us-east1-b, us-east1-c, us-east1-d|
|Western Europe|europe-west1-b, europe-west1-c, europe-west1-d|
|Eastern Asia|asia-east1-a, asia-east1-b, asia-east1-c|

* Resources that live in a zone are referred to as zonal resources. Virtual machine Instances and persistent disks live in a zone. To attach a persistent disk to a virtual machine instance, both resources must be in the same zone. Similarly, if you want to assign a static IP address to an instance, the instance must be in the same region as the static IP.

### Set the region and zone

1. Set the project region for this lab:

```sh
gcloud config set compute/region REGION
```

2. Create a variable for region:

```sh
export REGION=REGION_NAME
```

3. Create a variable for zone:

```sh
export ZONE=Zone_name
```

## Create a new instance with gcloud

* Instead of using the Cloud console to create a VM instance, use the command line tool gcloud, which is pre-installed in Google Cloud Shell. Cloud Shell is an interactive shell environment for Google Cloud loaded with all the development tools you need (gcloud, git, and others) and offers a persistent 5-GB home directory.

1. In the Cloud Shell, use gcloud to create a new VM instance from the command line:

```sh
gcloud compute instances create gcelab2 --machine-type e2-medium --zone=$ZONE
```

**Expected Output**

```sh
     Created [...gcelab2].
     NAME: gcelab2
     ZONE:  Zone
     MACHINE_TYPE: e2-medium
     PREEMPTIBLE:
     INTERNAL_IP: 10.128.0.3
     EXTERNAL_IP: 34.136.51.150
     STATUS: RUNNING
```

2. To see all the defaults, run:

```sh
gcloud compute instances create --help
```

3. You can also use SSH to connect to your instance via gcloud. Make sure to add your zone, or omit the --zone flag if you've set the option globally:

```sh
gcloud compute ssh gcelab2 --zone=$ZONE
```


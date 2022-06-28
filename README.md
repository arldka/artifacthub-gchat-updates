# ArtifactHUB Google Chat Notifications
A simple Go Google Cloud Function that handles ArtifactHUB Webhooks and relays them to a Google Chat space. 

# Setup

## Create a Google Chat webhook

### 1. (Optional) Create a new room for testing

### 2. Select the "Configure webhooks" menu:

<img alt="Configure webhooks" src="docs/images/configure_webhooks.png" height="300"/>

### 3. Create the new webhook

<img alt="Configure webhooks" src="docs/images/create_webhook_1.png" height="300"/>

### 4. Copy the webhook url

<img alt="Configure webhooks" src="docs/images/create_webhook_2.png" height="300"/>

## Create Cloud Functions

### 1. Create new Google Cloud Functions

<img alt="Configure webhooks" src="docs/images/create_cloud_function_1.png" height="300"/>

### 2. Select Google Cloud Storage ZIP source that points to the sourcecode and set the HTTP endpoint to notificationHandler

<img alt="Configure webhooks" src="docs/images/create_cloud_function_2.png" height="300"/>

### 3. Retreive the Cloud Functions Webhook URL

<img alt="Configure webhooks" src="docs/images/create_cloud_function_3.png" height="300"/>

## ArtifactHUB Setup

### 1. Go to Settings Webhook

<img alt="Configure webhooks" src="docs/images/artifact_hub_setup_1.png" height="300"/>

### 2. Create the Webhook and Paste the Cloud function HTTP Trigger url in the URL Post Section

<img alt="Configure webhooks" src="docs/images/artifact_hub_setup_2.png" height="300"/>

### 3. Add the Packages you want notifications about in the Packages section

<img alt="Configure webhooks" src="docs/images/artifact_hub_setup_3.png"/>

### 4. Click add to create the Webhook

# Local Development

Don't forget to set the following env variables !


```
export FUNCTION_TARGET=NotificationHandler
export WEBHOOK_URL=https://chat.googleapis.com/v1/spaces/AAA...AAA/messages?key=BBB...BBB

```
export FUNCTION_TARGET=NotificationHandler

export WEBHOOK_URL=https://chat.googleapis.com/v1/spaces/AAA...AAA/messages?key=BBB...BBB

The function used the cloud functions sdk.
You can run it like so:

```
cd cmd
go build
./cmd
```
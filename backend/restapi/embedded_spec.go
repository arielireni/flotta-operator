// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Flotta backend API for remote data storage",
    "title": "Flotta backend API",
    "contact": {
      "name": "Flotta flotta",
      "url": "https://github.com/project-flotta",
      "email": "flotta@redhat.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/flotta-backend/v1",
  "paths": {
    "/namespaces/{namespace}/devices/{device-id}/configuration": {
      "get": {
        "description": "Returns the device configuration",
        "tags": [
          "backend"
        ],
        "operationId": "GetDeviceConfiguration",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/device-configuration-response"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/device-configuration-response"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/devices/{device-id}/enrolment": {
      "post": {
        "description": "Initiates the process of enrolling the device",
        "tags": [
          "backend"
        ],
        "operationId": "EnrolDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          },
          {
            "name": "enrolment-info",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "swagger.yaml#/definitions/enrolment-info"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "201": {
            "description": "Created"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/devices/{device-id}/heartbeat": {
      "put": {
        "description": "Updates the heartbeat information of the device.",
        "tags": [
          "backend"
        ],
        "operationId": "UpdateHeartBeat",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          },
          {
            "name": "heartbeat",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "swagger.yaml#/definitions/heartbeat"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/devices/{device-id}/registration": {
      "get": {
        "description": "Returns a device registration status, which can be registered, unregistered or unknown.",
        "tags": [
          "backend"
        ],
        "operationId": "GetRegistrationStatus",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/device-registration-status-response"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/device-registration-status-response"
            }
          }
        }
      },
      "put": {
        "description": "Registers the device by providing its hardware configuration",
        "tags": [
          "backend"
        ],
        "operationId": "RegisterDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          },
          {
            "name": "registration-info",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "swagger.yaml#/definitions/registration-info"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Updated"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "device-configuration-response": {
      "type": "object",
      "properties": {
        "device-configuration": {
          "$ref": "swagger.yaml#/definitions/device-configuration-message"
        },
        "message": {
          "description": "Exposes the error message generated at the backend when there is an error (example HTTP code 500).",
          "type": "string"
        }
      }
    },
    "device-registration-status-response": {
      "type": "object",
      "properties": {
        "message": {
          "description": "Exposes the error message generated at the backend when there is an error (example HTTP code 500).",
          "type": "string"
        },
        "status": {
          "description": "Returns the device registration status, which can be one of the following {registered, unregistered, unknown}.",
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "description": "Exposes the error message generated at the backend when there is an error (example HTTP code 500).",
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Edge device backend storage management",
      "name": "backend"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Flotta backend API for remote data storage",
    "title": "Flotta backend API",
    "contact": {
      "name": "Flotta flotta",
      "url": "https://github.com/project-flotta",
      "email": "flotta@redhat.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/flotta-backend/v1",
  "paths": {
    "/namespaces/{namespace}/devices/{device-id}/configuration": {
      "get": {
        "description": "Returns the device configuration",
        "tags": [
          "backend"
        ],
        "operationId": "GetDeviceConfiguration",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/device-configuration-response"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/device-configuration-response"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/devices/{device-id}/enrolment": {
      "post": {
        "description": "Initiates the process of enrolling the device",
        "tags": [
          "backend"
        ],
        "operationId": "EnrolDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          },
          {
            "name": "enrolment-info",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/enrolmentInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "201": {
            "description": "Created"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/devices/{device-id}/heartbeat": {
      "put": {
        "description": "Updates the heartbeat information of the device.",
        "tags": [
          "backend"
        ],
        "operationId": "UpdateHeartBeat",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          },
          {
            "name": "heartbeat",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/heartbeat"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/devices/{device-id}/registration": {
      "get": {
        "description": "Returns a device registration status, which can be registered, unregistered or unknown.",
        "tags": [
          "backend"
        ],
        "operationId": "GetRegistrationStatus",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/device-registration-status-response"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/device-registration-status-response"
            }
          }
        }
      },
      "put": {
        "description": "Registers the device by providing its hardware configuration",
        "tags": [
          "backend"
        ],
        "operationId": "RegisterDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Namespace where the device resides",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Device ID",
            "name": "device-id",
            "in": "path",
            "required": true
          },
          {
            "name": "registration-info",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/registrationInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Updated"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "EnrolmentInfoFeatures": {
      "type": "object",
      "properties": {
        "hardware": {
          "$ref": "#/definitions/hardwareInfo"
        },
        "modelName": {
          "type": "string"
        }
      }
    },
    "LogsCollectionInformationSyslogConfig": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "protocol": {
          "type": "string"
        }
      }
    },
    "boot": {
      "type": "object",
      "properties": {
        "current_boot_mode": {
          "type": "string"
        },
        "pxe_interface": {
          "type": "string"
        }
      }
    },
    "configmapList": {
      "description": "List of configmaps used by the workload",
      "type": "array",
      "items": {
        "description": "ConfigMap kubernetes yaml specification",
        "type": "string"
      }
    },
    "containerMetrics": {
      "description": "Metrics container configuration",
      "type": "object",
      "properties": {
        "disabled": {
          "type": "boolean"
        },
        "path": {
          "description": "Path to use when retrieving metrics",
          "type": "string"
        },
        "port": {
          "description": "Port to use when retrieve the metrics",
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "cpu": {
      "type": "object",
      "properties": {
        "architecture": {
          "type": "string"
        },
        "count": {
          "type": "integer"
        },
        "flags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "frequency": {
          "type": "number"
        },
        "model_name": {
          "type": "string"
        }
      }
    },
    "dataConfiguration": {
      "description": "Configuration for data transfer",
      "type": "object",
      "properties": {
        "egress": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dataPath"
          }
        },
        "ingress": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dataPath"
          }
        }
      }
    },
    "dataPath": {
      "description": "Device-to-control plane paths mapping",
      "type": "object",
      "properties": {
        "source": {
          "description": "Path in the workload container",
          "type": "string"
        },
        "target": {
          "description": "Path in the control plane storage",
          "type": "string"
        }
      }
    },
    "device-configuration-response": {
      "type": "object",
      "properties": {
        "device-configuration": {
          "$ref": "#/definitions/deviceConfigurationMessage"
        },
        "message": {
          "description": "Exposes the error message generated at the backend when there is an error (example HTTP code 500).",
          "type": "string"
        }
      }
    },
    "device-registration-status-response": {
      "type": "object",
      "properties": {
        "message": {
          "description": "Exposes the error message generated at the backend when there is an error (example HTTP code 500).",
          "type": "string"
        },
        "status": {
          "description": "Returns the device registration status, which can be one of the following {registered, unregistered, unknown}.",
          "type": "string"
        }
      }
    },
    "deviceConfiguration": {
      "type": "object",
      "properties": {
        "heartbeat": {
          "$ref": "#/definitions/heartbeatConfiguration"
        },
        "log-collection": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/logsCollectionInformation"
          }
        },
        "metrics": {
          "$ref": "#/definitions/metricsConfiguration"
        },
        "os": {
          "$ref": "#/definitions/osInformation"
        },
        "storage": {
          "$ref": "#/definitions/storageConfiguration"
        }
      }
    },
    "deviceConfigurationMessage": {
      "type": "object",
      "properties": {
        "ansible_playbook": {
          "type": "string"
        },
        "configuration": {
          "$ref": "#/definitions/deviceConfiguration"
        },
        "device_id": {
          "description": "Device identifier",
          "type": "string"
        },
        "secrets": {
          "$ref": "#/definitions/secretList"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "$ref": "#/definitions/workloadList"
        },
        "workloads_monitoring_interval": {
          "description": "Defines the interval in seconds between the attempts to evaluate the workloads status and restart those that failed",
          "type": "integer",
          "minimum": 0,
          "exclusiveMinimum": true
        }
      }
    },
    "disk": {
      "type": "object",
      "properties": {
        "bootable": {
          "type": "boolean"
        },
        "by_id": {
          "description": "by-id is the World Wide Number of the device which guaranteed to be unique for every storage device",
          "type": "string"
        },
        "by_path": {
          "description": "by-path is the shortest physical path to the device",
          "type": "string"
        },
        "drive_type": {
          "type": "string"
        },
        "hctl": {
          "type": "string"
        },
        "id": {
          "description": "Determine the disk's unique identifier which is the by-id field if it exists and fallback to the by-path field otherwise",
          "type": "string"
        },
        "io_perf": {
          "$ref": "#/definitions/ioPerf"
        },
        "is_installation_media": {
          "description": "Whether the disk appears to be an installation media or not",
          "type": "boolean"
        },
        "model": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "serial": {
          "type": "string"
        },
        "size_bytes": {
          "type": "integer"
        },
        "smart": {
          "type": "string"
        },
        "vendor": {
          "type": "string"
        },
        "wwn": {
          "type": "string"
        }
      }
    },
    "enrolmentInfo": {
      "type": "object",
      "properties": {
        "features": {
          "type": "object",
          "properties": {
            "hardware": {
              "$ref": "#/definitions/hardwareInfo"
            },
            "modelName": {
              "type": "string"
            }
          }
        },
        "target_namespace": {
          "type": "string",
          "default": "default"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "description": "Exposes the error message generated at the backend when there is an error (example HTTP code 500).",
          "type": "string"
        }
      }
    },
    "eventInfo": {
      "type": "object",
      "properties": {
        "message": {
          "description": "Message describe the event which has occured.",
          "type": "string"
        },
        "reason": {
          "description": "Reason is single word description of the subject of the event.",
          "type": "string"
        },
        "type": {
          "description": "Either 'info' or 'warn', which reflect the importance of event.",
          "type": "string",
          "enum": [
            "info",
            "warn"
          ]
        }
      }
    },
    "gpu": {
      "type": "object",
      "properties": {
        "address": {
          "description": "Device address (for example \"0000:00:02.0\")",
          "type": "string"
        },
        "device_id": {
          "description": "ID of the device (for example \"3ea0\")",
          "type": "string"
        },
        "name": {
          "description": "Product name of the device (for example \"UHD Graphics 620 (Whiskey Lake)\")",
          "type": "string"
        },
        "vendor": {
          "description": "The name of the device vendor (for example \"Intel Corporation\")",
          "type": "string"
        },
        "vendor_id": {
          "description": "ID of the vendor (for example \"8086\")",
          "type": "string"
        }
      }
    },
    "hardwareInfo": {
      "description": "Hardware information",
      "type": "object",
      "properties": {
        "boot": {
          "$ref": "#/definitions/boot"
        },
        "cpu": {
          "$ref": "#/definitions/cpu"
        },
        "disks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/disk"
          }
        },
        "gpus": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/gpu"
          }
        },
        "hostname": {
          "type": "string"
        },
        "interfaces": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/interface"
          }
        },
        "memory": {
          "$ref": "#/definitions/memory"
        },
        "system_vendor": {
          "$ref": "#/definitions/systemVendor"
        }
      }
    },
    "hardwareProfileConfiguration": {
      "type": "object",
      "properties": {
        "include": {
          "type": "boolean"
        },
        "scope": {
          "type": "string",
          "enum": [
            "full",
            "delta"
          ]
        }
      }
    },
    "heartbeat": {
      "type": "object",
      "properties": {
        "events": {
          "description": "Events produced by device worker.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/eventInfo"
          }
        },
        "hardware": {
          "$ref": "#/definitions/hardwareInfo"
        },
        "status": {
          "type": "string",
          "enum": [
            "up",
            "degraded"
          ]
        },
        "upgrade": {
          "$ref": "#/definitions/upgradeStatus"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/workloadStatus"
          }
        }
      }
    },
    "heartbeatConfiguration": {
      "type": "object",
      "properties": {
        "hardware_profile": {
          "$ref": "#/definitions/hardwareProfileConfiguration"
        },
        "period_seconds": {
          "type": "integer"
        }
      }
    },
    "imageRegistries": {
      "description": "Image registries configuration",
      "type": "object",
      "properties": {
        "authFile": {
          "description": "Image registries authfile created by executing ` + "`" + `podman login` + "`" + ` or ` + "`" + `docker login` + "`" + ` (i.e. ~/.docker/config.json). https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/#log-in-to-docker-hub describes how the file can be created and how it is structured.",
          "type": "string"
        }
      }
    },
    "interface": {
      "type": "object",
      "properties": {
        "biosdevname": {
          "type": "string"
        },
        "client_id": {
          "type": "string"
        },
        "flags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "has_carrier": {
          "type": "boolean"
        },
        "ipv4_addresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "ipv6_addresses": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "mac_address": {
          "type": "string"
        },
        "mtu": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "product": {
          "type": "string"
        },
        "speed_mbps": {
          "type": "integer"
        },
        "vendor": {
          "type": "string"
        }
      }
    },
    "ioPerf": {
      "type": "object",
      "properties": {
        "sync_duration": {
          "description": "99th percentile of fsync duration in milliseconds",
          "type": "integer"
        }
      }
    },
    "logsCollectionInformation": {
      "description": "Log collection information",
      "type": "object",
      "properties": {
        "buffer_size": {
          "type": "integer",
          "format": "int32"
        },
        "kind": {
          "type": "string"
        },
        "syslog_config": {
          "type": "object",
          "properties": {
            "address": {
              "type": "string"
            },
            "protocol": {
              "type": "string"
            }
          }
        }
      }
    },
    "memory": {
      "type": "object",
      "properties": {
        "physical_bytes": {
          "type": "integer"
        },
        "usable_bytes": {
          "type": "integer"
        }
      }
    },
    "metrics": {
      "description": "Metrics endpoint configuration",
      "type": "object",
      "properties": {
        "allow_list": {
          "$ref": "#/definitions/metricsAllowList"
        },
        "containers": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/containerMetrics"
          }
        },
        "interval": {
          "description": "Interval(in seconds) to scrape metrics endpoint.",
          "type": "integer",
          "format": "int32"
        },
        "path": {
          "description": "Path to use when retrieving metrics",
          "type": "string"
        },
        "port": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "metricsAllowList": {
      "description": "Specification of system metrics to be collected",
      "type": "object",
      "properties": {
        "names": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "metricsConfiguration": {
      "description": "Defines metrics configuration for the device",
      "type": "object",
      "properties": {
        "receiver": {
          "$ref": "#/definitions/metricsReceiverConfiguration"
        },
        "retention": {
          "$ref": "#/definitions/metricsRetention"
        },
        "system": {
          "$ref": "#/definitions/systemMetricsConfiguration"
        }
      }
    },
    "metricsReceiverConfiguration": {
      "type": "object",
      "properties": {
        "caCert": {
          "type": "string"
        },
        "request_num_samples": {
          "type": "integer"
        },
        "timeout_seconds": {
          "type": "integer"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "metricsRetention": {
      "description": "Defines metrics data retention limits",
      "type": "object",
      "properties": {
        "max_hours": {
          "description": "Maximum time in hours metrics data files should kept on the device",
          "type": "integer",
          "format": "int32"
        },
        "max_mib": {
          "description": "Maximum size of metrics stored on disk",
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "osInformation": {
      "description": "OS lifecycle information",
      "type": "object",
      "properties": {
        "automatically_upgrade": {
          "description": "automatically upgrade the OS image",
          "type": "boolean"
        },
        "commit_id": {
          "description": "the last commit ID",
          "type": "string"
        },
        "hosted_objects_url": {
          "description": "the URL of the hosted commits web server",
          "type": "string"
        }
      }
    },
    "registrationInfo": {
      "type": "object",
      "properties": {
        "certificate_request": {
          "description": "Certificate Signing Request to be signed by flotta-operator CA",
          "type": "string"
        },
        "hardware": {
          "$ref": "#/definitions/hardwareInfo"
        }
      }
    },
    "s3StorageConfiguration": {
      "type": "object",
      "properties": {
        "aws_access_key_id": {
          "type": "string"
        },
        "aws_ca_bundle": {
          "type": "string"
        },
        "aws_secret_access_key": {
          "type": "string"
        },
        "bucket_host": {
          "type": "string"
        },
        "bucket_name": {
          "type": "string"
        },
        "bucket_port": {
          "type": "integer",
          "format": "int32"
        },
        "bucket_region": {
          "type": "string"
        }
      }
    },
    "secret": {
      "type": "object",
      "properties": {
        "data": {
          "description": "The secret's data section in JSON format",
          "type": "string"
        },
        "name": {
          "description": "Name of the secret",
          "type": "string"
        }
      }
    },
    "secretList": {
      "description": "List of secrets used by the workloads",
      "type": "array",
      "items": {
        "$ref": "#/definitions/secret"
      }
    },
    "storageConfiguration": {
      "type": "object",
      "properties": {
        "s3": {
          "$ref": "#/definitions/s3StorageConfiguration"
        }
      }
    },
    "systemMetricsConfiguration": {
      "description": "System metrics gathering configuration",
      "type": "object",
      "properties": {
        "allow_list": {
          "$ref": "#/definitions/metricsAllowList"
        },
        "disabled": {
          "description": "When true, turns system metrics collection off. False by default.",
          "type": "boolean"
        },
        "interval": {
          "description": "Interval(in seconds) to scrape metrics endpoint.",
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "systemVendor": {
      "type": "object",
      "properties": {
        "manufacturer": {
          "type": "string"
        },
        "product_name": {
          "type": "string"
        },
        "serial_number": {
          "type": "string"
        },
        "virtual": {
          "description": "Whether the machine appears to be a virtual machine or not",
          "type": "boolean"
        }
      }
    },
    "upgradeStatus": {
      "description": "Upgrade status",
      "type": "object",
      "properties": {
        "current_commit_ID": {
          "type": "string"
        },
        "last_upgrade_status": {
          "type": "string",
          "enum": [
            "succeeded",
            "failed"
          ]
        },
        "last_upgrade_time": {
          "type": "string"
        }
      }
    },
    "workload": {
      "type": "object",
      "properties": {
        "configmaps": {
          "$ref": "#/definitions/configmapList"
        },
        "data": {
          "$ref": "#/definitions/dataConfiguration"
        },
        "imageRegistries": {
          "$ref": "#/definitions/imageRegistries"
        },
        "labels": {
          "description": "Workload labels",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "log_collection": {
          "description": "Log collection target for this workload",
          "type": "string"
        },
        "metrics": {
          "$ref": "#/definitions/metrics"
        },
        "name": {
          "description": "Name of the workload",
          "type": "string"
        },
        "namespace": {
          "description": "Namespace of the workload",
          "type": "string"
        },
        "specification": {
          "type": "string"
        }
      }
    },
    "workloadList": {
      "description": "List of workloads deployed to the device",
      "type": "array",
      "items": {
        "$ref": "#/definitions/workload"
      }
    },
    "workloadStatus": {
      "type": "object",
      "properties": {
        "last_data_upload": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "status": {
          "type": "string",
          "enum": [
            "deploying",
            "running",
            "crashed",
            "stopped"
          ]
        }
      }
    }
  },
  "tags": [
    {
      "description": "Edge device backend storage management",
      "name": "backend"
    }
  ]
}`))
}

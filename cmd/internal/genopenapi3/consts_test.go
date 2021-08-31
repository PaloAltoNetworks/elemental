package genopenapi3

const regolitheINI = `
[regolithe]
product_name = dummy

[transformer]
name = gaia
url = go.aporeto.io/api
author =  Aporeto Inc.
email = dev@aporeto.com
version = 1.0
`

const typemapping = `
'[][]interface{}':
  openapi3:
    type: |-
      {
        "type": "array",
        "items": {
          "type": "array",
          "items": {
            "type": "object"
          }
        }
      }

'[][]string':
  openapi3:
    type: |-
      {
        "type": "array",
        "items": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }

'[]byte':
  openapi3:
    type: |-
      {
        "type": "string"
      }

'[]map[string]interface{}':
  openapi3:
    type: |-
      {
        "type": "array",
        "items": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        }
      }

'[]map[string]string':
  openapi3:
    type: |-
      {
        "type": "array",
        "items": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }

'[]time.Time':
  openapi3:
    type: |-
      {
        "type": "array",
        "items": {
          "type": "string",
          "format": "date-time"
        }
      }

_arch_list:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_audit_profile_rule_list:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_automation_entitlements:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_automation_events:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_cap_map:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_claims:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_elemental_identifiable:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_portlist:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_rendered_policy:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_syscall_action:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_syscall_rules:
  openapi3:
    type: |-
      {
        "type": "object"
      }

_vulnerability_level:
  openapi3:
    type: |-
      {
        "type": "object"
      }

elemental.Operation:
  openapi3:
    type: |-
      {
        "type": "object"
      }

json.RawMessage:
  openapi3:
    type: |-
      {
        "type": "object"
      }

map[string][]map[string]interface{}:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "array",
          "items": {
            "type": "object",
            "additionalProperties": {
              "type": "object"
            }
          }
        }
      }

map[string][]string:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }

map[string]bool:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "boolean"
        }
      }

map[string]int:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "integer"
        }
      }

map[string]interface{}:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "object"
        }
      }

map[string]map[string][]string:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "object",
          "additionalProperties": {
            "type": "array",
            "items": {
              "type": "string"
            }
          }
        }
      }

map[string]map[string]bool:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "object",
          "additionalProperties": {
            "type": "boolean"
          }
        }
      }

map[string]map[string]cloudnetworkquerydestination:
  openapi3:
    type: |-
      {
        "type": "object"
      }

map[string]map[string]interface{}:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        }
      }

map[string]string:
  openapi3:
    type: |-
      {
        "type": "object",
        "additionalProperties": {
          "type": "string"
        }
      }

network:
  openapi3:
    type: |-
      {
        "type": "object"
      }

networklist:
  openapi3:
    type: |-
      {
        "type": "object"
      }

pctimevalue:
  openapi3:
    type: |-
      {
        "type": "object"
      }

uiparametersexpression:
  openapi3:
    type: |-
      {
        "type": "object"
      }
`

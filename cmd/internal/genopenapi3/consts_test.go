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

const typeMapping = `
'[]byte':
  openapi3:
    type: |-
      {
        "type": "string"
      }
`

{{ header }}

package {{ package_name }}

import "github.com/aporeto-inc/elemental"

func init() {

    {% for spec in specifications.values() -%}
    elemental.RegisterIdentity({{spec.entity_name}}Identity)
    {% endfor -%}
}

func IdentifiableForIdentity(identity string) elemental.Identifiable{

  switch identity {
    {% for spec in specifications.values() -%}
      case {{ spec.entity_name }}Identity.Name:
        return New{{ spec.entity_name }}()
    {% endfor -%}
    default:
      return nil
  }
}

{{ header }}
{% set glob = {'prefix': ''} -%}

package {{ package_name }}

{% if package_name != 'elemental' -%}
import "github.com/aporeto-inc/elemental"
{% set _ = glob.update({'prefix': '{{ glob.prefix }}'}) -%}
{% endif %}


func init() {

    {% for spec in specifications.values() -%}
    {{ glob.prefix }}RegisterIdentity({{spec.entity_name}}Identity)
    {% endfor -%}
}

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) {{ glob.prefix }}Identifiable{

  switch identity {
    {% for spec in specifications.values() -%}
      case {{ spec.entity_name }}Identity.Name:
        return New{{ spec.entity_name }}()
    {% endfor -%}
    default:
      return nil
  }
}

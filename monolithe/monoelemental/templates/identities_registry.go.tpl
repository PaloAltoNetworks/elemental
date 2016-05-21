{{ header }}

package {{ package_name }}

import "github.com/aporeto-inc/elemental"

func init() {

    {% for spec in specifications.values() -%}
    elemental.RegisterIdentity({{spec.entity_name}}Identity)
    {% endfor -%}
}

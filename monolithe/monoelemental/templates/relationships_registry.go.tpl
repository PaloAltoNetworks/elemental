{{ header }}
{% set glob = {'prefix': ''} -%}

package {{package_name}}

{% if package_name != 'elemental' -%}
import "github.com/aporeto-inc/elemental"
{% set _ = glob.update({'prefix': 'elemental.'}) -%}
{% endif %}

const nodocString = "[nodoc]"

var relationshipsRegistry {{ glob.prefix }}RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() {{ glob.prefix }}RelationshipsRegistry {

  return relationshipsRegistry
}


func init() {
  relationshipsRegistry = {{ glob.prefix }}RelationshipsRegistry{}

  {% for spec in specifications.values() -%}
  //
  // Main Relationship for {{spec.rest_name}}
  //
  {{spec.entity_name}}MainRelationship := &{{ glob.prefix }}Relationship{
  {% if spec.allows_get -%}
    AllowsRetrieve: true,
  {% endif -%}
  {% if spec.allows_update -%}
    AllowsUpdate: true,
  {% endif -%}
  {% if spec.allows_delete -%}
    AllowsDelete: true,
  {% endif -%}
  }

  {% for child_api in spec.child_apis -%}
  {% set child_rest_name = child_api.rest_name -%}
  {% set child_spec = specifications[child_rest_name] -%}
  {% set child_resource_name = child_spec.resource_name -%}
  {% set child_entity_name = child_spec.entity_name -%}

  // Children relationship for {{child_resource_name}} in {{spec.rest_name}}
  {{spec.entity_name}}MainRelationship.AddChild(
    {{ glob.prefix }}IdentityFromName("{{child_rest_name}}"),
    &{{ glob.prefix }}Relationship{
    {% if child_api.allows_create -%}
      AllowsCreate: true,
    {% endif -%}
    {% if child_api.allows_update and child_api.relationship == "member" -%}
      AllowsPatch: true,
    {% endif -%}
    {% if child_api.allows_get -%}
      AllowsRetrieveMany: true,
      AllowsInfo: true,
    {% endif -%}
    },
  )
  {% endfor %}
  relationshipsRegistry[{{ glob.prefix }}IdentityFromName("{{spec.rest_name}}")] = {{spec.entity_name}}MainRelationship

  {% endfor -%}
}

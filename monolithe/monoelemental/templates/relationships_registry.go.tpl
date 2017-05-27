{{ header }}
{% set glob = {'prefix': ''} %}

package {{package_name}}

{% if package_name != 'elemental' %}
import "github.com/aporeto-inc/elemental"
{% set _ = glob.update({'prefix': 'elemental.'}) %}
{% endif %}

const nodocString = "[nodoc]" // nolint: varcheck

var relationshipsRegistry {{ glob.prefix }}RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() {{ glob.prefix }}RelationshipsRegistry {

  return relationshipsRegistry
}


func init() {
  relationshipsRegistry = {{ glob.prefix }}RelationshipsRegistry{}

  {% for rest_name, relation in relationships.iteritems() %}
  relationshipsRegistry[{{ glob.prefix }}IdentityFromName("{{rest_name}}")] = &{{ glob.prefix }}Relationship{
  {% if relation['allows_create']|length > 0 %}
    AllowsCreate: map[string]bool {
      {% for parent in relation['allows_create'] %}
      "{{parent}}" : true,
      {% endfor %}
    },
  {% endif %}
  {% if relation['allows_update']|length > 0 %}
    AllowsUpdate: map[string]bool {
      {% for parent in relation['allows_update'] %}
      "{{parent}}" : true,
      {% endfor %}
    },
    {% if relation['relationships'] == "member" %}
    AllowsPatch: map[string]bool {
      {% for parent in relation['allows_update'] %}
      "{{parent}}" : true,
      {% endfor %}
    },
    {% endif %}
  {% endif %}
  {% if relation['allows_delete']|length > 0 %}
    AllowsDelete: map[string]bool {
      {% for parent in relation['allows_delete'] %}
      "{{parent}}" : true,
      {% endfor %}
    },
  {% endif %}
  {% if relation['allows_get']|length > 0 %}
    AllowsRetrieve: map[string]bool {
      {% for parent in relation['allows_get'] %}
      "{{parent}}" : true,
      {% endfor %}
    },
    AllowsRetrieveMany: map[string]bool {
      {% for parent in relation['allows_get'] %}
      "{{parent}}" : true,
      {% endfor %}
    },
    AllowsInfo: map[string]bool {
      {% for parent in relation['allows_get'] %}
      "{{parent}}" : true,
      {% endfor %}
    },
  {% endif %}
  }
  {% endfor %}
}

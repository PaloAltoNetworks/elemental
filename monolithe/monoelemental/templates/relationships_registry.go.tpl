{{ header }}
{% set glob = {'prefix': ''} %}

package {{package_name}}

{% if package_name != 'elemental' %}
import "github.com/aporeto-inc/elemental"
{% set _ = glob.update({'prefix': 'elemental.'}) %}
{% endif %}

const nodocString = "[nodoc]"

var relationshipsRegistry {{ glob.prefix }}RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() {{ glob.prefix }}RelationshipsRegistry {

  return relationshipsRegistry
}


func init() {
  relationshipsRegistry = {{ glob.prefix }}RelationshipsRegistry{}

  {% for rest_name, relation in relationships.iteritems() %}
  relationshipsRegistry[{{ glob.prefix }}IdentityFromName("{{rest_name}}")] = &{{ glob.prefix }}Relationship{
    Parents: map[string]bool{
      {% for parent in relation['parents'] %}
      "{{parent}}": true,
      {% endfor %}
    },
  {% if relation['allows_create'] %}
    AllowsCreate: true,
  {% endif %}
  {% if relation['allows_update'] and relation['relationship'] == "member" %}
    AllowsPatch: true,
  {% endif %}
  {% if relation['allows_get'] %}
    AllowsRetrieve: true,
    AllowsRetrieveMany: true,
    AllowsInfo: true,
  {% endif %}
  }
  {% endfor %}
}

{{ header }}

package {{ package_name }}

import "github.com/aporeto-inc/elemental"
{% for imp in imports -%}
import "{{imp}}"
{% endfor -%}

// {{ specification.entity_name }}Attributes represents the various attributes name of {{ specification.entity_name }}.
type {{ specification.entity_name }}Attributes int
const (
{% for attribute in specification.attributes -%}
    {{ specification.entity_name }}Attribute{{attribute.local_name[0:1].upper() + attribute.local_name[1:]}}{% if loop.index == 1 %}  {{ specification.entity_name }}Attributes = iota{% endif %}
{% endfor -%}
)

{% for attr, constant in constants.iteritems() -%}
// {{ constant.type }} represents the possible values for attribute "{{attr}}".
type {{ constant.type }} string
const (
{% for value in constant['values'] -%}
    {{ value.name }} {{ constant.type }} = "{{ value.value }}"
{% endfor -%}
)
{% endfor -%}

// {{specification.entity_name}}Identity represents the Identity of the object
var {{specification.entity_name}}Identity = elemental.Identity {
    Name:     "{{specification.rest_name}}",
    Category: "{{specification.resource_name}}",
}

{% if not specification.is_root -%}
// {{specification.entity_name_plural}}List represents a list of {{specification.entity_name_plural}}
type {{specification.entity_name_plural}}List []*{{specification.entity_name}}
{%- endif %}

{% set glob = {'identifier': ''} -%}
// {{specification.entity_name}} represents the model of a {{specification.rest_name}}
type {{specification.entity_name}} struct {
    {% for attribute in specification.attributes -%}
    {% set field_name = attribute.local_name[0:1].upper() + attribute.local_name[1:] -%}
    {% set json_tags = 'json:"%s,omitempty"' % attribute.local_name if attribute.exposed else 'json:"-"' -%}
    {% set primary_key = ',primarykey' if attribute.primary_key else '' -%}
    {% set cql_tags = 'cql:"%s%s,omitempty"' % (attribute.local_name.lower(), primary_key) if attribute.stored else 'cql:"-"' -%}
    {% set type = attribute.local_type.split(';')[0] -%}
    {% if attribute.name in constants -%}
    {% set type = constants[attribute.name]['type'] -%}
    {%- endif -%}
    {{ field_name }} {{ type }} `{{json_tags}} {{cql_tags}}`
    {% if attribute.identifier -%}
    {% set _ = glob.update({'identifier': field_name}) -%}
    {% endif -%}
    {% endfor -%}
    {%- if specification.is_root %}
    Token string `json:"APIKey,omitempty"`
    Organization string `json:"enterprise,omitempty"`
    {%- endif %}
}

// New{{specification.entity_name}} returns a new *{{specification.entity_name}}
func New{{specification.entity_name}}() *{{specification.entity_name}} {

    return &{{specification.entity_name}}{
        {% for attribute in specification.attributes -%}
        {% set field_name = attribute.local_name[0:1].upper() + attribute.local_name[1:] -%}
        {% if attribute.type == 'external' -%}
        {% set constructor = attribute.local_type.split(';')[1] -%}
        {% if constructor -%}
        {{ field_name }}: {{ constructor }},
        {% endif %}
        {% elif attribute.default_value -%}
        {% set enclosing_format = '"%s"' if attribute.type in ['string', 'enum'] else '%s' -%}
        {{field_name}}: {{ enclosing_format % attribute.default_value}},
        {% endif -%}
        {% endfor -%}
    }
}

// Identity returns the Identity of the object.
func (o *{{specification.entity_name}}) Identity() elemental.Identity {

    return {{specification.entity_name}}Identity
}

// Identifier returns the value of the object's unique identifier.
func (o *{{specification.entity_name}}) Identifier() string {

    return o.{{ glob.identifier }}
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *{{specification.entity_name}}) SetIdentifier(ID string) {

    o.{{ glob.identifier }} = ID
}

// Validate valides the current information stored into the structure.
func (o *{{specification.entity_name}}) Validate() elemental.Errors {

    errors := elemental.Errors{}

    {% for attribute in specification.attributes -%}
    {% set field_name = attribute.local_name[0:1].upper() + attribute.local_name[1:] -%}
    {% set attribute_name = attribute.local_name -%}

    {% if attribute.allowed_choices != None -%}
    if err := elemental.ValidateStringInList("{{ attribute_name }}", string(o.{{ field_name }}), []string{"{{ attribute.allowed_choices|join('", "') }}"}); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}

    {% if attribute.allowed_chars != None -%}
    if err := elemental.ValidatePattern("{{ attribute_name }}", o.{{ field_name }}, "{{ attribute.allowed_chars }}"); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}

    {% if attribute.max_length != None -%}
    if err := elemental.ValidateMaximumLength("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.max_length }}, false); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}

    {% if attribute.min_length != None -%}
    if err := elemental.ValidateMinimumLength("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.min_length }}, false); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}

    {% if attribute.max_value != None -%}
    {% if attribute.type == "float" -%}
    if err := elemental.ValidateMaximumFloat("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.max_value }}, false); err != nil {
        errors = append(errors, err)
    }

    {% else -%}
    if err := elemental.ValidateMaximumInt("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.max_value }}, false); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}
    {% endif -%}

    {% if attribute.min_value != None -%}
    {% if attribute.type == "float" -%}
    if err := elemental.ValidateMinimumFloat("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.min_value }}, false); err != nil {
        errors = append(errors, err)
    }

    {% else -%}
    if err := elemental.ValidateMinimumInt("{{ attribute_name }}", o.{{ field_name }}, {{ attribute.min_value }}, false); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}
    {% endif -%}

    {% if attribute.required -%}
    {% if attribute.type == "string" -%}
    if err := elemental.ValidateRequiredString("{{ attribute_name }}", o.{{ field_name }}); err != nil {
        errors = append(errors, err)
    }

    {% endif -%}
    {% endif -%}

    {% endfor -%}
    return errors
}

{% if specification.is_root -%}
// APIKey returns a the API Key
func (o *{{specification.entity_name}}) APIKey() string {

    return o.Token
}

// SetAPIKey sets a the API Key
func (o *{{specification.entity_name}}) SetAPIKey(key string) {

    o.Token = key
}

{% endif -%}

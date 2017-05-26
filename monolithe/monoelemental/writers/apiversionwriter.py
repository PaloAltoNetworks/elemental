# Author: Antoine Mercadal
# See LICENSE file for full LICENSE
# Copyright 2016 Aporeto.

import os
from monolithe.lib import TaskManager
from monolithe.generators.lib import TemplateFileWriter


class APIVersionWriter(TemplateFileWriter):
    """
    """

    def __init__(self, monolithe_config, api_info):
        """
        """
        super(APIVersionWriter, self).__init__(package="monoelemental")

        output = monolithe_config.get_option("output", "transformer")
        self.name = monolithe_config.get_option("name", "transformer")

        self.output_directory = "%s/elemental" % output
        self.api_info = api_info

        header_path = "%s/elemental/__code_header" % output
        if os.path.exists(header_path):
            with open(header_path, "r") as f:
                self.header_content = f.read()
        else:
            self.header_content = ""

    def perform(self, specifications):
        """
        """
        task_manager = TaskManager()
        for rest_name, specification in specifications.items():
            task_manager.start_task(method=self._write_model, specification=specification, specification_set=specifications)
        task_manager.wait_until_exit()

        self._write_registry(specifications=specifications)

        relationships = {}

        # Compute relations to create an inversed registry
        for rest_name, specification in specifications.items():
            self.add_relation(relationships=relationships, specification=specification)

            for child_api in specification.child_apis:
                child_specification = specifications[child_api.rest_name]
                self.add_relation(relationships, child_specification, specification, child_api.relationship)

        self._write_relationships(relationships=relationships)
        self._format()

    def add_relation(self, relationships, specification, from_parent_specification=None, child_relationship=''):
        """Add specification in relations registry"""
        if specification.rest_name not in relationships:
            # Add with default values
            relationships[specification.rest_name] = {
                'allows_get': specification.allows_get,
                'allows_update': specification.allows_update,
                'allows_delete': specification.allows_delete,
                'parents': [] if from_parent_specification is None else [from_parent_specification.rest_name],
                'relationship': child_relationship
            }
        else:
            # Relation already exists, so we update it
            if not relationships[specification.rest_name]['allows_get'] and specification.allows_get:
                relationships[specification.rest_name]['allows_get'] = specification.allows_get

            if not relationships[specification.rest_name]['allows_update'] and specification.allows_update:
                relationships[specification.rest_name]['allows_update'] = specification.allows_update

            if not relationships[specification.rest_name]['allows_delete'] and specification.allows_delete:
                relationships[specification.rest_name]['allows_delete'] = specification.allows_delete

            if from_parent_specification is not None and from_parent_specification.rest_name not in relationships[specification.rest_name]['parents']:
                relationships[specification.rest_name]['parents'].append(from_parent_specification.rest_name)

            if child_relationship is not '' and relationships[specification.rest_name]['relationship'] == "":
                relationships[specification.rest_name]['relationship'] = child_relationship

    def _write_model(self, specification, specification_set):
        """
        """
        filename = '%s.go' % specification.entity_name.lower()
        constants, imports = self._extract_additional_information(specification)

        self.write(destination=self.output_directory, filename=filename, template_name="model.go.tpl",
                   specification=specification,
                   specification_set=specification_set,
                   package_name=self.name,
                   header=self.header_content,
                   constants=constants,
                   imports=imports,
                   model_version=self.api_info["version"])

    def _write_registry(self, specifications):
        """
        """
        filename = 'identities_registry.go'
        self.write(destination=self.output_directory, filename=filename, template_name="identities_registry.go.tpl",
                   specifications=specifications,
                   package_name=self.name,
                   header=self.header_content,
                   model_version=self.api_info["version"])

    def _write_relationships(self, relationships):
        """
        """
        filename = 'relationships_registry.go'
        self.write(destination=self.output_directory, filename=filename, template_name="relationships_registry.go.tpl",
                   relationships=relationships,
                   package_name=self.name,
                   header=self.header_content)

    def _format(self):
        """
        """
        os.system("gofmt -w '%s' >/dev/null 2>&1" % self.output_directory)

    def _extract_additional_information(self, specification):
        """
        """
        constants = {}
        imports = []

        for attribute in specification.attributes:

            if attribute.type == 'external':
                tokens = attribute.local_type.split(';')
                if len(tokens) == 3 and tokens[2] is not '':
                    if tokens[2] not in imports:
                        imports.append(tokens[2])

            if attribute.type == 'time' and 'time' not in imports:
                imports.append('time')

            if attribute.type == 'enum':
                if attribute.allowed_choices and len(attribute.allowed_choices) > 0:

                    name = attribute.name
                    constants[name] = {}
                    go_name = name[0:1].upper() + name[1:]
                    constants[name]['type'] = "%s%sValue" % (specification.entity_name, go_name)
                    constants[name]['values'] = []

                    for choice in attribute.allowed_choices:
                        const_name = choice.replace('_', ' ').title().replace('', '')
                        const_name = const_name[0:1].upper() + const_name[1:]
                        constants[name]['values'].append({'value': choice, 'name': "%s%s%s" % (specification.entity_name, go_name, const_name)})

        return constants, imports

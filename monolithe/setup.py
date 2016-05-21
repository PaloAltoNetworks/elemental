# Author: Antoine Mercadal
# See LICENSE file for full LICENSE
# Copyright 2016 Aporeto.

from setuptools import setup, find_packages

setup(name='monoelemental',
      version='1.0',
      description='model generator package.',
      packages=find_packages(exclude=['ez_setup', 'examples', 'tests', '.git', '.gitignore', 'README.md']),
      include_package_data=True,
      entry_points={'monolithe.plugin.lang.elemental': ['info=monoelemental:plugin_info']})

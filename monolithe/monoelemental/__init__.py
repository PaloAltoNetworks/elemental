# Author: Antoine Mercadal
# See LICENSE file for full LICENSE
# Copyright 2016 Aporeto.

from .writers.apiversionwriter import APIVersionWriter
from .converter import get_type_name

__all__ = ['APIVersionWriter', 'plugin_info', 'get_type_name']


def plugin_info():
    """
    """
    return {
        'VanillaWriter': None,
        'APIVersionWriter': APIVersionWriter,
        'PackageWriter': None,
        'CLIWriter': None,
        'get_idiomatic_name': None,
        'get_type_name': get_type_name
    }

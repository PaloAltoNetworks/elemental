# Author: Antoine Mercadal
# See LICENSE file for full LICENSE
# Copyright 2016 Aporeto.


def get_type_name(type_name, sub_type=None):
    """
    """
    if type_name in ("string", "enum"):
        return "string"

    if type_name == "float":
        return "float64"

    if type_name == "boolean":
        return "bool"

    if type_name == "list":
        st = sub_type if sub_type else "interface{}"
        return "[]%s" % st

    if type_name == "integer":
        return "int"

    if type_name == "time":
        return "time.Time"

    return "interface{}"

# coding: utf-8

"""
    Merlin

    API Guide for accessing Merlin's model management, deployment, and serving functionalities

    The version of the OpenAPI document: 0.14.0
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import Any, ClassVar, Dict, List, Optional
from pydantic import BaseModel, StrictInt, StrictStr
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class ResourceRequest(BaseModel):
    """
    ResourceRequest
    """ # noqa: E501
    min_replica: Optional[StrictInt] = None
    max_replica: Optional[StrictInt] = None
    cpu_request: Optional[StrictStr] = None
    cpu_limit: Optional[StrictStr] = None
    memory_request: Optional[StrictStr] = None
    gpu_name: Optional[StrictStr] = None
    gpu_request: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["min_replica", "max_replica", "cpu_request", "cpu_limit", "memory_request", "gpu_name", "gpu_request"]

    model_config = {
        "populate_by_name": True,
        "validate_assignment": True
    }


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of ResourceRequest from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={
            },
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of ResourceRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "min_replica": obj.get("min_replica"),
            "max_replica": obj.get("max_replica"),
            "cpu_request": obj.get("cpu_request"),
            "cpu_limit": obj.get("cpu_limit"),
            "memory_request": obj.get("memory_request"),
            "gpu_name": obj.get("gpu_name"),
            "gpu_request": obj.get("gpu_request")
        })
        return _obj



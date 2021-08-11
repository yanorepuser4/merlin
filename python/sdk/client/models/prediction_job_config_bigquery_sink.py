# coding: utf-8

"""
    Merlin

    API Guide for accessing Merlin's model management, deployment, and serving functionalities  # noqa: E501

    OpenAPI spec version: 0.14.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from client.configuration import Configuration


class PredictionJobConfigBigquerySink(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'table': 'str',
        'staging_bucket': 'str',
        'result_column': 'str',
        'save_mode': 'SaveMode',
        'options': 'dict(str, str)'
    }

    attribute_map = {
        'table': 'table',
        'staging_bucket': 'staging_bucket',
        'result_column': 'result_column',
        'save_mode': 'save_mode',
        'options': 'options'
    }

    def __init__(self, table=None, staging_bucket=None, result_column=None, save_mode=None, options=None, _configuration=None):  # noqa: E501
        """PredictionJobConfigBigquerySink - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._table = None
        self._staging_bucket = None
        self._result_column = None
        self._save_mode = None
        self._options = None
        self.discriminator = None

        if table is not None:
            self.table = table
        if staging_bucket is not None:
            self.staging_bucket = staging_bucket
        if result_column is not None:
            self.result_column = result_column
        if save_mode is not None:
            self.save_mode = save_mode
        if options is not None:
            self.options = options

    @property
    def table(self):
        """Gets the table of this PredictionJobConfigBigquerySink.  # noqa: E501


        :return: The table of this PredictionJobConfigBigquerySink.  # noqa: E501
        :rtype: str
        """
        return self._table

    @table.setter
    def table(self, table):
        """Sets the table of this PredictionJobConfigBigquerySink.


        :param table: The table of this PredictionJobConfigBigquerySink.  # noqa: E501
        :type: str
        """

        self._table = table

    @property
    def staging_bucket(self):
        """Gets the staging_bucket of this PredictionJobConfigBigquerySink.  # noqa: E501


        :return: The staging_bucket of this PredictionJobConfigBigquerySink.  # noqa: E501
        :rtype: str
        """
        return self._staging_bucket

    @staging_bucket.setter
    def staging_bucket(self, staging_bucket):
        """Sets the staging_bucket of this PredictionJobConfigBigquerySink.


        :param staging_bucket: The staging_bucket of this PredictionJobConfigBigquerySink.  # noqa: E501
        :type: str
        """

        self._staging_bucket = staging_bucket

    @property
    def result_column(self):
        """Gets the result_column of this PredictionJobConfigBigquerySink.  # noqa: E501


        :return: The result_column of this PredictionJobConfigBigquerySink.  # noqa: E501
        :rtype: str
        """
        return self._result_column

    @result_column.setter
    def result_column(self, result_column):
        """Sets the result_column of this PredictionJobConfigBigquerySink.


        :param result_column: The result_column of this PredictionJobConfigBigquerySink.  # noqa: E501
        :type: str
        """

        self._result_column = result_column

    @property
    def save_mode(self):
        """Gets the save_mode of this PredictionJobConfigBigquerySink.  # noqa: E501


        :return: The save_mode of this PredictionJobConfigBigquerySink.  # noqa: E501
        :rtype: SaveMode
        """
        return self._save_mode

    @save_mode.setter
    def save_mode(self, save_mode):
        """Sets the save_mode of this PredictionJobConfigBigquerySink.


        :param save_mode: The save_mode of this PredictionJobConfigBigquerySink.  # noqa: E501
        :type: SaveMode
        """

        self._save_mode = save_mode

    @property
    def options(self):
        """Gets the options of this PredictionJobConfigBigquerySink.  # noqa: E501


        :return: The options of this PredictionJobConfigBigquerySink.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._options

    @options.setter
    def options(self, options):
        """Sets the options of this PredictionJobConfigBigquerySink.


        :param options: The options of this PredictionJobConfigBigquerySink.  # noqa: E501
        :type: dict(str, str)
        """

        self._options = options

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(PredictionJobConfigBigquerySink, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, PredictionJobConfigBigquerySink):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, PredictionJobConfigBigquerySink):
            return True

        return self.to_dict() != other.to_dict()

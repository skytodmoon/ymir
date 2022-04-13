# coding: utf-8

from __future__ import absolute_import

from datetime import date, datetime  # noqa: F401
from typing import List, Dict  # noqa: F401

from src import util
from src.swagger_models.api_response import ApiResponse  # noqa: F401,E501
from src.swagger_models.base_model_ import Model
from src.swagger_models.dataset_result_result import DatasetResultResult  # noqa: F401,E501


class DatasetResult(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """
    def __init__(self, code: int=None, request_id: str=None, message: str=None, result: DatasetResultResult=None):  # noqa: E501
        """DatasetResult - a model defined in Swagger

        :param code: The code of this DatasetResult.  # noqa: E501
        :type code: int
        :param request_id: The request_id of this DatasetResult.  # noqa: E501
        :type request_id: str
        :param message: The message of this DatasetResult.  # noqa: E501
        :type message: str
        :param result: The result of this DatasetResult.  # noqa: E501
        :type result: DatasetResultResult
        """
        self.swagger_types = {
            'code': int,
            'request_id': str,
            'message': str,
            'result': DatasetResultResult
        }

        self.attribute_map = {
            'code': 'code',
            'request_id': 'request_id',
            'message': 'message',
            'result': 'result'
        }
        self._code = code
        self._request_id = request_id
        self._message = message
        self._result = result

    @classmethod
    def from_dict(cls, dikt) -> 'DatasetResult':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The DatasetResult of this DatasetResult.  # noqa: E501
        :rtype: DatasetResult
        """
        return util.deserialize_model(dikt, cls)

    @property
    def code(self) -> int:
        """Gets the code of this DatasetResult.


        :return: The code of this DatasetResult.
        :rtype: int
        """
        return self._code

    @code.setter
    def code(self, code: int):
        """Sets the code of this DatasetResult.


        :param code: The code of this DatasetResult.
        :type code: int
        """

        self._code = code

    @property
    def request_id(self) -> str:
        """Gets the request_id of this DatasetResult.


        :return: The request_id of this DatasetResult.
        :rtype: str
        """
        return self._request_id

    @request_id.setter
    def request_id(self, request_id: str):
        """Sets the request_id of this DatasetResult.


        :param request_id: The request_id of this DatasetResult.
        :type request_id: str
        """

        self._request_id = request_id

    @property
    def message(self) -> str:
        """Gets the message of this DatasetResult.


        :return: The message of this DatasetResult.
        :rtype: str
        """
        return self._message

    @message.setter
    def message(self, message: str):
        """Sets the message of this DatasetResult.


        :param message: The message of this DatasetResult.
        :type message: str
        """

        self._message = message

    @property
    def result(self) -> DatasetResultResult:
        """Gets the result of this DatasetResult.


        :return: The result of this DatasetResult.
        :rtype: DatasetResultResult
        """
        return self._result

    @result.setter
    def result(self, result: DatasetResultResult):
        """Sets the result of this DatasetResult.


        :param result: The result of this DatasetResult.
        :type result: DatasetResultResult
        """

        self._result = result
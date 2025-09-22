import httpx
from typing import Dict, Any, Optional, TypedDict, Required

from .result import Result, ResultOK, ResultError
from .json_ext import JSONExt


class SinglebaseError(Exception):
    """Custom exception for Singlebase SDK errors."""


class PayloadType(TypedDict, total=False):
    op: Required[str]  # required
    # any arbitrary fields allowed


class Client:
    """
    Singlebase API client for synchronous and asynchronous requests.
    """

    BASE_API_URL = "https://cloud.singlebaseapis.com/api"

    def __init__(
        self,
        api_key: str,
        api_url: Optional[str] = None,
        endpoint_key: Optional[str] = None,
        headers: Optional[dict] = None,
    ):
        if not api_key:
            raise SinglebaseError("MISSING_API_KEY")
        if not api_url and not endpoint_key:
            raise SinglebaseError("MISSING_ENDPOINT_KEY")

        self._api_key = api_key
        self._headers = headers or {}

        if api_url:
            self._api_url = api_url
        else:
            self._api_url = f"{self.BASE_API_URL}/{endpoint_key}"

    @staticmethod
    def validate_payload(data: dict) -> PayloadType:
        """
        Validate that payload matches PayloadType contract:
        - Must contain 'op'
        - Must be a string
        - Allows arbitrary extra fields
        """
        if not isinstance(data, dict):
            raise TypeError("Payload must be a dict")
        if not isinstance(data.get("op"), str):
            raise TypeError("INVALID_OPERATION_TYPE")
        return data  # type: ignore

    def call(
        self,
        payload: PayloadType,
        headers: Optional[dict] = None,
        bearer_token: Optional[str] = None,
    ) -> Result:
        """Synchronously dispatch a request to the API using httpx."""
        try:
            self.validate_payload(payload)

            _headers = {
                **self._headers,
                **(headers or {}),
                "x-api-key": self._api_key,
                "x-sbc-sdk-client": "singlebase-py",
            }
            if bearer_token:
                _headers["Authorization"] = f"Bearer {bearer_token}"

            r = httpx.post(self._api_url, json=payload, headers=_headers, timeout=10.0)
            resp = JSONExt.loads(r.text)

            if r.status_code == httpx.codes.OK:
                return ResultOK(
                    data=resp.get("data"),
                    meta=resp.get("meta"),
                    status_code=200,
                    ok=True,
                )
            else:
                return ResultError(
                    error=resp.get("error"), status_code=r.status_code, ok=False
                )
        except Exception as e:
            return ResultError(error=f"EXCEPTION: {e}", status_code=500, ok=False)

    async def call_async(
        self,
        payload: PayloadType,
        headers: Optional[dict] = None,
        bearer_token: Optional[str] = None,
    ) -> Result:
        """Asynchronously dispatch a request to the API using httpx."""
        try:
            self.validate_payload(payload)

            _headers = {
                **self._headers,
                **(headers or {}),
                "x-api-key": self._api_key,
                "x-sbc-sdk-client": "singlebase-py",
            }
            if bearer_token:
                _headers["Authorization"] = f"Bearer {bearer_token}"

            async with httpx.AsyncClient(timeout=10.0) as client:
                r = await client.post(self._api_url, json=payload, headers=_headers)
                resp = JSONExt.loads(r.text)

            if r.status_code == httpx.codes.OK:
                return ResultOK(
                    data=resp.get("data"),
                    meta=resp.get("meta"),
                    status_code=200,
                    ok=True,
                )
            else:
                return ResultError(
                    error=resp.get("error"), status_code=r.status_code, ok=False
                )
        except Exception as e:
            return ResultError(error=f"EXCEPTION: {e}", status_code=500, ok=False)

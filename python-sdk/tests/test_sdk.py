import pytest
import httpx
import respx
from singlebase import (
    Client,
    Result,
    ResultOK,
    ResultError,
    JSONExt,
    upload_presigned_file,
)


def test_result_ok_and_error():
    ok = ResultOK(data={"foo": "bar"})
    err = ResultError(error="failed", status_code=400)
    assert ok.ok is True
    assert err.ok is False
    assert err.status_code == 400


def test_jsonext_serialization():
    d = {"when": "2025-01-01T00:00:00Z"}
    s = JSONExt.dumps(d)
    assert isinstance(s, str)
    back = JSONExt.loads(s)
    assert isinstance(back, dict)


def test_result_get_data_success_and_default():
    r = ResultOK(
        data={
            "address": {
                "city": {
                    "city_fullname": "San Francisco",
                    "zipcode": 94107,
                }
            }
        }
    )

    # Full data if no path
    assert isinstance(r.get_data(), dict)

    # Nested dot path
    assert r.get_data("address.city.city_fullname") == "San Francisco"
    assert r.get_data("address.city.zipcode") == 94107

    # Missing key returns default
    assert r.get_data("address.country", default="USA") == "USA"


def test_result_get_data_type_error():
    r = ResultOK(data={"user": {"id": 123}})

    # user.id is an int → trying to traverse deeper should raise
    with pytest.raises(TypeError):
        r.get_data("user.id.value")


@respx.mock
def test_client_call_success():
    route = respx.post("https://cloud.singlebaseapis.com/api/test").mock(
        return_value=httpx.Response(200, json={"data": {"msg": "ok"}, "meta": {}})
    )

    c = Client(api_key="abc", endpoint_key="test")
    r = c.dispatch({"op": "ping"})
    assert isinstance(r, ResultOK)
    assert r.data["msg"] == "ok"
    assert route.called


@respx.mock
def test_client_call_error():
    respx.post("https://cloud.singlebaseapis.com/api/test").mock(
        return_value=httpx.Response(400, json={"error": "Bad Request"})
    )

    c = Client(api_key="abc", endpoint_key="test")
    r = c.dispatch({"op": "ping"})
    assert isinstance(r, ResultError)
    assert r.error == "Bad Request"


def test_upload_file(tmp_path):
    file = tmp_path / "hello.txt"
    file.write_text("hello")

    with pytest.raises(httpx.HTTPStatusError):
        upload_presigned_file(str(file), {"url": "http://127.0.0.1:9", "fields": {}})

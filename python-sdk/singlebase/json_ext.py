import json
import arrow
import datetime


class JSONExt:
    """ 
    Utility class for JSON serialization/deserialization with datetime/arrow support.
    """

    @classmethod
    def dumps(cls, data: dict) -> str:
        """Serialize dict to JSON string, with datetime handling."""
        return json.dumps(data, default=cls._serialize)

    @classmethod
    def loads(cls, data: str):
        """Deserialize JSON string to dict, with datetime handling."""
        if not data:
            return None
        if isinstance(data, list):
            return [json.loads(v) if v else None for v in data]
        return json.loads(data, object_hook=cls._deserialize)

    @classmethod
    def _serialize(cls, o):
        return cls._timestamp_to_str(o)

    @classmethod
    def _deserialize(cls, json_dict):
        for k, v in json_dict.items():
            if isinstance(v, str) and cls._timestamp_valid(v):
                json_dict[k] = arrow.get(v)
        return json_dict

    @staticmethod
    def _timestamp_valid(dt_str) -> bool:
        try:
            datetime.datetime.fromisoformat(dt_str.replace("Z", "+00:00"))
        except Exception:
            return False
        return True

    @staticmethod
    def _timestamp_to_str(dt) -> str:
        if isinstance(dt, arrow.Arrow):
            return dt.for_json()
        elif isinstance(dt, (datetime.date, datetime.datetime)):
            return dt.isoformat()
        return dt

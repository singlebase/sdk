from dataclasses import dataclass, field, asdict
from typing import Optional, Dict, Any


@dataclass
class Result:
    """
    Represents the result of an API operation.
    """

    data: Dict[str, Any] = field(default_factory=dict)
    meta: Dict[str, Any] = field(default_factory=dict)
    ok: bool = True
    error: Optional[str] = None
    status_code: int = 200

    def to_dict(self) -> dict:
        """Convert the result object to a dictionary."""
        return asdict(self)

    def __repr__(self):
        return f"<Result ok={self.ok} status={self.status_code} error={self.error!r}>"

    def get_data(self, path: Optional[str] = None, default: Any = None) -> Any:
        """
        Retrieve a value from `data` using dot-notation.

        Args:
            path (str, optional): Dot-notation path to the key. 
                                  If None, return the full data object.
            default (Any, optional): Default value if key not found.

        Returns:
            Any: The retrieved value or default.

        Raises:
            TypeError: If traversal encounters a non-dict where a dict was expected.

        r = ResultOK(data={
            "address": {
                "city": {
                    "city_fullname": "San Francisco",
                    "zipcode": 94107
                }
            }
        })

        print(r.get_data())  
        # {'address': {'city': {'city_fullname': 'San Francisco', 'zipcode': 94107}}}

        print(r.get_data("address.city.city_fullname"))  
        # "San Francisco"
        """
        if path is None or path == "":
            return self.data

        current = self.data
        for part in path.split("."):
            if not isinstance(current, dict):
                raise TypeError(
                    f"Cannot traverse '{part}' — expected dict, got {type(current).__name__}"
                )
            if part not in current:
                return default
            current = current[part]
        return current
    
class ResultOK(Result):
    """Represents a successful API operation result."""

    def __init__(self, **kw):
        super().__init__(**kw)
        self.ok = True


class ResultError(Result):
    """Represents a failed API operation result."""

    def __init__(self, **kw):
        super().__init__(**kw)
        self.ok = False

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

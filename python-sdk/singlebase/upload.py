import httpx
import aiofiles
from pathlib import Path


def upload_presigned_file(filepath: str, data: dict) -> bool:
    """
    Synchronously upload a file using presigned URL data.

    Args:
        filepath (str): Path to the file to upload
        data (dict): Presigned URL data containing 'url' and 'fields'

    Returns:
        bool: True if upload was successful
    """
    with open(filepath, "rb") as f2u:
        files = {"file": (Path(filepath).name, f2u)}
        resp = httpx.post(data["url"], data=data["fields"], files=files)
        resp.raise_for_status()
        return True


async def upload_presigned_file_async(filepath: str, data: dict) -> bool:
    """
    Asynchronously upload a file using presigned URL data.

    Args:
        filepath (str): Path to the file to upload
        data (dict): Presigned URL data containing 'url' and 'fields'

    Returns:
        bool: True if upload was successful
    """
    async with aiofiles.open(filepath, "rb") as f2u:
        content = await f2u.read()

    files = {"file": (Path(filepath).name, content)}
    async with httpx.AsyncClient() as client:
        resp = await client.post(data["url"], data=data["fields"], files=files)
        resp.raise_for_status()
        return True

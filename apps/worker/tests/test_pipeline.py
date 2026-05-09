import pytest

from worker.pipeline import dispatch_job


@pytest.mark.asyncio
async def test_dispatch_placeholder_accepts_job():
    await dispatch_job({"job_type": "validate_video"})

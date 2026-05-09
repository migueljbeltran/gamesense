import asyncio
import logging

from worker.config import settings
from worker.pipeline import dispatch_job


logging.basicConfig(level=settings.log_level)
logger = logging.getLogger(__name__)


async def poll_forever() -> None:
    logger.info("Worker started", extra={"queue": settings.sqs_queue_url})
    while True:
        await asyncio.sleep(settings.worker_poll_interval_seconds)


async def run_once(job: dict) -> None:
    await dispatch_job(job)


if __name__ == "__main__":
    asyncio.run(poll_forever())

import logging


logger = logging.getLogger(__name__)


async def dispatch_job(job: dict) -> None:
    job_type = job.get("job_type", "unknown")
    logger.info("Dispatch placeholder", extra={"job_type": job_type})

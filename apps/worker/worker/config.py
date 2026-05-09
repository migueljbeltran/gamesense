from functools import lru_cache

from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    model_config = SettingsConfigDict(env_file=".env.dev", extra="ignore")

    environment: str = "development"
    log_level: str = "DEBUG"
    sqs_queue_url: str = "http://localstack:4566/000000000000/gamesense-jobs"
    worker_poll_interval_seconds: int = 5
    worker_max_messages_per_poll: int = 1
    artifacts_dir: str = "/artifacts/local"


@lru_cache
def get_settings() -> Settings:
    return Settings()


settings = get_settings()

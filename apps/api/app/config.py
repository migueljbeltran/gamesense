from functools import lru_cache

from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    model_config = SettingsConfigDict(env_file=".env.dev", extra="ignore")

    environment: str = "development"
    log_level: str = "DEBUG"
    api_cors_origins: str = "http://localhost:3000"
    database_url: str = "postgresql+asyncpg://gamesense:gamesense@postgres:5432/gamesense"
    database_pool_size: int = 5
    database_max_overflow: int = 10


@lru_cache
def get_settings() -> Settings:
    return Settings()


settings = get_settings()

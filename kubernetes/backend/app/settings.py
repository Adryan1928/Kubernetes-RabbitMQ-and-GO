from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    app_name: str = "StudyFlow"
    app_env: str = "development"
    database_url: str = "postgresql+psycopg://studyflow:studyflow@postgres:5432/studyflow"
    redis_host: str = "redis"
    redis_port: int = 6379
    redis_queue: str = "studyflow:plans"
    stats_cache_key: str = "studyflow:stats"

    model_config = SettingsConfigDict(env_file=".env", extra="ignore")


settings = Settings()

import json

from redis import Redis

from app.settings import settings


def get_redis_client() -> Redis:
    return Redis(host=settings.redis_host, port=settings.redis_port, decode_responses=True)


def enqueue_plan(job_id: int) -> None:
    redis_client = get_redis_client()
    redis_client.lpush(settings.redis_queue, json.dumps({"job_id": job_id}))
    redis_client.delete(settings.stats_cache_key)

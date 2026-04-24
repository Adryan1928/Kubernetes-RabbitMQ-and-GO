import json
import logging
import time
from contextlib import suppress

from sqlalchemy import select

from app.database import Base, SessionLocal, engine
from app.models import StudyPlanRequest
from app.queue import get_redis_client
from app.services import generate_study_plan
from app.settings import settings

logging.basicConfig(level=logging.INFO, format="%(asctime)s [%(levelname)s] %(message)s")
logger = logging.getLogger("studyflow-worker")

Base.metadata.create_all(bind=engine)


def process_job(job_id: int) -> None:
    db = SessionLocal()
    try:
        job = db.scalar(select(StudyPlanRequest).where(StudyPlanRequest.id == job_id))
        if not job:
            logger.warning("Job %s nao encontrado", job_id)
            return

        job.status = "processing"
        db.commit()

        time.sleep(4)

        job.generated_plan = generate_study_plan(
            learner_name=job.learner_name,
            topic=job.topic,
            level=job.level,
            duration_days=job.duration_days,
            goals=job.goals,
        )
        job.status = "completed"
        db.commit()
        logger.info("Job %s finalizado", job_id)
    finally:
        db.close()


def main() -> None:
    redis_client = get_redis_client()
    logger.info("Worker ouvindo a fila %s", settings.redis_queue)

    while True:
        message = redis_client.brpop(settings.redis_queue, timeout=5)
        if not message:
            continue

        _, raw_payload = message
        with suppress(json.JSONDecodeError, KeyError, ValueError):
            payload = json.loads(raw_payload)
            process_job(int(payload["job_id"]))
            redis_client.delete(settings.stats_cache_key)


if __name__ == "__main__":
    main()
